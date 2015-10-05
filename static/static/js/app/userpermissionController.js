(function () {

    angular.module('baseApp')
        .controller('userpermissionControlller', ['$scope', '$rootScope', '$http', '$mdDialog', userpermissionControlller]);

    function userpermissionControlller($scope, $rootScope, $http, $mdDialog) {
        var self = this;

        self.recommendedPermissions = [
          {key:'ADMIN', display: 'Admin'},
          {key:'HR', display: 'HR'},
          {key:'ENTITYMASTER', display: 'Entity Master'}
        ];

        $scope.showLoading = true;
        $scope.currentPermissionsDisplay = [];
        //$scope.user populated at the end


        self.buildEmptyUserObject = function() {
          return {
            Firstname: "",
            Lastname: "",
            Username: "",
            Mobile: "",
            Password: "",
            MainEmail: {
              Address: ""
            },
            Permissions: [],
            AuthMethods: [],
            MainAddress: {
              Address1: "",
              Address2: "",
              City: "",
              State: "",
              Postcode: 0
            }
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

        self.load = function(username) {
          $scope.showLoading = true;
          $http.get('/user?username='+username, {}).then(function (response) {
            user = response.data;
            $scope.user = user;
            $scope.showLoading = false;
            console.log($scope.user);

            $scope.currentPermissionsDisplay = [];
            for(var i = 0; i < user.Permissions.length; i++)
            {
              $scope.currentPermissionsDisplay.push({key: user.Permissions[i].Name});
            }
          }, function errorCallback(response) {
            console.log(response);
            self.createDialog(response, "Server Error");
          });
        };


        //listen for event from userexitController that tells us that we are editing permissions
        var unbind1 = $rootScope.$on('editperms', function(event, username){
            console.log('User permissions controller got told for user', username);
            $scope.showLoading = true;
            $scope.user = self.buildEmptyUserObject()

            self.load(username);
        });

        self.addPerm = function(permName) {
          console.log(permName);
          self.searchText = "";
          self.selectedItem = null;
          self.showLoading = true;

          $http.get('/user/permission/add?username='+$scope.user.Username+'&perm='+permName, {}).then(function (response) {
            self.load($scope.user.Username);
          }, function errorCallback(response){
            self.createDialog(response, "Server Error");
            self.load($scope.user.Username);
          });
        };

        $scope.$on('$destroy', unbind1);
        $scope.user = self.buildEmptyUserObject();
    }
})();
