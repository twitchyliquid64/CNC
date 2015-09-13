<!DOCTYPE html>
<html lang="en">

<head>
    <link rel="stylesheet" href="/static/css/angular-material.min.css">
    <link rel="stylesheet" href="/static/css/r2k9.css">
    <link rel="stylesheet" href="/static/fonts/roboto.css">
    <link rel="stylesheet" href="/static/fonts/material-icons/materialicons.css">
    <meta name="viewport" content="initial-scale=1" />

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">

</head>

<body layout="column" ng-app="baseApp">
  {!{ template "bannertop" . }!}

    <div layout="row" flex>

        <div layout="column" flex="99" layout-align="center center" id="content">
            <md-content layout="row" flex="99" class="md-padding">
                <div layout="column" flex="99">
          			     <h1>System login</h1>
          		  </div>
            </md-content>
        </div>
    </div>

    <!-- Angular Material Dependencies -->
    <script src="/static/js/angular/angular.min.js"></script>
    <script src="/static/js/angular/angular-animate.min.js"></script>
    <script src="/static/js/angular/angular-aria.min.js"></script>
    <script src="/static/js/angular/angular-messages.min.js"></script>
    <script src="/static/js/angular/angular-material.min.js"></script>
    <script src="/static/js/angular/angular-dragdrop.min.js"></script>

    <!-- Base App Dependencies -->

    <!-- Module declaration needs to come first -->
    <script src="/static/js/app/baseApp.js"></script>
    <script src="/static/js/app/loggerService.js"></script>

    <!-- Module peripherals can come after in any order -->
    <script src="/static/js/app/parametricController.js"></script>
    <script src="/static/js/app/mainController.js"></script>

  </body>
</html>
