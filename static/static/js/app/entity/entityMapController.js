(function () {

    angular.module('baseApp')
        .controller('entityMapController', ['$scope', '$rootScope', '$http', '$mdDialog', '$location', '$routeParams', '$mdToast', entityMapController]);

    function entityMapController($scope, $rootScope, $http, $mdDialog, $location, $routeParams, $mdToast) {
        var self = this;
        $scope.showLoading = false;
        $scope.connected = false;
        $scope.wasConnected = false;

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
          $http.get('/entityLocations?id='+$routeParams.entityID, {}).then(function (response) {
            locs = response.data;
            $scope.locs = locs;
            console.log($scope.locs);
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
