<md-content class="content" flex ng-show="main.focus == 'resource-form'">
  <md-data-table-toolbar>
    <h2 class="md-title" flex="50">Plugins <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon>
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
    min-height: 270px;
  }

  </style>


  <md-input-container flex layout-fill>
    <label>Name</label>
    <input ng-model="resource.Name" type="text">
  </md-input-container>


  <md-input-container flex layout="row" class="md-block">
    <div flex="50">
      <md-checkbox ng-model="resource.IsExecutable" aria-label="Checkbox 1">
        Javascript Code
      </md-checkbox>
    </div>
    <div flex="50">
      <md-checkbox ng-model="resource.IsTemplate" aria-label="Checkbox 1">
        Template
      </md-checkbox>
    </div>
  </md-input-container>

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
