(function() {

    var app = angular.module('baseApp', [
      'md.data.table',
      'ngRoute',
      'ngMaterial']);

    //routing
    app.config(['$routeProvider',
      function($routeProvider) {
        $routeProvider.when('/admin/users', {templateUrl: '/view/users'});
        $routeProvider.when('/admin/entities', {templateUrl: '/view/entities', controller: 'entityViewerAdminController'});
    }]);

    //material colour scheme
    app.config(function($mdThemingProvider, $mdIconProvider){
      $mdThemingProvider.theme('default')
                          .primaryPalette('teal')
                          .accentPalette('brown');
    });

})();
