(function() {

    angular.module('baseApp', ['md.data.table', 'ngMaterial'])
    .config(function($mdThemingProvider, $mdIconProvider){
      $mdThemingProvider.theme('default')
                          .primaryPalette('teal')
                          .accentPalette('brown');
    });

})();
