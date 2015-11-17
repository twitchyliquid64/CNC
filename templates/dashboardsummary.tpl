<md-content class="content" flex ng-show="main.focus == 'summary'" layout="column" layout-fill>
  <md-content>
    <h2>Summary</h2>
  </md-content>

  <md-content flex layout="row" layout-fill>
    <md-content flex="99" layout="column" layout-fill>
      <md-subheader class="md-no-sticky" layout="column">
        <span flex layout-fill>System Log</span>
        <md-icon flex class="md-avatar" md-font-library="material-icons" ng-show="isConnected()">check</md-icon>
        <md-icon flex class="md-avatar" md-font-library="material-icons" ng-hide="isConnected()">close</md-icon>
      </md-subheader>

      <div class="logElement" layout="row" ng-repeat="msg in getLogMsgs()">
        <span class="logComponent">[{{msg.Component | uppercase}}]</span>
        <span >{{msg.Message}}</span>
      </div>
    </md-content>


    <md-content flex="40" layout="column">
      <md-list>
        <md-subheader class="md-no-sticky">System Status</md-subheader>
        <md-list-item class="md-2-line">
          <md-icon class="md-avatar" md-font-library="material-icons">people</md-icon>
          <div class="md-list-item-text">
            <h3>Database</h3>
            <p style="color: red;">TRACKING NOT IMPLEMENTED</p>
          </div>
        </md-list-item>
      </md-list>
    </md-content>
  </md-content>
</md-content>
