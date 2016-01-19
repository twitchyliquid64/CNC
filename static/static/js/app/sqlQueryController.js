(function () {

    angular.module('baseApp').directive('myEnter', function () {
        return function (scope, element, attrs) {
            element.bind("keydown keypress", function (event) {
                if(event.which === 13) {
                    scope.$apply(function (){
                        scope.$eval(attrs.myEnter);
                    });

                    event.preventDefault();
                }
            });
        };
    });


    angular.module('baseApp')
        .controller('sqlQueryController', ['$scope', '$http', '$mdDialog', '$mdToast', '$interval', sqlQueryController]);

    function sqlQueryController($scope, $http, $mdDialog, $mdToast, $interval) {
        var self = this;
        $scope.showLoading = false;
        $scope.connected = false;
        $scope.wasConnected = false;
        $scope.tableHide = true;
        $scope.SQL = "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE';";
        $scope.tableCols = [];
        $scope.tableRows = [];
        $scope.err = "";

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
            if (msgType == "error"){
              $scope.err = d.Error;
            } else if (msgType == "columns") {
              $scope.tableCols = d.Cols;
            } else if (msgType == "data") {
              $scope.tableRows = d.Results;
              $scope.tableHide = false;
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

        $scope.do = function(){
          var sql = $scope.SQL;
          $scope.tableHide = true;
          $scope.tableCols = [];
          $scope.tableRows = [];
          $scope.err = "";
          ws.send(JSON.stringify({Type: 'query', Query: sql}));
        }

    }
})();
