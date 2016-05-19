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
            codeGraph = self.codeGraph = new graphing.CodeGraph($('#codegraph-window'));

            for (var i = 0; i < graphing.blocks.length; i++)
              codeGraph.addCode(graphing.blocks[i]);
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
            console.log($scope.resource);
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
