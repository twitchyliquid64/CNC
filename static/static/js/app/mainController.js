(function () {

    angular.module('baseApp')
        .controller('mainController', ['$mdSidenav', 'logger', mainController]);

    function mainController($mdSidenav, $logger) {
        var self = this;

        self.focus = 'logging';
		self.logMessages = [[0,"System started."]];
		self.startTime = Date.now();
		
		self.getLogMsgs = function() {
			return $logger.logMsgs;
		};

        self.activate = function (element) {
            self.focus = element;
        };

        self.toggle = function () {
            $mdSidenav('left').toggle();
        };

	$logger.log("Main controller initialised.");
    }
})();
