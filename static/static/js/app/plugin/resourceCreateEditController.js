(function () {

    angular.module('baseApp')
        .controller('resourceCreateEditController', ['$scope', '$rootScope', '$http', '$mdDialog', '$location', '$routeParams', '$mdToast', resourceCreateEditController]);

    function resourceCreateEditController($scope, $rootScope, $http, $mdDialog, $location, $routeParams, $mdToast) {
        var self = this;
        $scope.showLoading = true;
        $scope.resourceSelected = [];
        $scope.mode = 'js';

        if($location.path().indexOf("newresource", 0) !== -1){
          $scope.isCreateMode = true;
        } else {
          $scope.isCreateMode = false;
        }

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
          $scope.resource.Data = "";
          console.log($scope.resource);

            if ($scope.isCreateMode) {
              self.processCreate();
            } else {
              self.processSave();
            }
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

        self.processCreate = function() {
          $http({
            method: 'POST',
            url: '/plugins/newresource',
            data: $scope.resource
          }).then(function successCallback(response) {
              console.log(response);
              self.createDialog("Resource created successfully. To apply your changes, please restart the plugin.", "Plugin Resources");
            }, function errorCallback(response) {
              console.log(response);
              self.createDialog("Server responded with error: " + response.data, "Server Error");
          });
        };

        self.load = function() {
          $http.get('/resource?resourceid='+$routeParams.resourceID, {}).then(function (response) {
            resource = response.data;
            $scope.resource = resource;
            $scope.showLoading = false;
            self.editor.setValue(atob(resource.Data));
            console.log($scope.entity);
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

        if ($scope.isCreateMode) {
          $scope.showLoading = false;
        } else {
          self.load();
        }

        $scope.showReference = function() {
          window.open("/ref/api");
        }
        $scope.toggleMode = function(){
          if ($scope.mode == 'js'){
            $scope.mode = 'html';
            self.editor.session.setMode("ace/mode/html");
          }else if ($scope.mode == 'html'){
            $scope.mode = 'js';
            self.editor.session.setMode("ace/mode/javascript");
          }
        }
    }
})();
