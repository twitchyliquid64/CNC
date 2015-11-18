<md-content class="content" flex ng-show="main.focus == 'summary'" layout="column" layout-fill>
  <md-content>
    <h2>Summary</h2>
  </md-content>

  <md-content flex layout="row" layout-fill layout-wrap>
    <md-content flex="99" layout="column" layout-fill>
      <md-subheader class="md-no-sticky" layout="column">
        <span flex layout-fill>System Log</span>
        <md-icon flex class="md-avatar" md-font-library="material-icons" ng-show="isConnected()">check</md-icon>
        <md-icon flex class="md-avatar" md-font-library="material-icons" ng-hide="isConnected()">close</md-icon>
      </md-subheader>

      <div class="logElement" layout="row" ng-repeat="msg in getLogMsgs()">
        <span class="logComponent">[{{msg.Component | uppercase}}]</span>{{msg.Message}}
      </div>
    </md-content>


    <md-content flex="40" layout="column">
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
</md-content>
