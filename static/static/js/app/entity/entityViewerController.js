(function () {

    angular.module('baseApp')
        .controller('entityViewerController', ['$scope', '$http', '$mdDialog', '$location', '$routeParams', '$mdToast', '$interval', entityViewerController]);

    function entityViewerController($scope, $http, $mdDialog, $location, $routeParams, $mdToast, $interval) {
        var self = this;
        $scope.showLoading = true;
        $scope.connected = false;
        $scope.wasConnected = false;
        $scope.msgs = [];

        var ws = new WebSocket("wss://" + location.hostname+(location.port ? ':'+location.port: '') + "/ws/entityUpdates");
        $scope.$on('$destroy', function(event) {
          ws.close();
          $interval.cancel(self.timer);
          console.log("Destroying entity resources");
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
            var msgType = d.Type;
            if (msgType == "status"){
              $scope.msgs.push(d);
              self.processUpdate(d);
              if ($scope.msgs.length > 8){
                $scope.msgs.shift();
              }
            }
          });
        };

        ws.onclose = function()
        {
          $scope.$apply(function(){
            $scope.connected = false;
          });
        };




        self.processUpdate = function(data){
          $scope.entity.updatedat_mom = moment.unix(data.Created);
          $scope.entity.UpdatedAt_time = $scope.entity.updatedat_mom.fromNow();

          $scope.entity.LastStatString = data.Content;
          $scope.entity.LastStatStyle = data.Style;
          $scope.entity.LastStatIcon = data.Icon;
          $scope.entity.LastStatMeta = data.StyleMeta;
        };

        self.buildEmptyEntityObject = function() {
          return {
            ID: $routeParams.entityID,
            Name: $routeParams.entityID,
            Category: "",
            Icon: ""
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

        self.preprocessEntity = function(entities) {
            $scope.entity.createdat_mom = moment($scope.entity.CreatedAt)
            $scope.entity.CreatedAt_time = $scope.entity.createdat_mom.fromNow();
            $scope.entity.updatedat_mom = moment($scope.entity.UpdatedAt)
            $scope.entity.UpdatedAt_time = $scope.entity.updatedat_mom.fromNow();
        }

        self.updateUpdatedTime = function(){
          if($scope.entity.updatedat_mom){
            $scope.entity.UpdatedAt_time = $scope.entity.updatedat_mom.fromNow();
          }
        }
        self.timer = $interval(self.updateUpdatedTime, 5000);


        $scope.entity = self.buildEmptyEntityObject();

        $http.get('/entity?entityID='+$routeParams.entityID, {}).then(function (response) {
          entity = response.data;
          $scope.entity = entity;
          self.preprocessEntity();
          $scope.showLoading = false;
          console.log($scope.entity);
        }, function errorCallback(response) {
          console.log(response);
          self.createDialog(response, "Server Error");
        });
    }
})();
