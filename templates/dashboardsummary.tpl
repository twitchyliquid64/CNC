<md-content class="content" flex ng-show="main.focus == 'summary'" layout="column" layout-fill>
  <md-content>
    <h2>Summary</h2>
  </md-content>

  <md-content flex layout="row" layout-fill layout-wrap>
    <md-content flex="99" flex-sm="100" layout="column" layout-fill style="min-height: 400px;">
      <md-subheader class="md-no-sticky" layout="column">
        <span flex layout-fill>System Log</span>
        <md-icon flex class="md-avatar" md-font-library="material-icons" ng-show="isConnected()">check</md-icon>
        <md-icon flex class="md-avatar" md-font-library="material-icons" ng-hide="isConnected()">close</md-icon>
      </md-subheader>

      <p class="logElement" layout="row" ng-repeat="msg in getLogMsgs()" style="margin: 0px;">
        <span class="logComponent">[{{msg.Component | uppercase}}] </span> {{msg.Message}}
      </p>
    </md-content>


    <md-content flex="35" flex-sm="100" layout="column">
      <md-list>
        <md-subheader class="md-no-sticky">System Status
          <md-icon flex class="md-avatar" md-font-library="material-icons" ng-show="updateState=='done'">check</md-icon>
          <md-icon flex class="md-avatar" md-font-library="material-icons" ng-show="updateState=='loading'">more_horiz</md-icon>
          <md-icon flex class="md-avatar" md-font-library="material-icons" ng-show="updateState=='error'">error</md-icon>
        </md-subheader>
        <md-list-item class="md-2-line" ng-repeat="component in components">
          <md-icon class="md-avatar" md-font-library="material-icons">{{component.Icon}}</md-icon>
          <div class="md-list-item-text">
            <h3>{{component.Name}}</h3>
            <p ng-class="{red: component.State=='Fault', green: component.State=='OK', amber: component.State=='Disabled'}">{{component.State}}</p>
          </div>
        </md-list-item>
      </md-list>
    </md-content>
  </md-content>

  <md-content flex layout="row">
    <md-button flex class="md-raised" ng-disabled="disableActions" ng-click="reloadTemplates();"><i class="material-icons" style="vertical-align: middle;">refresh</i> Reload Templates</md-button>
    <md-button flex class="md-raised" ng-disabled="disableActions" ><i class="material-icons" style="vertical-align: middle;">data_usage</i> Clean DB</md-button>
  </md-content>
</md-content>
