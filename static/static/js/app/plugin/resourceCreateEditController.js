(function () {

    angular.module('baseApp')
        .controller('resourceCreateEditController', ['$scope', '$rootScope', '$http', '$mdDialog', '$location', '$routeParams', '$mdToast', resourceCreateEditController]);

    function resourceCreateEditController($scope, $rootScope, $http, $mdDialog, $location, $routeParams, $mdToast) {
        var self = this;
        $scope.showLoading = true;
        $scope.resourceSelected = [];
        $scope.resTypes = [
          {'code': 'JSC', 'name': 'Javascript Code'},
          {'code': 'TPL', 'name': 'Template'}
        ]

        var resourceToAceModes = {'JSC': 'ace/mode/javascript', 'TPL': "ace/mode/html"};
        function setMode(mode) {
          $scope.mode = mode;
          if (mode in resourceToAceModes) {
            self.editor.session.setMode(resourceToAceModes[mode]);
          }
        }

        self.buildEmptyResourceObject = function() {
          return {
            PluginID: parseInt($routeParams.pluginID),
            Name: "",
            ResType: "JSC",
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
          $scope.resource.Data = "";
          console.log($scope.resource);

          self.processSave();
        }

        self.processSave = function() {
          $http({
            method: 'POST',
            url: '/plugins/saveresource',
            data: $scope.resource
          }).then(function successCallback(response) {
              console.log(response);
              self.createDialog("Resource saved successfully. To apply your changes, please restart the plugin.", "Plugin Resources");
            }, function errorCallback(response) {
              console.log(response);
              self.createDialog("Server responded with error: " + response.data, "Server Error");
          });
        }

        self.load = function() {
          $http.get('/resource?resourceid='+$routeParams.resourceID, {}).then(function (response) {
            resource = response.data;
            $scope.resource = resource;
            $scope.showLoading = false;

            setMode(resource.ResType);
            self.editor.setValue(atob(resource.Data));
            console.log($scope.resource);
          }, function errorCallback(response) {
            console.log(response);
            self.createDialog(response, "Server Error");
          });
        }

        $scope.process = self.process;
        $scope.resource = self.buildEmptyResourceObject();

        self.editor = ace.edit("editor");
        self.editor.setTheme("ace/theme/github");
        self.editor.session.setMode("ace/mode/javascript");

        setMode('JSC')
        self.load();

        $scope.showReference = function() {
          window.open("/ref/api");
        }

        $scope.toggleMode = function(){
          setMode($scope.mode == 'JSC' ? 'TPL' : 'JSC');
        }
    }
})();
