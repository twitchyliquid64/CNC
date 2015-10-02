(function () {

    angular.module('baseApp')
        .controller('usereditController', ['$scope', '$rootScope', '$http', '$mdDialog', usereditController]);

    function usereditController($scope, $rootScope, $http, $mdDialog) {
        var self = this;
        $scope.isNewUserMode = false;
        $scope.user = {
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
                  $mdDialog.show(
                    $mdDialog.alert()
                      .parent(angular.element(document.querySelector('#popupContainer')))
                      .clickOutsideToClose(true)
                      .title('Server Response')
                      .content("New user created successfully.")
                      .ariaLabel('Create User Dialog')
                      .ok('OK')
                  );
                }
              }, function errorCallback(response) {
                console.log(response);
                $mdDialog.show(
                  $mdDialog.alert()
                    .parent(angular.element(document.querySelector('#popupContainer')))
                    .clickOutsideToClose(true)
                    .title('Server Error')
                    .content(response)
                    .ariaLabel('Create User Error Dialog')
                    .ok('OK')
                );
            });
          } else { //update

          }
        };

        //listen for event from userController that tells us that we are creating a new user
        var unbind = $rootScope.$on('newuser', function(event){
            console.log('User edit controller got told to trigger for new user');
            $scope.isNewUserMode = true;
            $scope.user = {//reset the Struct
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
        });

        $scope.$on('$destroy', unbind);
    }
})();
