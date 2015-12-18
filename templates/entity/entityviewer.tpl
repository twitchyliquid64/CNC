<md-content class="content" flex ng-show="main.focus == 'entity-view'">
  <md-data-table-toolbar>
    <h2 class="md-title" flex="50" ng-click="main.activateRouted('/admin/entities', 'entities')">
      <md-icon md-font-library="material-icons">keyboard_arrow_left</md-icon>
      Entities <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon> {{entity.Name}}</h2>
    <div class="md-toolbar-tools">
      <span flex></span>
      <md-icon flex class="md-avatar" md-font-library="material-icons" ng-show="(!connected) && (!wasConnected)">more_horiz</md-icon>
      <md-icon flex class="md-avatar" md-font-library="material-icons" ng-show="(!connected) && wasConnected">close</md-icon>
      <md-icon flex class="md-avatar" md-font-library="material-icons" ng-show="connected">check</md-icon>
    </div>
  </md-data-table-toolbar>

  <div layout="row" layout-align="space-around" ng-show="showLoading">
    <md-progress-circular md-mode="indeterminate"></md-progress-circular>
  </div>

  <md-content flex layout="row" layout-sm="column" layout-fill layout-wrap ng-hide="showLoading">
    <md-content flex="60" flex-sm="100" layout="column">
      <md-list>
        <md-list-item class="md-2-line">
          <md-icon class="md-avatar" md-font-library="material-icons">http</md-icon>
          <div class="md-list-item-text">
            <h3>API Key</h3>
            <p>{{entity.APIKey}}</p>
          </div>
        </md-list-item>


        <md-list-item class="md-2-line">
          <md-icon class="md-avatar" md-font-library="material-icons">access_time</md-icon>
          <div class="md-list-item-text">
            <h3>Lasted Updated</h3>
            <p>{{entity.UpdatedAt_time}}</p>
          </div>
        </md-list-item>

        <md-list-item class="md-2-line">
          <md-icon class="md-avatar" md-font-library="material-icons">assessment</md-icon>
          <div class="md-list-item-text">
            <h3>Current Status</h3>
            <p entity-status></p>
          </div>
        </md-list-item>
      </md-list>
    </md-content>

    <md-content flex layout="column" layout-fill layout-align="center center" style="color: rgba(0, 0, 0, 0.54);">
      <i class="material-icons" style="vertical-align: middle;font-size: 450%;">{{entity.Icon}}</i>
      <p><b><md-icon class="md-avatar" md-font-library="material-icons">group_work</md-icon>Category: </b>{{entity.Category}}<br>
      <b><md-icon class="md-avatar" md-font-library="material-icons">access_time</md-icon>Created: </b>{{entity.CreatedAt_time}}</p>
    </md-content>

  </md-content>

  <hr ng-hide="showLoading" flex/>

  <md-content layout="column" layout-fill layout-wrap ng-hide="showLoading">
    <h3>Activity Log</h3>

    <p class="logElement" layout="row" ng-repeat="msg in msgs" style="margin: 0px;">
      <span class="logComponent">[{{msg.Created | uppercase}}] </span> {{msg.Content}}
    </p>
  </md-content>
</md-content>
