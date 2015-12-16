(function () {

    angular.module('baseApp')
        .controller('pluginCreateController', ['$scope', '$rootScope', '$http', '$mdDialog', '$location', '$routeParams', '$mdToast', pluginCreateController]);

    function pluginCreateController($scope, $rootScope, $http, $mdDialog, $location, $routeParams, $mdToast) {
        var self = this;


        self.buildEmptyPluginObject = function() {
          return {
            Name: "",
            Icon: "",
            Description: "",
            Enabled: false
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
            url: '/plugins/new',
            data: $scope.plugin
          }).then(function successCallback(response) {
              console.log(response);
              self.createDialog("New plugin created successfully.", "Plugins");
            }, function errorCallback(response) {
              console.log(response);
              self.createDialog("Server responded with error: " + response.data, "Server Error");
          });
        }


        self.process = function() {
          console.log($scope.plugin);
          self.create();
        };

        //random list of icons to choose from
        $scope.icons = ["add", "memory", "bug_report", "change_history", "explore", "grade", "favorite", "event",
                        "star_rate", "work", "call", "speaker_phone", "radio", "videocam", "sd_storage", "wifi_tethering",
                        "computer", "laptop", "router", "scanner", "phone_android", "directions_bus", "directions_car"];

        $scope.process = self.process;
        $scope.plugin = self.buildEmptyPluginObject();
    }
})();
