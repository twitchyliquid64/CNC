<md-content class="content" flex ng-show="main.focus == 'summary'" layout="column" layout-fill>
  <md-content>
    <h2>Summary</h2>
  </md-content>

  <md-content flex layout="row" layout-fill>
    <md-content flex="99" layout="column" layout-fill>
      <md-subheader class="md-no-sticky">System Log</md-subheader>

      <div class="logElement" layout="row">
        <span class="logComponent">MESSENGER</span>
        <span >Initialisation completed</span>
      </div>
      <div class="logElement" layout="row">
        <span class="logComponent">WEB-GATEWAY</span>
        <span >Request with invalid route dropped</span>
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
