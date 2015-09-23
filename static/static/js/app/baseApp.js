(function() {

    angular.module('baseApp', ['ngMaterial'])
    .config(function($mdThemingProvider, $mdIconProvider){
      $mdThemingProvider.theme('default')
                          .primaryPalette('teal')
                          .accentPalette('brown');
    });

})();
