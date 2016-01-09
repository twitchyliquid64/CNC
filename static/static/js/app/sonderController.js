(function () {

    angular.module('baseApp')
        .controller('sonderController', ['loggerService', '$scope', '$interval', '$http', sonderController]);

    function sonderController(loggerService, $scope, $interval, $http) {
        var self = this;
        self.sonderDate = moment([2016, 11, 20]);

        self.timer = $interval(self.update, 60000);
        $scope.$on('$destroy', function(event) {
          $interval.cancel(self.timer);
        });
        self.update = function(){

        }

        self.getDays = function(){
          return self.sonderDate.diff(moment(), 'days');
        }
        self.getHours = function(){
          return self.sonderDate.diff(moment(), 'hours') % 24;
        }
        self.getMinutes = function(){
          return self.sonderDate.diff(moment(), 'minutes') % 60;
        }
    }
})();
