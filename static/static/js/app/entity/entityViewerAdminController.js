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
            $scope.entities = entities;
            $scope.showLoading = false;
            console.log($scope.entities);
          });
        };

        self.updateEntities();
        $scope.refresh = self.updateEntities;
    };

})();
