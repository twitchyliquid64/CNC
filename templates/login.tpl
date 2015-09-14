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

    <div layout="row" flex="70">

        <div layout="column" flex="99" layout-align="center center" id="content">
          <div layout="row" flex="" layout-padding="" layout-fill="" layout-align="center center" class="ng-scope">
            <div flex="40" flex-lg="50" flex-md="70" flex-sm="100">

              <md-toolbar style="padding: 22px; text-align: center;">
                <div>
                  <i class="material-icons" style="font-size:48px;" >person</i>
                </div>
                <h1 class="md-headline ng-scope" style="text-align: center;">Login</h1>
              </md-toolbar>

              <md-content class="md-padding">
                <form name="login" class="ng-valid-email ng-invalid ng-invalid-required">
                  <md-input-container>
                    <label for="username" class="ng-scope">Username</label>
                    <input id="username" label="username" name="username" type="text" required class="ng-invalid ng-invalid-required">
                  </md-input-container>

                  <md-input-container>
                    <label for="password" class="ng-scope">Password</label>
                    <input id="password" label="password" name="password" type="password" required class="ng-invalid ng-invalid-required">
                  </md-input-container>

                  <button style="width: 100%;"
                  class="md-raised md-primary md-button md-scope"
                  aria-label="Log in" tabindex="0" aria-disabled="true">Log in</button>
                </form>
              </md-content>


            </div>
          </div>
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
