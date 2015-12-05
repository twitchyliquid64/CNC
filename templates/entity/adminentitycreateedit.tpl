<md-content class="content" flex ng-show="main.focus == 'entity-edit'">
  <md-data-table-toolbar>
    <h2 class="md-title" flex="50" ng-click="main.activateRouted('/admin/entities', 'entities')">
      <md-icon md-font-library="material-icons">keyboard_arrow_left</md-icon>
      Entities <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon> <span ng-show="isNewEntityMode">New</span><span ng-hide="isNewEntityMode">Edit</span></h2>

    <div class="md-toolbar-tools">
      <span flex></span>
    </div>

  </md-data-table-toolbar>


  <p>Entity Details</p>
  <md-input-container flex layout-fill>
    <label>Name</label>
    <input ng-model="entity.Name" type="text">
  </md-input-container>

  <md-input-container flex  layout-fill>
    <label>Category</label>
    <input ng-model="entity.Category" type="text">
  </md-input-container>

  <md-input-container flex>
    <md-autocomplete flex required
      md-input-name="autocompleteField"
      md-no-cache="true"
      md-items="icon in icons"
      md-item-text="entity.Icon"
      md-search-text="entity.Icon"
      md-selected-item="entity.Icon"
      md-floating-label="Icon">
      <md-item-template>
        <i class="material-icons" style="vertical-align: middle;">{{icon}}</i> {{icon}}
      </md-item-template>
    </md-autocomplete>
  </md-input-container>

  <button flex="" layout-fill=""
  class="md-raised md-primary md-button md-scope"
  ng-click="process()"
  aria-label="Create Entity" tabindex="0" aria-disabled="true">
   <i class="material-icons" style="vertical-align: middle;">save</i>
  <span ng-show="isNewEntityMode" style="vertical-align: middle;">Create</span>
  <span ng-hide="isNewEntityMode" style="vertical-align: middle;">Save</span>
  </button>

</md-content>
