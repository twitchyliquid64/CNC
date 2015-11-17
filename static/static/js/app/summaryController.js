(function () {

    angular.module('baseApp')
        .controller('summaryController', ['loggerService', '$scope', summaryController]);

    function summaryController(loggerService, $scope) {
        var self = this;

        $scope.getLogMsgs = function() {
          return loggerService.msgs;
        };
        $scope.isConnected = function() {
          return loggerService.connected;
        }
    }
})();
