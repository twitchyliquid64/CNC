(function () {

    angular.module('baseApp')
        .controller('userController', ['$scope', '$rootScope', '$http', '$mdDialog', userController]);

    function userController($scope, $rootScope, $http, $mdDialog) {
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
          console.log("userController.updateUsers()")
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

        self.deleteUser = function(username, ev) {
          var confirm = $mdDialog.confirm()
                .title('Confirm user deletion')
                .content('Are you sure you want to delete ' + username + '\'s account?')
                .ariaLabel('Confirm user deletion')
                .targetEvent(ev)
                .ok('Yes')
                .cancel('Abort');
          $mdDialog.show(confirm).then(function() {
            $http.get('/user/delete?username='+username, {}).then(function (response) { //get user data to display in table
              $scope.showLoading = true;
              self.updateUsers();
            });
          }, function() {  });
        };

        self.resetPassword = function(username, ev) {
            $mdDialog.show({
              controller: ResetPasswordDialogController,
              template: '<md-dialog aria-label="Reset Password"  ng-cloak ">' +
              '<md-toolbar><div class="md-toolbar-tools"><h2>Reset Password</h2><span flex></span>' +
              '<md-button class="md-icon-button" ng-click="cancel()">' +
              '<md-icon md-font-library="material-icons">close</md-icon>' +
              '</md-button>' +
              '</div></md-toolbar>' +
              '' +
              '<md-dialog-content style="max-width:800px;max-height:810px; " layout="row">' +
              '<md-input-container>' +
              '<label>New Password</label>' +
              '<input ng-model="pass" flex>' +
              '</md-input-container>' +
              '<md-button ng-click="done()">Set Password </md-button>' +
              '</md-dialog-content>' +
              '</md-dialog>',
              parent: angular.element(document.body),
              targetEvent: ev,
              clickOutsideToClose:true
            })
            .then(function(pass) {
              console.log(pass);
              $http.get('/user/updatepass?username='+username+"&pass="+pass, {}).then(function (response) { //get user data to display in table
                $scope.showLoading = true;
                self.updateUsers();
              });
            }, function() {
              //cancelled
            });
        }

        //done after controller initialisation
        self.updateUsers();
    };

    function ResetPasswordDialogController($scope, $mdDialog) {
      $scope.pass = '';
      $scope.hide = function() {
        $mdDialog.hide();
      };
      $scope.cancel = function() {
        $mdDialog.cancel();
      };
      $scope.done = function() {
        $mdDialog.hide($scope.pass);
      };
    };
})();
