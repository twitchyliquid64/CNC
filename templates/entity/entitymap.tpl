<md-content class="content" flex ng-show="main.focus == 'entity-map'">
  <md-data-table-toolbar>
    <h2 class="md-title" flex="50" ng-click="main.activateRouted('/admin/entities', 'entities')">
      <md-icon md-font-library="material-icons">keyboard_arrow_left</md-icon>
      Entities <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon> {{entity.Name}}</h2>
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

    <style>
     @media screen and (min-width: 200px) {
         #map_canvas {
             margin: 0 auto;
             height: 650px;
             min-width: 150px;
             max-width: 2250px;
             float: left;
             width: 100%;
         }
     }
  </style>

  <div id="map_canvas"></div>
  </md-content>
</md-content>
