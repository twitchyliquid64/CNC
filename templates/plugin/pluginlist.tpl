<md-content class="content" flex ng-show="main.focus == 'plugins'">
  <md-data-table-toolbar>
    <h2 class="md-title">Plugins</h2>

    <div class="md-toolbar-tools">
      <span flex></span>
      <md-button class="ng-icon-button" ng-click="" aria-label="Refresh">
        <md-icon md-font-library="material-icons">refresh</md-icon>
      </md-button>
      <md-button class="ng-icon-button" ng-click="main.activateRouted('/admin/plugins/new', 'plugin-new')" aria-label="Add Plugin">
        <md-icon md-font-library="material-icons">add</md-icon>
      </md-button>
    </div>
  </md-data-table-toolbar>

  <div layout="row" layout-align="space-around" ng-show="showLoading">
    <md-progress-circular md-mode="indeterminate"></md-progress-circular>
  </div>
  <style>
  .small-icons {
    min-width: 16px;
    padding: 2px;
    margin: 2px;
  }
  </style>

  <md-content ng-hide="showLoading">

    <md-card>
      <md-card-content>
        <md-data-table-toolbar>
          <h2 flex="50" class="md-title"><md-icon md-font-library="material-icons">cloud</md-icon> Weather Alerter</h2>
          <div class="md-toolbar-tools">
            <span flex hide-sm></span>
            <md-switch ng-model="blue" aria-label="Enable" class="md-block"></md-switch>
          </div>
        </md-data-table-toolbar>

      </md-card-content>
      <md-card-actions layout="row" layout-align="end center">
        <p layout-margin layout-padding class="green">Running</p>
        <span flex hide-sm>
        </span>
        <md-button><md-icon md-font-library="material-icons">close</md-icon> Stop</md-button>
        <md-button><md-icon md-font-library="material-icons">edit</md-icon> Edit</md-button>
      </md-card-actions>
    </md-card>

  </md-content>
</md-content>
