<md-content class="content" flex ng-show="main.focus == 'resource-form'">
  <md-data-table-toolbar>
    <h2 class="md-title" ng-click="main.activateRouted('/admin/plugin/'+resource.PluginID, 'plugin-edit')" flex="50">
                                  <md-icon md-font-library="material-icons">keyboard_arrow_left</md-icon>
                                  Plugins <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon>
                                  Resources <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon>
                                  <span ng-show="isCreateMode">New</span>
                                  <span ng-hide="isCreateMode">Edit</span></h2>

    <div class="md-toolbar-tools">
      <span flex></span>
    </div>

  </md-data-table-toolbar>

  <style>
  #editor {
    width: 95%;
    height: 80%;
    min-height: 420px;
  }

</style>

  <md-input-container flex layout-fill>
    <label>Name</label>
    <input ng-model="resource.Name" type="text">
  </md-input-container>

  <md-input-container>
    <md-select ng-model="resource.ResType" aria-label="Resource Type">
      <md-option ng-value="type.code" ng-repeat="type in resTypes">{{type.name}}</md-option>
    </md-select>
  </md-input-container>

  <button flex=""
  class="md-primary md-button md-scope"
  ng-click="showReference()"
  aria-label="View API Reference" tabindex="0" aria-disabled="true">
   <i class="material-icons" style="vertical-align: middle;">code</i>
   <span style="vertical-align: middle;"> API Reference</span>
  </button>

  <button flex=""
  class="md-primary md-button md-scope"
  ng-click="toggleMode()"
  aria-label="Toggle editor mode" tabindex="0" aria-disabled="true">
  <i class="material-icons" style="vertical-align: middle;" ng-show="mode == 'js'">code</i>
  <i class="material-icons" style="vertical-align: middle;" ng-show="mode == 'html'">web</i>
   <span style="vertical-align: middle;"> Toggle Editor mode</span>
  </button>


  <md-input-container class="md-block">
    <pre id="editor">console.log("Hello world!");</pre>
  </md-input-container>


  <button flex="" layout-fill=""
  class="md-raised md-primary md-button md-scope"
  ng-click="process()"
  aria-label="Save Plugin" tabindex="0" aria-disabled="true">
   <i class="material-icons" style="vertical-align: middle;">save</i>
   <span ng-show="isCreateMode" style="vertical-align: middle;"> Create</span>
   <span ng-hide="isCreateMode" style="vertical-align: middle;"> Save Changes</span>
  </button>

</md-content>
