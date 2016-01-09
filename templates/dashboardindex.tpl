<!DOCTYPE html>
<html lang="en">

  <head>
      <title>CNC Dashboard</title>
      {!{template "headcontent"}!}
  </head>

  <body layout="column" ng-app="baseApp" ng-controller="mainController as main" ng-cloak>
      <md-toolbar layout="row" flex="10">
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

          <md-subheader class="md-no-sticky">Main Menu</md-subheader>
          <md-list-item>
              <md-button ng-click="main.activate('messenger')" ng-disabled="true">
                <md-icon md-font-library="material-icons">message</md-icon> Messenger
              </md-button>
          </md-list-item>
          <md-list-item>
              <md-button ng-click="main.activate('mail')" ng-disabled="true">
                <md-icon md-font-library="material-icons">email</md-icon> Mail
              </md-button>
          </md-list-item>
          <md-list-item>
              <md-button ng-click="main.activate('sonder')">
                <md-icon md-font-library="material-icons">wallpaper</md-icon> End of Sonder
              </md-button>
          </md-list-item>

          {!{if .IsAdmin}!}
          <md-subheader class="md-no-sticky">Admin</md-subheader>
          <md-list-item>
              <md-button ng-click="main.activateRouted('/admin/dashboard', 'summary')">
                <md-icon md-font-library="material-icons">tune</md-icon> Summary
              </md-button>
          </md-list-item>
          <md-list-item>
              <md-button ng-click="main.activateRouted('/admin/users', 'users')">
                <md-icon md-font-library="material-icons">people</md-icon> Users
              </md-button>
          </md-list-item>
          <md-list-item>
              <md-button ng-click="main.activateRouted('/admin/entities', 'entities')">
                <md-icon md-font-library="material-icons">settings_input_antenna</md-icon> Entities
              </md-button>
          </md-list-item>
          <md-list-item>
              <md-button ng-click="main.activateRouted('/admin/plugins', 'plugins');">
                <md-icon md-font-library="material-icons">memory</md-icon> Plugins
              </md-button>
          </md-list-item>
          {!{end}!}

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
        {!{template "usercreateeditpage"}!}

        {!{template "userpermissions"}!}
        {!{end}!}


        {!{if eq .User.Username  "testadmin"}!}
        <md-content class="content" flex ng-show="main.focus == 'sonder'" ng-controller="sonderController as sonder">
          <h2>End of Sonder</h2>
          <p>This is your contract with yourself. If you have not engaged yourself in the area which your mind revolves around, by the 20th November 2016, you must aggressively pursue personal change to these ends
          to the detriment of other pursuits. This is not the start of something, it is merely a safety net - to avoid an end.</p>

          <div flex style="text-align: center;" layout-align="center">
            <h3 flex>{{sonder.getDays()}} days</h3>
            <p flex>{{sonder.getHours()}} hours, {{sonder.getMinutes()}} minutes</p>
          </div>

          <md-list>
            <md-subheader class="md-no-sticky">Motivations
            </md-subheader>
            <md-list-item class="md-2-line">
              <md-icon class="md-avatar" md-font-library="material-icons">cross</md-icon>
              <div class="md-list-item-text">
                <h3>Avoidance, Distration</h3>
                <p>You center your whole life around this one thing, yet you direct all your efforts and thoughts elsewhere, particularly your hobbies (distractions).</p>
              </div>
            </md-list-item>
          </md-list>

        </md-content>
        {!{end}!}

        <md-content flex ng-show="main.isRoutingMode" layout-fill>
          <div ng-view layout-fill></div>
        </md-content>

      </div>
    </div>

    {!{template "tailcontent"}!}
    <script src="/static/js/app/mainController.js"></script>
    <script src="/static/js/app/summaryController.js"></script>
    <script src="/static/js/app/sonderController.js"></script>
    <script src="/static/js/app/user/userController.js"></script>
    <script src="/static/js/app/user/usereditController.js"></script>
    <script src="/static/js/app/user/userpermissionController.js"></script>
    <script src="/static/js/app/entity/entityViewerAdminController.js"></script>
    <script src="/static/js/app/entity/entityViewerController.js"></script>
    <script src="/static/js/app/entity/entityFormController.js"></script>
    <script src="/static/js/app/entity/entityStatusDirective.js"></script>
    <script src="/static/js/app/plugin/pluginListController.js"></script>
    <script src="/static/js/app/plugin/pluginCreateController.js"></script>
    <script src="/static/js/app/plugin/pluginEditController.js"></script>
    <script src="/static/js/app/plugin/resourceCreateEditController.js"></script>
    <script src="/static/js/app/services/loggerService.js"></script>
    <script src="/static/js/ace/src-min/ace.js" type="text/javascript" charset="utf-8"></script>

    <script type="text/javascript" src="https://maps.google.com/maps/api/js?sensor=false"></script>
    <script src="/static/js/app/entity/entityMapController.js"></script>

  </body>
</html>
