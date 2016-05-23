<md-content class="content" flex ng-show="main.focus == 'resource-form'">
  <md-data-table-toolbar>
    <h2 class="md-title" ng-click="main.activateRouted('/admin/plugin/'+resource.PluginID, 'plugin-edit')" flex="50">
                                  <md-icon md-font-library="material-icons">keyboard_arrow_left</md-icon>
                                  Plugins <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon>
                                  Resources <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon>
                                  Edit</h2>

    <div class="md-toolbar-tools">
      <span flex></span>
    </div>

  </md-data-table-toolbar>

  <style>

  .editor {
    height: 80%;
    min-height: 420px;
  }

  .repeated-item {
    border-right: 1px solid #ddd;
    box-sizing: border-box;
    display: table-cell;
    height: 50px;
    text-align: center;
  }

  #horizontal-container {
    display: block;
    height: 60px;
    overflow-x: scroll;
  }

</style>

  <div layout="row" flex layout-fill>
    <md-input-container flex>
      <label>Name</label>
      <input ng-model="resource.Name" type="text">
    </md-input-container>

    <md-input-container style="margin-top: 13px">
      <md-select ng-model="resource.ResType" ng-change="setMode(resource.ResType)">
        <md-option ng-value="type.code" ng-repeat="type in resTypes">{{type.name}}</md-option>
      </md-select>
    </md-input-container>
  </div>

  <div id="text-window">
    <div layout="row" flex>
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
    </div>

    <md-input-container layout="row" flex>
      <pre id="code-editor" class="editor" flex>console.log("Hello world!");</pre>
    </md-input-container>
  </div>

  <div id="graph-window">
    <div id="graph-editor" class="editor" layout="row" flex layout-fill>
    </div>
    <md-virtual-repeat-container id="horizontal-container" md-orient-horizontal>
      <div ng-repeat="block in blocks" class="repeated-item" flex>
        <md-button ng-click="codeGraph.addCode(block)" class="md-raised">{{block.name}}</md-button>
      </div>
    </md-virtual-repeat-container>
  </div>

  <button flex layout-fill layout="row"
  class="md-raised md-primary md-button md-scope"
  ng-click="process()"
  aria-label="Save Plugin" tabindex="0" aria-disabled="true">
   <i class="material-icons" style="vertical-align: middle;">save</i>
   <span style="vertical-align: middle;"> Save Changes</span>
  </button>

</md-content>
