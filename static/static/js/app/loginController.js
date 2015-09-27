(function () {

    angular.module('baseApp')
        .controller('loginController', ['$http', '$mdDialog', loginController]);

    function loginController($http, $mdDialog) {
        var self = this;
        self.isLoggingIn = false;

        self.doLogin = function() {
          self.isLoggingIn = true;

          console.log("Logging in as: ", self.username);

          $http({
            method: 'POST',
            url: '/login',
            data: {user:self.username, pass:self.password},
            headers: {'Content-Type': 'application/x-www-form-urlencoded'},
            transformRequest: function(obj) {
                var str = [];
                for(var p in obj)
                str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
                return str.join("&");
            },
          }).
          then(function(response) {
            console.log(response.data);
            if (response.data == "GOOD") //login successful
            {
              window.location.href = '/';
            } else {
              self.isLoggingIn = false;

              $mdDialog.show(
                $mdDialog.alert()
                  .parent(angular.element(document.querySelector('#popupContainer')))
                  .clickOutsideToClose(true)
                  .title('Authentication Error')
                  .content('The supplied credentials are invalid. Please try again.')
                  .ariaLabel('Login Error Dialog')
                  .ok('OK')
              );
            }
          }, function(response) {
            alert(response);
          });
        };
    }
})();
