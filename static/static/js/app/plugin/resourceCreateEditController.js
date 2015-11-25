(function () {

    angular.module('baseApp')
        .controller('resourceCreateEditController', ['$scope', '$rootScope', '$http', '$mdDialog', '$location', '$routeParams', '$mdToast', resourceCreateEditController]);

    function resourceCreateEditController($scope, $rootScope, $http, $mdDialog, $location, $routeParams, $mdToast) {
        var self = this;
        $scope.showLoading = true;
        $scope.resourceSelected = [];

        self.buildEmptyResourceObject = function() {
          return {
            PluginID: parseInt($routeParams.pluginID),
            Name: "",
            IsExecutable: true,
            IsTemplate: false,
            Data: ""
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

        self.process = function() {
          var content = self.editor.getValue();
          $scope.resource.JSONData = content;
          console.log($scope.resource);

          $http({
            method: 'POST',
            url: '/plugins/newresource',
            data: $scope.resource
          }).then(function successCallback(response) {
              console.log(response);
              if (response.data == "GOOD") {
                self.createDialog("Resource created successfully. To apply your changes, please restart the plugin.", "Plugin Resources");
              } else {
                self.createDialog("Server responded with error: " + response.data, "Server Error");
              }
            }, function errorCallback(response) {
              console.log(response);
              self.createDialog(response.data, "Server Error");
          });
        };

        $scope.process = self.process;
        $scope.resource = self.buildEmptyResourceObject();

        self.editor = ace.edit("editor");
        self.editor.setTheme("ace/theme/github");
        self.editor.session.setMode("ace/mode/javascript");
    }
})();
