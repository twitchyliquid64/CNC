(function () {

    angular.module('baseApp')
        .controller('entityViewerAdminController', ['$scope', '$rootScope', '$http', '$mdDialog', entityViewerAdminController]);

    function entityViewerAdminController($scope, $rootScope, $http, $mdDialog) {
        var self = this;
        $scope.entities = [];
        $scope.showLoading = true;
        $scope.selected = [];

        self.updateEntities = function(){
          console.log("entityViewerAdminController.updateUsers()")
          $scope.showLoading = true;
          $http.get('/entities', {}).then(function (response) { //get entities data to display in table
            entities = response.data;
            $scope.entities = self.preprocessEntities(entities);
            $scope.showLoading = false;
            console.log($scope.entities);
          });
        };

        self.preprocessEntities = function(entities) {
          for(var i = 0;i < entities.length; i++) { //parse the time into objects and generate helper strings
            entities[i].createdat_mom = moment(entities[i].CreatedAt)
            entities[i].CreatedAt_time = entities[i].createdat_mom.fromNow();
            entities[i].updatedat_mom = moment(entities[i].UpdatedAt)
            entities[i].UpdatedAt_time = entities[i].updatedat_mom.fromNow();
          }
          return entities;
        }

        self.updateEntities();
        $scope.refresh = self.updateEntities;
    };

})();
