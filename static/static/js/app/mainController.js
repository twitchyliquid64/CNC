(function () {

    angular.module('baseApp')
        .controller('mainController', ['$mdSidenav', '$rootScope', '$location', mainController]);

    function mainController($mdSidenav, $rootScope, $location) {
        var self = this;

        self.isRoutingMode = false;
        self.focus = 'summary';

        self.activateRouted = function(route, element) {
          console.log("Now activating section: " + element + " on route: " + route);
          self.focus = element;
          $rootScope.$broadcast('component.changed', element);
          $mdSidenav('left').close()
          $location.path(route);
          self.isRoutingMode = true;
        };

        self.activate = function (element) {
          console.log("Now activating section: " + element);
          self.focus = element;
          self.isRoutingMode = false;
          $rootScope.$broadcast('component.changed', element);
          $mdSidenav('left').close()
          $location.path("/");
        };

        self.logout = function() {
          window.location.href = '/logout';
        };

        self.toggle = function () {
            $mdSidenav('left').toggle();
        };


        //runs on page load
        var p = $location.path();
        console.log("Path:", p);
        //if the page URL is set to a section, switch the UI to that section.
        if(p == '/admin/users') {
          self.activateRouted('/admin/users', 'users');
        } else if(p == '/admin/entities') {
          self.activateRouted('/admin/entities', 'entities');
        } else if(p == '/admin/entities/new') {
          self.activateRouted('/admin/entities/new', 'entity-edit');
        } else if(p == '/admin/dashboard') {
          self.activateRouted('/admin/dashboard', 'summary');
        } else if(p.startsWith('/admin/entity/')) {
          self.activateRouted(p, 'entity-edit');
        } else if(p == "/" || p == ""){ // default TODO: Make it do something else for ppl who are not admins
          self.activateRouted('/admin/dashboard', 'summary');
        }
    }
})();
