(function () {

    angular.module('baseApp')
        .controller('usereditController', ['$scope', '$rootScope', '$http', '$mdDialog', usereditController]);

    function usereditController($scope, $rootScope, $http, $mdDialog) {
        var self = this;
        $scope.isNewUserMode = false;
        $scope.showLoading = false; //not used ATM
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

        //called by UI to create / commit changes
        self.process = function() {
          if ($scope.isNewUserMode)
          {
            if ($scope.user.Password != "")
            {
              $scope.user.AuthMethods.push({MethodType: "PASSWD", Value: $scope.user.Password});
            }

            $http({
              method: 'POST',
              url: '/users/new',
              data: $scope.user
            }).then(function successCallback(response) {
                console.log(response);
                if (response.data == "GOOD") {
                  self.createDialog("New user created successfully.", "Users");
                } else {
                  self.createDialog("Server responded with error: " + response.data, "Server Error");
                }
              }, function errorCallback(response) {
                console.log(response);
                self.createDialog(response, "Server Error");
            });
          } else { //edit user mode
            $http({
              method: 'POST',
              url: '/users/edit',
              data: $scope.user
            }).then(function successCallback(response) {
                console.log(response);
                if (response.data == "GOOD") {
                  self.createDialog("User edited successfully.", "Users");
                } else {
                  self.createDialog("Server responded with error: " + response.data, "Server Error");
                }
              }, function errorCallback(response) {
                console.log(response);
                self.createDialog(response, "Server Error");
            });
          }
        };

        //listen for event from userController that tells us that we are creating a new user
        var unbind1 = $rootScope.$on('newuser', function(event){
            console.log('User edit controller got told to trigger for new user');
            $scope.isNewUserMode = true;
            $scope.user = self.buildEmptyUserObject();
        });

        //listen for event from userController that tells us that we are editing an existing user
        var unbind2 = $rootScope.$on('edituser', function(event, username){
            console.log('User edit controller got told to trigger for edit user', username);
            $scope.isNewUserMode = false;
            $scope.showLoading = true;
            $scope.user = self.buildEmptyUserObject()

            $http.get('/user?username='+username, {}).then(function (response) {
              user = response.data;
              $scope.user = user;
              $scope.showLoading = false;
              console.log($scope.user);
            }, function errorCallback(response) {
              console.log(response);
              self.createDialog(response, "Server Error");
          });
        });

        self.editPermsUser = function(username) {
          console.log(username);
          $rootScope.$broadcast('editperms', username);
        };


        $scope.$on('$destroy', unbind1);
        $scope.$on('$destroy', unbind2);
        $scope.user = self.buildEmptyUserObject();
    }
})();
