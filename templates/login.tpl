<!DOCTYPE html>
<html lang="en">

<head>
    <title>CNC</title>
    {!{template "headcontent"}!}
</head>

<body layout="row" ng-app="baseApp" ng-controller="loginController as lc">
    <div layout="row" flex="100">

          <div layout="row" flex="" layout-padding="" layout-fill="" layout-align="center center" class="ng-scope">
            <div flex="40" flex-lg="50" flex-md="70" flex-sm="100" class="">

              <md-toolbar style="padding: 22px; text-align: center;">
                <div>
                  <i class="material-icons" style="font-size:48px;" >person</i>
                </div>
                <h1 class="md-headline ng-scope" style="text-align: center;">Login</h1>
                <md-progress-linear md-mode="indeterminate" ng-show="lc.isLoggingIn"></md-progress-linear>
              </md-toolbar>

              <md-content class="md-whiteframe-z1 md-padding">
                <form name="login" class="ng-valid-email ng-invalid ng-invalid-required">
                  <md-input-container>
                    <label for="username" class="ng-scope">Username</label>
                    <input id="username" label="username" ng-model="lc.username" name="username" type="text" required class="ng-invalid-required">
                  </md-input-container>

                  <md-input-container>
                    <label for="password" class="ng-scope">Password</label>
                    <input id="password" label="password" name="password" ng-model="lc.password" type="password" required class="ng-invalid-required">
                  </md-input-container>

                  <button flex="" layout-fill=""
                  class="md-raised md-primary md-button md-scope"
                  ng-click="lc.doLogin()"
                  ng-disabled="lc.isLoggingIn"
                  aria-label="Log in" tabindex="0" aria-disabled="true">
                  <span style="vertical-align: middle;">{{lc.isLoggingIn === true ? "Signing in ... Please Wait" : "Sign In"}}</span> <i class="material-icons" style="vertical-align: middle;">keyboard_arrow_right</i>
                </button>
                </form>
              </md-content>


            </div>
          </div>
    </div>
    {!{template "tailcontent"}!}
    <script src="/static/js/app/loginController.js"></script>
  </body>
</html>
