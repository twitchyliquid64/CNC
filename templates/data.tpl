<md-content class="content" flex ng-show="main.focus == 'data'">
  <md-data-table-toolbar>
    <h2 class="md-title" flex="50"> Database</h2>
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

  <md-content flex layout="column" layout-fill layout-wrap ng-hide="showLoading">
    <md-content flex="50" layout="column">
      <p>Please enter your SQL query in the box below, and press enter.</p>
      <input type="text" ng-model="SQL" my-enter="do()" layout-fill flex></input>
    </md-content>
    <md-content flex="85" layout="column" layout-align="center center" md-padding layout-fill>
      <style>
      tr,td {
        border: 1px solid black;
      }
      </style>


      <table style="margin-top: 12px; width: 80%;	border-collapse: collapse;" ng-hide="tableHide">
        <thead style="font-weight: bold;">
          <tr>
            <td ng-repeat="col in tableCols track by $index">{{col}}</td>
          </tr>
        </thead>
        <tbody>
          <tr ng-repeat="row in tableRows track by $index">
            <td ng-repeat="element in row track by $index">{{element}}</td>
          </tr>
        </tbody>
      </table>

      <div ng-show="err != ''">
        {{err}}
      </div>
    </md-content>
  </md-content>

</md-content>
