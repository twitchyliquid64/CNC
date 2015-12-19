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
              if ($scope.locs.length > 5){
                $scope.locs.pop();
              }
              self.setCurrentMarker();
            }
          });
        };

        ws.onclose = function()
        {
          $scope.$apply(function(){
            $scope.connected = false;
          });
        };





        self.setCurrentMarker = function(){
          console.log($scope.locs);
          if ($scope.locs != null && $scope.locs.length > 0)
          {
            self.currentMarker = new google.maps.Marker({
                        title: "Current position",
                        icon: new google.maps.MarkerImage("/static/img/bluecircle.png"),
                        position: new google.maps.LatLng($scope.locs[0].Latitude, $scope.locs[0].Longitude)
                    });
            self.currentMarker.setMap(self.map);

            if ($scope.locs[0].Accuracy > 1){
              // Add circle overlay and bind to marker
              var circle = new google.maps.Circle({
                map: self.map,
                radius: $scope.locs[0].Accuracy,
                fillColor: '#2525AA'
              });
              circle.bindTo('center', self.currentMarker, 'position');
            }
            self.map.panTo(new google.maps.LatLng($scope.locs[0].Latitude, $scope.locs[0].Longitude));
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

        self.initialDownloadLocationData = function(){
          $http.get('/entityLocations?limit=1&id='+$routeParams.entityID, {}).then(function (response) {
            locs = response.data;
            for(var i = 0; i < locs.length; i++)
              locs[i] = self.processRow(locs[i]);

            console.log($scope.locs);
            if (locs != []){
              $scope.locs = locs;
              self.setCurrentMarker();
            }
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
        self.initialDownloadLocationData();
    }
})();
