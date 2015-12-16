(function () {

    angular.module('baseApp')
        .controller('pluginListController', ['$scope', '$rootScope', '$http', '$mdDialog', pluginListController]);

    function pluginListController($scope, $rootScope, $http, $mdDialog) {
        var self = this;
        $scope.plugins = [];
        $scope.showLoading = false;

        self.refresh = function() {
          console.log("pluginListController.refresh()")
          $scope.showLoading = true;
          $http.get('/plugins', {}).then(function (response) { //get entities data to display in table
            plugins = response.data;
            $scope.running = plugins.Running || [];
            $scope.stopped = plugins.Disabled || [];
            $scope.plugins = [];
            for(var i = 0; i < $scope.running.length ;i++)
              $scope.plugins[$scope.plugins.length] = $scope.running[i];
            for(var i = 0; i < $scope.stopped.length ;i++)
              $scope.plugins[$scope.plugins.length] = $scope.stopped[i];
            $scope.showLoading = false;
            console.log($scope.plugins);
          });
        }

        self.switchChanged = function(pluginID, state){
          console.log("pluginListController.switchChanged(): ", pluginID, state);
          $http.get('/plugins/changestate?pluginid=' + pluginID + '&state='+state, {}).then(function (response) {
            console.log(response); // TODO: Check
          }, function errorMessage(response){
            console.log(response)
            self.createDialog("Could not switch state", "Server Error");
          });
        }

        self.deletePlugin = function (pluginID, pluginName, ev) {
          var confirm = $mdDialog.confirm()
                .title('Confirm plugin deletion')
                .content('Are you sure you want to delete plugin \'' + pluginName + '\'?')
                .ariaLabel('Confirm plugin deletion')
                .targetEvent(ev)
                .ok('Yes')
                .cancel('Abort');
          $mdDialog.show(confirm).then(function() {
            $http.get('/plugins/deleteplugin?pluginid='+pluginID, {}).then(function (response) {
              $scope.showLoading = true;
              self.refresh();
            });
          }, function errorCallback(response) {
            console.log(response);
            self.createDialog(response.data, "Server Error");
          });
        }

        $scope.refresh = self.refresh;
        $scope.switchChanged = self.switchChanged;
        $scope.deletePlugin = self.deletePlugin;
        self.refresh();
    };

})();
