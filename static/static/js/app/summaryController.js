(function () {

    angular.module('baseApp')
        .controller('summaryController', ['loggerService', '$scope', '$interval', '$http', summaryController]);

    function summaryController(loggerService, $scope, $interval, $http) {
        var self = this;
        $scope.components = [];
        $scope.updateState = 'loading';

        $scope.getLogMsgs = function() {
          return loggerService.msgs;
        };
        $scope.isConnected = function() {
          return loggerService.connected;
        }

        self.update = function(){
        $scope.updateState = 'loading';
        $http.get('/sys-status', {}).then(function (response) {
            components = response.data;
            $scope.components = components;
            $scope.updateState = 'done';
            console.log(components);
        }, function errorCallback(response) {
            console.log(response);
            $scope.updateState = 'error';
          });
        }


        self.updateTimer = $interval(self.update, 1000 * 23);
        $scope.$on('$destroy', function() {
          $interval.cancel(self.updateTimer);
        });
        self.update();
    }
})();
