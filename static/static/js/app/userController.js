(function () {

    angular.module('baseApp')
        .controller('userController', ['$scope', '$rootScope', '$http', userController]);

    function userController($scope, $rootScope, $http) {
        var self = this;
        $scope.users = [];
        $scope.showLoading = true;
        $scope.selected = [];

        //called when notified about page change
        var unbind = $rootScope.$on('component.changed', function(event, component){
            console.log('User controller got event:', component);

            if (component == "users") //if our page is now visible
            {

              $http.get('/users', {}).then(function (response) { //get user data to display in table
                users = response.data;
                for (i = 0; i < users.length; i++)//build the permissions chip data
                {
                  users[i].perms = [];
                  for (y = 0; y < users[i].Permissions.length; y++)
                  {
                    users[i].perms[users[i].perms.length] = users[i].Permissions[y].Name;
                  }
                }
                $scope.users = users;
                $scope.showLoading = false;

              });
            }else { //improve performance by unloading data when our page is no longer visible
              $scope.users = [];
              $scope.showLoading = true;
            }
        });

        $scope.$on('$destroy', unbind);
    }
})();
