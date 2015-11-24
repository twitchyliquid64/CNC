(function () {

    angular.module('baseApp')
        .controller('pluginEditController', ['$scope', '$rootScope', '$http', '$mdDialog', '$location', '$routeParams', '$mdToast', pluginEditController]);

    function pluginEditController($scope, $rootScope, $http, $mdDialog, $location, $routeParams, $mdToast) {
        var self = this;
        $scope.showLoading = true;


        self.buildEmptyPluginObject = function() {
          return {
            ID: $routeParams.pluginID,
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

        self.load = function() {
          $http.get('/plugin?pluginid='+$routeParams.pluginID, {}).then(function (response) {
            plugin = response.data;
            $scope.plugin = plugin;
            $scope.showLoading = false;
            console.log($scope.plugin);
          }, function errorCallback(response) {
            console.log(response);
            self.createDialog(response, "Server Error");
          });
        }

        self.process = function() {
          console.log($scope.plugin);
          self.update();
        };

        //random list of icons to choose from
        $scope.icons = ["add", "memory", "bug_report", "change_history", "explore", "grade", "favorite", "event",
                        "star_rate", "work", "call", "speaker_phone", "radio", "videocam", "sd_storage", "wifi_tethering",
                        "computer", "laptop", "router", "scanner", "phone_android", "directions_bus", "directions_car"];

        $scope.process = self.process;
        $scope.plugin = self.buildEmptyPluginObject();
        self.load();
    }
})();
