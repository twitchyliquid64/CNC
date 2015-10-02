<!DOCTYPE html>
<html lang="en">

  <head>
      <title>CNC Dashboard</title>
      {!{template "headcontent"}!}
  </head>

  <body layout="column" ng-app="baseApp" ng-controller="mainController as main">
      <md-toolbar layout="row" flex="10" class="md-whiteframe-z1">
        <h1 flex><md-icon md-font-library="material-icons" style="font-size: 250%;">av_timer</md-icon> CNC</h1>

          <div class="md-toolbar-tools" flex layout-align="end center">
              <md-button ng-click="main.toggle()" hide-gt-sm class="md-icon-button">
                  <md-icon aria-label="Menu" md-svg-icon="/static/img/menu.svg"></md-icon>
              </md-button>
              <md-button ng-click="main.logout()">
                <md-icon md-font-library="material-icons">exit_to_app</md-icon>
              </md-button>
          </div>

      </md-toolbar>

    <div layout="row" flex>
      <md-sidenav class="site-sidenav md-sidenav-left md-whiteframe-z2"
                    md-component-id="left"
                    md-is-locked-open="$mdMedia('gt-sm')">

        <md-list><!--Put ng-repeat in the md-list -->
          {!{if .IsAdmin}!}
          <md-subheader class="md-no-sticky">Admin</md-subheader>
          <md-list-item>
              <md-button ng-click="main.activate('summary')">
                <md-icon md-font-library="material-icons">tune</md-icon> Summary
              </md-button>
          </md-list-item>
          <md-list-item>
              <md-button ng-click="main.activate('users')">
                <md-icon md-font-library="material-icons">people</md-icon> Users
              </md-button>
          </md-list-item>
          <md-list-item>
              <md-button ng-click="main.activate('data')">
                <md-icon md-font-library="material-icons">storage</md-icon> Data
              </md-button>
          </md-list-item>
          <md-list-item>
              <md-button ng-click="main.activate('plugins')">
                <md-icon md-font-library="material-icons">memory</md-icon> Plugins
              </md-button>
          </md-list-item>
          {!{end}!}

          <md-subheader class="md-no-sticky">Comms</md-subheader>
          <md-list-item>
              <md-button ng-click="main.activate('messenger')">
                <md-icon md-font-library="material-icons">message</md-icon> Messenger
              </md-button>
          </md-list-item>
          <md-list-item>
              <md-button ng-click="main.activate('mail')">
                <md-icon md-font-library="material-icons">email</md-icon> Mail
              </md-button>
          </md-list-item>

          <md-subheader class="md-no-sticky">Other</md-subheader>
          <md-list-item>
              <md-button ng-click="main.activate('assets')">
                <md-icon md-font-library="material-icons">local_shipping</md-icon> Attached Assets
              </md-button>
          </md-list-item>
        </md-list>

      </md-sidenav>

      <div flex layout="column" tabIndex="-1" role="main" class="md-whiteframe-z2">

        {!{if .IsAdmin}!}
        <md-content class="content" flex ng-show="main.focus == 'summary'">
          <h2>Summary</h2>
          <p>server uptime, resources, running plugins etc</p>
          <p><b>Is Admin: </b>{!{.IsAdmin}!}</p>
          <p><b>Username: </b>{!{.User.Username}!}</p>
          <p><b>First Name: </b>{!{.User.Firstname}!}</p>
          <p><b>Last Name: </b>{!{.User.Lastname}!}</p>
        </md-content>

        {!{template "userpage"}!}

        {!{template "usercreateeditpage"}!}

        <md-content class="content" flex ng-show="main.focus == 'data'">
          <h2>Data</h2>
          <p>all custom datasets and active streams will go here.</p>
        </md-content>
        {!{end}!}

      </div>
    </div>

    {!{template "tailcontent"}!}
    <script src="/static/js/app/mainController.js"></script>
    <script src="/static/js/app/userController.js"></script>
    <script src="/static/js/app/usereditController.js"></script>
  </body>
</html>
