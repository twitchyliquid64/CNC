(function () {

    angular.module('baseApp')
        .controller('pluginListController', ['$scope', '$rootScope', '$http', '$mdDialog', pluginListController]);

    function pluginListController($scope, $rootScope, $http, $mdDialog) {
        var self = this;
        $scope.plugins = [];
        $scope.showLoading = false;

        self.refresh = function() {
        }

        $scope.refresh = self.refresh;
    };

})();
