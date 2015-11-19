(function () {

  angular.module('baseApp')
    .service('loggerService', ['$rootScope', function($rootScope) {
      var self = this;
      var location = window.location;
      self.msgs = [];
      self.connected = false;
      var ws = new WebSocket("wss://" + location.hostname+(location.port ? ':'+location.port: '') + "/ws/logging");

      ws.onopen = function()
      {
        console.log("Logger ws opened.");
        $rootScope.$apply(function(){
          self.connected = true;
        });
      };

      ws.onmessage = function (evt)
      {
        var received_msg = evt.data;
        console.log(evt.data);
        $rootScope.$apply(function(){
          self.msgs.push(JSON.parse(evt.data));
          if (self.msgs.length > 15){
            self.msgs.shift();
          }
        });
      };

      ws.onclose = function()
      {
        console.log("Logger ws closed.");
        $rootScope.$apply(function(){
          self.connected = false;
        });
      };
    }]);
})();
