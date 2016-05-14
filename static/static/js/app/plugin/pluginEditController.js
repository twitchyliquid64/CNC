(function () {

    angular.module('baseApp')
        .controller('pluginEditController', ['$scope', '$rootScope', '$http', '$mdDialog', '$location', '$routeParams', '$mdToast', pluginEditController]);

    function pluginEditController($scope, $rootScope, $http, $mdDialog, $location, $routeParams, $mdToast) {
        var self = this;
        $scope.showLoading = true;
        $scope.resourceSelected = [];
        $scope.loadingMode = function() {
          if ($scope.showLoading)return "indeterminate";
          return "";
        }


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

        self.createNewResource = function() {
          var newResource = {
            'ID': 0,
            'PluginID': $scope.plugin.ID,
            'Name':  'New Resource',
            'ResType': 'JSC',
            'JSONData': 'console.log("Hello world!")'
          }

          $http({
            method: 'POST',
            url: '/plugins/newresource',
            data: newResource
          }).then(function successCallback(response) {
              $scope.main.activateRouted('/admin/resource/'+response.data.ID, 'resource-form')
            }, function errorCallback(response) {
              console.log(response);
              self.createDialog("Server responded with error: " + response.data.error, "Server Error");
          });
        }

        self.load = function() {
          $scope.showLoading = true;
          $http.get('/plugin?pluginid='+$routeParams.pluginID, {}).then(function (response) {
            plugin = response.data;
            $scope.plugin = plugin;
            $scope.iconSearchText = plugin.Icon;
            $scope.showLoading = false;
            console.log($scope.plugin);
          }, function errorCallback(response) {
            console.log(response);
            self.createDialog(response, "Server Error");
          });
        }


        self.deleteResource = function(resourceID, ev) {
          var confirm = $mdDialog.confirm()
                .title('Confirm resource deletion')
                .content('Are you sure you want to delete resource ' + resourceID + '?')
                .ariaLabel('Confirm resource deletion')
                .targetEvent(ev)
                .ok('Yes')
                .cancel('Abort');
          $mdDialog.show(confirm).then(function() {
            $http.get('/plugins/deleteresource?resourceid='+resourceID, {}).then(function (response) { //get user data to display in table
              $scope.showLoading = true;
              self.load();
            });
          }, function errorCallback(response) {
            console.log(response);
            self.createDialog(response.data, "Server Error");
          });
        };


        self.update = function() {
          $http({
            method: 'POST',
            url: '/plugins/edit',
            data: $scope.plugin
          }).then(function successCallback(response) {
              console.log(response);
              $mdToast.show(
                $mdToast.simple()
                  .content('Plugin details updated successfully.')
                  .position('bottom')
                  .hideDelay(3000)
              );
            }, function errorCallback(response) {
              console.log(response);
              self.createDialog("Server responded with error: " + response.data, "Server Error");
          });
        }

        self.process = function() {
          console.log($scope.plugin);
          self.update();
        };

        var icons = ["add", "memory", "bug_report", "change_history", "explore", "grade", "favorite", "event",
                        "star_rate", "work", "call", "speaker_phone", "radio", "videocam", "sd_storage", "wifi_tethering",
                        "computer", "laptop", "router", "scanner", "phone_android", "directions_bus", "directions_car"]

        function filterByContains(array, text) {
          return array.filter(function(item) {
            return item.indexOf(text) !== -1;
          })
        }

        $scope.iconSearchText = null;
        $scope.getIcons = function() {
          console.log($scope.iconSearchText)
          return $scope.iconSearchText ? filterByContains(icons, $scope.iconSearchText) : icons;
        }

        //random list of icons to choose from
        $scope.process = self.process;
        $scope.deleteResource = self.deleteResource;
        $scope.deletePlugin = self.deletePlugin;
        $scope.createNewResource = self.createNewResource;
        $scope.plugin = self.buildEmptyPluginObject();
        self.load();
    }
})();
