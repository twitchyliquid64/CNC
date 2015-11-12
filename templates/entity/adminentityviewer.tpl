<md-content class="content" flex ng-show="main.focus == 'entities'">
  <md-data-table-toolbar>
    <h2 class="md-title">Entities</h2>

    <div class="md-toolbar-tools">
      <span flex></span>
      <md-button class="ng-icon-button" ng-click="refresh()" aria-label="Refresh">
        <md-icon md-font-library="material-icons">refresh</md-icon>
      </md-button>
      <md-button class="ng-icon-button" ng-click="" aria-label="Add Entity">
        <md-icon md-font-library="material-icons">add</md-icon>
      </md-button>
    </div>
  </md-data-table-toolbar>

  <div layout="row" layout-sm="column" layout-align="space-around" ng-show="showLoading">
    <md-progress-circular md-mode="indeterminate"></md-progress-circular>
  </div>
  <style>
  .small-icons {
    min-width: 16px;
    padding: 2px;
    margin: 2px;
  }
  </style>

  <md-data-table-container ng-hide="showLoading">
    <table md-data-table md-row-select="selected" md-progress="deferred">
      <thead>
        <tr>
          <th name="Name"></th>
          <th name="Last Updated"></th>
          <th name="Status"></th>
          <th name="Actions"></th>
        </tr>
      </thead>
      <tbody>
        <tr ng-repeat="entity in entities">
          <td>{{entity.Name}}</td>
          <td></td>
          <td></td>
          <td>
          </td>
        </tr>
      </tbody>
    </table>
  </md-data-table-container>
</md-content>
