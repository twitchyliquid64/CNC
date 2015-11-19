(function () {

    angular.module('baseApp')
        .controller('entityViewerController', ['$scope', '$http', '$mdDialog', '$location', '$routeParams', '$mdToast', entityViewerController]);

    function entityViewerController($scope, $http, $mdDialog, $location, $routeParams, $mdToast) {
        var self = this;
        $scope.showLoading = true;
        //$scope.entity populated at the end

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
