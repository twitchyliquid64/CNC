<!DOCTYPE html>
<html lang="en">

<head>
    <title>CNC</title>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="description" content="">
    <meta name="viewport" content="initial-scale=1, maximum-scale=1, width=device-width" />

    <link rel="stylesheet" href="/static/css/r2k9.css">
    <link rel="stylesheet" href="/static/fonts/roboto.css">
    <link rel="stylesheet" href="/static/fonts/material-icons/materialicons.css">
    <link rel="stylesheet" href="/static/css/angular-material.css">
</head>

<body layout="row" ng-app="baseApp">
    <div layout="row" flex="100">

          <div layout="row" flex="" layout-padding="" layout-fill="" layout-align="center center" class="ng-scope">
            <div flex="40" flex-lg="50" flex-md="70" flex-sm="100" class="">

              <md-toolbar style="padding: 22px; text-align: center;">
                <div>
                  <i class="material-icons" style="font-size:48px;" >person</i>
                </div>
                <h1 class="md-headline ng-scope" style="text-align: center;">Login</h1>
              </md-toolbar>

              <md-content class="md-whiteframe-z1 md-padding">
                <form name="login" class="ng-valid-email ng-invalid ng-invalid-required">
                  <md-input-container>
                    <label for="username" class="ng-scope">Username</label>
                    <input id="username" label="username" name="username" type="text" required class="ng-invalid-required">
                  </md-input-container>

                  <md-input-container>
                    <label for="password" class="ng-scope">Password</label>
                    <input id="password" label="password" name="password" type="password" required class="ng-invalid-required">
                  </md-input-container>

                  <button flex="" layout-fill=""
                  class="md-raised md-primary md-button md-scope"
                  aria-label="Log in" tabindex="0" aria-disabled="true"><span style="vertical-align: middle;">Sign in</span> <i class="material-icons" style="vertical-align: middle;">keyboard_arrow_right</i></button>
                </form>
              </md-content>


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
