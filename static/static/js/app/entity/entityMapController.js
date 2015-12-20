function valueOrDash(input, units){
  if (input == -100){
    return "---";
  }
  return input + units;
}

(function () {

    angular.module('baseApp')
        .controller('entityMapController', ['$scope', '$rootScope', '$http', '$mdDialog', '$location', '$routeParams', '$mdToast', '$interval', entityMapController]);

    function entityMapController($scope, $rootScope, $http, $mdDialog, $location, $routeParams, $mdToast, $interval) {
        var self = this;
        self.currentMarker = null;
        $scope.showLoading = false;
        $scope.connected = false;
        $scope.wasConnected = false;
        $scope.locs = [{TimeUpdatedString: '---', Latitude: '-', Longitude: '-', SpeedKph: '-', Course: '-', Accuracy: '-'}];
        self.historicalMarkers = [];
        $scope.historicalShown = 8;
        $scope.loadingData = true;

        var crosshairsIcon = {
          url: "/static/img/cross-hairs.gif",
          size: new google.maps.Size(30, 30),
          origin: new google.maps.Point(0, 0),
          anchor: new google.maps.Point(7, 7),
          scaledSize: new google.maps.Size(14, 14)
        };

        var spotIcon = {
          url: "/static/img/bluecircle.png",
          size: new google.maps.Size(30, 30),
          origin: new google.maps.Point(0, 0),
          anchor: new google.maps.Point(8, 8),
          scaledSize: new google.maps.Size(16, 16)
        };


        self.updateUpdatedTime = function(){
          if ($scope.locs != null && $scope.locs.length > 0)
          {
            $scope.locs[0].TimeUpdatedString = $scope.locs[0].TimeUpdated.fromNow();
          }
        }

        self.timer = $interval(self.updateUpdatedTime, 5000);






        var ws = new WebSocket("wss://" + location.hostname+(location.port ? ':'+location.port: '') + "/ws/entityUpdates?id=" + $routeParams.entityID);
        $scope.$on('$destroy', function(event) {
          ws.close();
          $interval.cancel(self.timer);
        });

        ws.onopen = function()
        {
          $scope.$apply(function(){
            $scope.connected = true;
            $scope.wasConnected = true;
          });
        };

        ws.onmessage = function (evt)
        {
          $scope.$apply(function(){
            var d = JSON.parse(evt.data);
            console.log(d);
            var msgType = d.Type;
            if (msgType == "location"){
              $scope.locs.unshift(self.processRow(d));
              if ($scope.locs.length > ($scope.historicalShown+1)){
                $scope.locs.pop();
              }
              self.updateMap();
            }
          });
        };

        ws.onclose = function()
        {
          $scope.$apply(function(){
            $scope.connected = false;
          });
        };





        self.updateMap = function(){
          console.log($scope.locs);
          if ($scope.locs != null && $scope.locs.length > 0)
          {
            //delete existing position marker / circle
            if( self.currentMarker != null)self.currentMarker.setMap(null);
            if( self.circle != null)self.circle.setMap(null);

            //create new marker
            self.currentMarker = new google.maps.Marker({
                        title: "Current position",
                        icon: spotIcon,
                        position: new google.maps.LatLng($scope.locs[0].Latitude, $scope.locs[0].Longitude)
                    });
            self.currentMarker.setMap(self.map);

            //create new accuracy circle if accuracy information is available
            if ($scope.locs[0].Accuracy > 1){
              var circle = new google.maps.Circle({
                map: self.map,
                radius: $scope.locs[0].Accuracy,
                fillColor: '#2525AA'
              });
              circle.bindTo('center', self.currentMarker, 'position');
              self.circle = circle;
            }

            //center the map on the new location
            self.map.panTo(new google.maps.LatLng($scope.locs[0].Latitude, $scope.locs[0].Longitude));

            //delete existing historical points
            for(var i = 0; i < self.historicalMarkers.length; i++) {
              self.historicalMarkers[i].setMap(null);
            }
            self.historicalMarkers = [];

            //display historical points
            var lastPoint = [$scope.locs[0].Latitude, $scope.locs[0].Longitude];
            for(var i = 1; i < $scope.locs.length; i++) {
              if ($scope.locs[i].Accuracy < 150){
                self.historicalMarkers[self.historicalMarkers.length] = new google.maps.Marker({
                      title: moment($scope.locs[i].CreatedAt).format("dddd, MMMM Do YYYY, h:mm:ss a"),
                      icon: crosshairsIcon,
                      position: new google.maps.LatLng($scope.locs[i].Latitude, $scope.locs[i].Longitude)
                  });
                self.historicalMarkers[self.historicalMarkers.length-1].setMap(self.map);

                var infowindow = new google.maps.InfoWindow({
                  content: makeContentWindow(moment($scope.locs[i].CreatedAt).format("dddd, MMMM Do YYYY, h:mm:ss a"), $scope.locs[i].Accuracy, $scope.locs[i].Course,  $scope.locs[i].SpeedKph)
                });

                b = function(win,mark){
                  mark.addListener('click', function(){
                    win.open(self.map, this);
                  });
                }
                b(infowindow, self.historicalMarkers[self.historicalMarkers.length-1]);


                var line = new google.maps.Polyline({
                  path: [{lat: lastPoint[0], lng: lastPoint[1]}, {lat: $scope.locs[i].Latitude, lng: $scope.locs[i].Longitude}],
                  strokeColor: '#FF0000',
                  strokeOpacity: 1.0,
                  strokeWeight: 2
                })
                line.setMap(self.map);
                self.historicalMarkers[self.historicalMarkers.length] = line;
                lastPoint = [$scope.locs[i].Latitude, $scope.locs[i].Longitude];
              }
            }
          }
        }

        self.downloadEntityData = function(){
          $http.get('/entity?entityID='+$routeParams.entityID, {}).then(function (response) {
            entity = response.data;
            $scope.entity = entity;
          }, function errorCallback(response) {
            console.log(response);
            self.createDialog(response, "Server Error");
          });
        }

        self.initialDownloadLocationData = function(limit){
          $scope.loadingData = true;
          $http.get('/entityLocations?limit=' + limit + '&id='+$routeParams.entityID, {}).then(function (response) {
            locs = response.data;
            for(var i = 0; i < locs.length; i++)
              locs[i] = self.processRow(locs[i]);

            console.log($scope.locs);
            if (locs != []){
              $scope.locs = locs;
              self.updateMap();
            }
            $scope.loadingData = false;
          }, function errorCallback(response) {
            console.log(response);
            self.createDialog(response, "Server Error");
          });
        }

        self.createDialog = function(message, title) {
          $mdDialog.show(
            $mdDialog.alert()
              .parent(angular.element(document.querySelector('#popupContainer')))
              .clickOutsideToClose(true)
              .title(title)
              .content(message)
              .ariaLabel(title)
              .ok('OK')
          );
        };

        self.processRow = function(row){
          ret = row;
          ret.SpeedKph = valueOrDash(ret.SpeedKph, "km/h");
          ret.Course = valueOrDash(ret.Course, "");
          ret.AccuracyDisplay = valueOrDash(ret.Accuracy, "m");
          ret.SatNum = valueOrDash(ret.SatNum, "");
          ret.TimeUpdated = moment(ret.CreatedAt);
          ret.TimeUpdatedString = ret.TimeUpdated.fromNow();
          return ret;
        }

        self.initMap = function(){
          var centerlatlng = new google.maps.LatLng(-33.915803, 151.195242);
          var myOptions = {
              zoom: 15,
              center: centerlatlng,
              mapTypeId: google.maps.MapTypeId.ROADMAP
          };
          self.map = new google.maps.Map(document.getElementById("map_canvas"), myOptions);
        }

        self.initMap();
        self.downloadEntityData();
        self.initialDownloadLocationData($scope.historicalShown);


        $scope.$watch(function(scope) {return scope.historicalShown;},
              function() {self.initialDownloadLocationData($scope.historicalShown);}
        );

    }


    function makeContentWindow(tim, acc, cours, spd){
      return '<div id="content">'+
          '<div id="siteNotice">'+
          '</div>'+
          '<h1 id="firstHeading" class="firstHeading">' +  tim + '</h1>'+
          '<div id="bodyContent">'+
          '<p><b>Accuracy: </b>' + acc + 'm</p>'+
          '<p><b>Heading: </b>' + cours + ' degrees</p>'+
          '<p><b>Speed: </b>' + spd + '</p>'+
          '</div>'+
          '</div>';
    }
})();
