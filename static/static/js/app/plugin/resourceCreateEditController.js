(function () {

    angular.module('baseApp')
        .controller('resourceCreateEditController', ['$scope', '$rootScope', '$http', '$mdDialog', '$location', '$routeParams', '$mdToast', resourceCreateEditController]);

    function resourceCreateEditController($scope, $rootScope, $http, $mdDialog, $location, $routeParams, $mdToast) {
        var self = this;
        $scope.showLoading = true;
        $scope.resourceSelected = [];
        $scope.resTypes = [
          {'code': 'JSC', 'name': 'Javascript Code'},
          {'code': 'TPL', 'name': 'Template'},
          {'code': 'GRA', 'name': 'Code Graph'}
        ]

        var resourceToAceModes = {'JSC': 'ace/mode/javascript', 'TPL': "ace/mode/html"};
        function setMode(mode) {
          $scope.mode = mode;
          if (mode in resourceToAceModes) {
            $('#codegraph-window').hide();
            $('#editor-window').show();
            setEditorMode(resourceToAceModes[mode]);
          } else if (mode == 'GRA') {
            $('#editor-window').hide()
            $('#codegraph-window').show();
            setGraphMode();
          }
        }

        function setEditorMode(aceMode) {
          if (self.editor === undefined) {
            self.editor = ace.edit("editor");
            self.editor.setTheme("ace/theme/github");
            self.editor.setValue(atob(resource.Data));
          }

          self.editor.session.setMode(aceMode);
        }

        function setGraphMode() {
          if (self.codeGraph === undefined) {
            var content = $scope.resource.Data ?
              atob($scope.resource.Data) :
              undefined;

            codeGraph = self.codeGraph = new graphing.CodeGraph($('#codegraph-window'), content);
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
          var content = $scope.mode === 'GRA' ?
            self.codeGraph.toJsonString() :
            self.editor.getValue();

          $scope.resource.JSONData = content;
          $scope.resource.Data = "";

          self.processSave();
        }

        self.processSave = function() {
          $http({
            method: 'POST',
            url: '/plugins/saveresource',
            data: $scope.resource
          }).then(function successCallback(response) {
              self.createDialog("Resource saved successfully. To apply your changes, please restart the plugin.", "Plugin Resources");
            }, function errorCallback(response) {
              console.log(response);
              self.createDialog("Server responded with error: " + response.data.error, "Server Error");
          });
        }

        self.load = function() {
          $http.get('/resource?resourceid='+$routeParams.resourceID, {}).then(function (response) {
            resource = response.data;
            $scope.resource = resource;
            $scope.showLoading = false;

            setMode(resource.ResType);
          }, function errorCallback(response) {
            console.log(response);
            self.createDialog(response, "Server Error");
          });
        }

        $scope.showGraph = false;
        $scope.process = self.process;
        $scope.resource = self.buildEmptyResourceObject();
        $scope.setMode = setMode;

        self.load();

        $scope.showReference = function() {
          window.open("/ref/api");
        }

        $scope.toggleMode = function(){
          setMode($scope.mode == 'JSC' ? 'TPL' : 'JSC');
        }
    }
})();
