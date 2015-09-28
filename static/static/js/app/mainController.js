(function () {

    angular.module('baseApp')
        .controller('mainController', ['$mdSidenav', mainController]);

    function mainController($mdSidenav) {
        var self = this;

        self.focus = 'summary';

        self.activate = function (element) {
            self.focus = element;
        };

        self.logout = function() {
          window.location.href = '/logout';
        };

        self.toggle = function () {
            $mdSidenav('left').toggle();
        };
    }
})();
