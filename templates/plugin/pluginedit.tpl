<md-content class="content" flex ng-show="main.focus == 'plugin-edit'">
  <md-data-table-toolbar>
    <h2 class="md-title" flex="50">Plugins <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon> Edit</h2>

    <div class="md-toolbar-tools">
      <span flex></span>
    </div>

  </md-data-table-toolbar>

  <style>
  .small-icons {
    min-width: 16px;
    padding: 2px;
    margin: 2px;
  }
  </style>

  <md-content flex layout="row" layout-fill layout-wrap>
    <md-content flex="99" flex-sm="100" layout="column" layout-fill >
      <p>Plugin Details</p>
      <md-input-container flex layout-fill>
        <label>Name</label>
        <input ng-model="plugin.Name" type="text">
      </md-input-container>

      <md-input-container flex>
        <md-autocomplete flex required
          md-input-name="autocompleteField"
          md-no-cache="true"
          md-items="icon in icons"
          md-item-text="plugin.Icon"
          md-search-text="plugin.Icon"
          md-selected-item="plugin.Icon"
          md-floating-label="Icon">
          <md-item-template>
            <i class="material-icons" style="vertical-align: middle;">{{icon}}</i> {{icon}}
          </md-item-template>
        </md-autocomplete>
      </md-input-container>

      <md-input-container class="md-block">
            <label>Description</label>
            <textarea ng-model="plugin.Description" columns="1" md-maxlength="150" rows="5"></textarea>
      </md-input-container>

      <md-switch ng-model="plugin.Enabled" aria-label="Enabled" class="md-block">
        Enabled
      </md-switch>
    </md-content>

    <md-content md-margin flex="42" flex-sm="100" layout="column">
      <md-list>
        <p>Resources</p>
        <md-list-item class="md-2-line" ng-repeat="resource in plugin.Resources">
          <md-icon class="md-avatar" md-font-library="material-icons">code</md-icon>
          <div class="md-list-item-text">
            <h3>resource.Name</h3>
          </div>
        </md-list-item>
        <div ng-if="plugin.Resources.length == 0" >
          <span ><i>This plugin has no resources.</i></span>
        </div>
      </md-list>
    </md-content>
  </md-content>

  <button flex="" layout-fill=""
  class="md-raised md-primary md-button md-scope"
  ng-click="process()"
  aria-label="Save Plugin" tabindex="0" aria-disabled="true">
   <i class="material-icons" style="vertical-align: middle;">save</i>
  <span style="vertical-align: middle;"> Save Changes</span>
  </button>

</md-content>
