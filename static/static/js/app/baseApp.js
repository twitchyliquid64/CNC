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
        $routeProvider.when('/admin/entities/new', {templateUrl: '/view/entities/form', controller: 'entityFormController'});
        $routeProvider.when('/admin/entity/:entityID', {templateUrl: '/view/entities/form', controller: 'entityFormController'});
        $routeProvider.when('/admin/dashboard', {templateUrl: '/view/dashboard/summary', controller: 'summaryController'});

        $routeProvider.when('/entity/:entityID', {templateUrl: '/view/entity', controller: 'entityViewerController'});
        $routeProvider.when('/entitymap/:entityID', {templateUrl: '/view/entitymap', controller: 'entityMapController'});

        $routeProvider.when('/admin/data', {templateUrl: '/view/data', controller: 'sqlQueryController'});


        $routeProvider.when('/admin/plugins', {templateUrl: '/view/plugins', controller: 'pluginListController'});
        $routeProvider.when('/admin/plugins/new', {templateUrl: '/view/plugins/newform', controller: 'pluginCreateController'})
        $routeProvider.when('/admin/plugin/:pluginID', {templateUrl: '/view/plugins/editform', controller: 'pluginEditController'});
        $routeProvider.when('/admin/newresource/:pluginID', {templateUrl: '/view/plugins/resourceform', controller: 'resourceCreateEditController'});
        $routeProvider.when('/admin/resource/:resourceID', {templateUrl: '/view/plugins/resourceform', controller: 'resourceCreateEditController'});

    }]);

    //material colour scheme
    app.config(function($mdThemingProvider, $mdIconProvider){
      $mdThemingProvider.theme('default')
                          .primaryPalette('teal')
                          .accentPalette('amber');
    });

})();
