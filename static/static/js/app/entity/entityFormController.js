(function () {

    angular.module('baseApp')
        .controller('entityFormController', ['$scope', '$rootScope', '$http', '$mdDialog', '$location', '$routeParams', '$mdToast', entityFormController]);

    function entityFormController($scope, $rootScope, $http, $mdDialog, $location, $routeParams, $mdToast) {
        var self = this;

        if($location.path().indexOf("new", this.length - "new".length) !== -1){ //if the url (/entities/?) ends in new
          $scope.isNewEntityMode = true;
        } else {
          $scope.isNewEntityMode = false;
        }
        $scope.showLoading = false; //not used ATM
        //$scope.entity populated at the end

        self.buildEmptyEntityObject = function() {
          return {
            Name: "",
            Category: "",
            Icon: ""
          };
        };

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

        self.create = function() {
          $http({
            method: 'POST',
            url: '/entities/new',
            data: $scope.entity
          }).then(function successCallback(response) {
              console.log(response);
              if (response.data == "GOOD") {
                self.createDialog("New entity created successfully.", "Entities");
              } else {
                self.createDialog("Server responded with error: " + response.data, "Server Error");
              }
            }, function errorCallback(response) {
              console.log(response);
              self.createDialog(response.data, "Server Error");
          });
        }

        self.saveChanges = function() {
          $http({
            method: 'POST',
            url: '/entities/edit',
            data: $scope.entity
          }).then(function successCallback(response) {
              console.log(response);
              if (response.data == "GOOD") {
                $mdToast.show(
                  $mdToast.simple()
                    .content('Entity details updated successfully.')
                    .position('bottom')
                    .hideDelay(3000)
                );
              } else {
                self.createDialog("Server responded with error: " + response.data, "Server Error");
              }
            }, function errorCallback(response) {
              console.log(response);
              self.createDialog(response.data, "Server Error");
          });
        }

        self.process = function() {
          console.log($scope.entity);
          if ($scope.isNewEntityMode) {
            self.create();
          }else {
            self.saveChanges();
          }
        };

        //random list of icons to choose from
        $scope.icons = ["add", "memory", "bug_report", "change_history", "explore", "grade", "favorite", "event",
                        "star_rate", "work", "call", "speaker_phone", "radio", "videocam", "sd_storage", "wifi_tethering",
                        "computer", "laptop", "router", "scanner", "phone_android", "directions_bus", "directions_car"];

        $scope.process = self.process;
        $scope.entity = self.buildEmptyEntityObject();

        // if we are editing an existing entity we need to load it's existing fields
        if (!$scope.isNewEntityMode) {
          $http.get('/entity?entityID='+$routeParams.entityID, {}).then(function (response) {
            entity = response.data;
            $scope.entity = entity;
            $scope.showLoading = false;
            console.log($scope.entity);
          }, function errorCallback(response) {
            console.log(response);
            self.createDialog(response, "Server Error");
          });
        }
    }
})();
