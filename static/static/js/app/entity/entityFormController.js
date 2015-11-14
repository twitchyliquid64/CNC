(function () {

    angular.module('baseApp')
        .controller('entityFormController', ['$scope', '$rootScope', '$http', '$mdDialog', '$location', entityFormController]);

    function entityFormController($scope, $rootScope, $http, $mdDialog, $location) {
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

        self.process = function() {
          console.log($scope.entity);
          if ($scope.isNewEntityMode) {
            self.create();
          }
        };

        //random list of icons to choose from
        $scope.icons = ["add", "memory", "bug_report", "change_history", "explore", "grade", "favorite", "event",
                        "star_rate", "work", "call", "speaker_phone", "radio", "videocam", "sd_storage", "wifi_tethering",
                        "computer", "laptop", "router", "scanner", "phone_android", "directions_bus", "directions_car"];

        $scope.process = self.process;
        $scope.entity = self.buildEmptyEntityObject();
    }
})();
