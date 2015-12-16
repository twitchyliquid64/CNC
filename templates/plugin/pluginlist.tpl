<md-content class="content" flex ng-show="main.focus == 'plugins'">
  <md-data-table-toolbar>
    <h2 class="md-title">Plugins</h2>

    <div class="md-toolbar-tools">
      <span flex></span>
      <md-button class="ng-icon-button" ng-click="refresh()" aria-label="Refresh">
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

    <md-card ng-repeat="plugin in plugins">
      <md-card-content>
        <md-data-table-toolbar>
          <h2 flex="50" class="md-title"><md-icon md-font-library="material-icons">{{plugin.Icon}}</md-icon> {{plugin.Name}}</h2>
          <div class="md-toolbar-tools">
            <span flex hide-sm></span>
            <md-switch ng-model="plugin.Enabled" ng-change="switchChanged(plugin.ID, plugin.Enabled)" aria-label="Enable" class="md-block"></md-switch>
          </div>
        </md-data-table-toolbar>

      </md-card-content>
      <md-card-actions layout="row" layout-align="end center">
        <span ng-if="!plugin.HasCrashed">
          <p hide-sm layout-margin ng-show="plugin.Enabled" layout-padding class="green">Running</p>
          <p hide-sm layout-margin ng-hide="plugin.Enabled" layout-padding class="amber">Disabled</p>
        </span>
        <span ng-if="plugin.HasCrashed">
          <p hide-sm layout-margin  layout-padding class="red">{{plugin.ErrorStr}}</p>
        </span>
        <span flex hide-sm>
        </span>
        <md-button ng-disabled="plugin.Enabled" ng-click="deletePlugin(plugin.ID, plugin.Name, $event)"><md-icon md-font-library="material-icons">close</md-icon> Delete</md-button>
        <md-button ng-click="main.activateRouted('/admin/plugin/'+plugin.ID, 'plugin-edit')"><md-icon md-font-library="material-icons">edit</md-icon> Edit</md-button>
      </md-card-actions>
    </md-card>

  </md-content>
</md-content>
