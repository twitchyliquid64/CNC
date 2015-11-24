<md-content class="content" flex ng-show="main.focus == 'plugin-new'">
  <md-data-table-toolbar>
    <h2 class="md-title" flex="50">Plugins <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon> New</h2>

    <div class="md-toolbar-tools">
      <span flex></span>
    </div>

  </md-data-table-toolbar>


  <md-content>
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

  <button flex="" layout-fill=""
  class="md-raised md-primary md-button md-scope"
  ng-click="process()"
  aria-label="Create Plugin" tabindex="0" aria-disabled="true">
   <i class="material-icons" style="vertical-align: middle;">save</i>
  <span style="vertical-align: middle;">Create</span>
  </button>

</md-content>
