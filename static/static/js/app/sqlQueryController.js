(function () {

    angular.module('baseApp')
        .controller('sqlQueryController', ['$scope', '$http', '$mdDialog', '$mdToast', '$interval', sqlQueryController]);

    function sqlQueryController($scope, $http, $mdDialog, $mdToast, $interval) {
        var self = this;
        $scope.showLoading = true;
        $scope.connected = false;
        $scope.wasConnected = false;

        var ws = new WebSocket("wss://" + location.hostname+(location.port ? ':'+location.port: '') + "/ws/sql");
        $scope.$on('$destroy', function(event) {
          ws.close();
          $interval.cancel(self.timer);
          console.log("Destroying data resources");
        });
        //$scope.entity populated at the end

        ws.onopen = function()
        {
          $scope.$apply(function(){
            $scope.connected = true;
            $scope.wasConnected = true;
          });
        };

        ws.onmessage = function (evt)
        {
          $scope.$apply(function(){
            var d = JSON.parse(evt.data);
            console.log(d);

            var msgType = d.Type;
            if (msgType == "status"){

            }
          });
        };

        ws.onclose = function()
        {
          $scope.$apply(function(){
            $scope.connected = false;
          });
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

    }
})();
