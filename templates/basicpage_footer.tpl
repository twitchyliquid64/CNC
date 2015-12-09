</md-content>
</div>
</div>


    <!-- Angular Material Dependencies - angular moved to head-->
    <script src="/static/js/angular/angular-animate.min.js"></script>
    <script src="/static/js/angular/angular-route.min.js"></script>
    <script src="/static/js/angular/angular-aria.min.js"></script>
    <script src="/static/js/angular/angular-messages.min.js"></script>
    <script src="/static/js/angular/angular-material.min.js"></script>
    <script src="/static/js/angular/angular-dragdrop.min.js"></script>
    <script src="/static/js/md-data-table.min.js"></script>

    <!-- Base App Dependencies -->

    <!-- Module declaration needs to come first -->
    <script src="/static/js/app/baseApp.js"></script>

    <!-- Other dependencies -->
    <script src="/static/js/moment.min.js"></script>

<script>
(function() {

var app = angular.module('apiApp', [
  'md.data.table',
  'vAccordion',
  'ngMaterial']);

//material colour scheme
app.config(function($mdThemingProvider, $mdIconProvider){
  $mdThemingProvider.theme('default')
                      .primaryPalette('{{.PrimaryColour}}')
                      .accentPalette('{{.AccentColour}}');
});

})();
</script>

<script src="/static/js/v-accordion.min.js"></script>
</body>
</html>
