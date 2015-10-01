(function () {

    angular.module('baseApp')
        .controller('mainController', ['$mdSidenav', '$rootScope', mainController]);

    function mainController($mdSidenav, $rootScope) {
        var self = this;

        self.focus = 'summary';

        self.activate = function (element) {
          console.log("Now activating section: " + element);
          self.focus = element;
          $rootScope.$broadcast('component.changed', element);
        };

        self.logout = function() {
          window.location.href = '/logout';
        };

        self.toggle = function () {
            $mdSidenav('left').toggle();
        };
    }
})();
