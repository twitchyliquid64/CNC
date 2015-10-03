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
            if (component == "users") //if our page is now visible
            {
              self.updateUsers();
            }else { //improve performance by unloading data when our page is no longer visible
              $scope.users = [];
              $scope.showLoading = true;
            }
        });

        $scope.$on('$destroy', unbind);

        self.updateUsers = function(){
          $http.get('/users', {}).then(function (response) { //get user data to display in table
            users = response.data;
            $scope.users = users;
            $scope.showLoading = false;
          });
        };

        self.newUser = function() {
          $rootScope.$broadcast('newuser'); //tell everyone who is listening
        };

        self.editUser = function(username) {
          console.log(username);
          $rootScope.$broadcast('edituser', username);
        };

        self.deleteUser = function(username) {
          $http.get('/user/delete?username='+username, {}).then(function (response) { //get user data to display in table
            $scope.showLoading = true;
            self.updateUsers();
          });
        }
    }
})();
