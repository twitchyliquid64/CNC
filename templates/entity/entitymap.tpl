<div class="content" flex ng-show="main.focus == 'entity-map'">
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

  <div flex layout="column" layout-wrap ng-hide="showLoading">
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

    <div flex layout="row" layout-wrap style="padding-left: 3px; margin-left: 14px;">
      <div flex="50" layout="column">
        <div style="color: rgba(0, 0, 0, 0.54); font-size: 12px;">Speed</div>
        <div style="color: rgba(0, 0, 0, 1.00); font-size: 20px;">{{locs[0].SpeedKph}}</div>
      </div>


      <div flex="20" flex-sm="50" layout="column">
        <div style="color: rgba(0, 0, 0, 0.54); font-size: 12px;">Latitude</div>
        <div style="color: rgba(0, 0, 0, 1.00); font-size: 20px;white-space: nowrap;overflow: hidden;text-overflow: ellipsis;">{{locs[0].Latitude}}</div>
      </div>

      <div flex="20" flex-sm="50" layout="column">
        <div style="color: rgba(0, 0, 0, 0.54); font-size: 12px;">Longitude</div>
        <div style="color: rgba(0, 0, 0, 1.00); font-size: 20px;white-space: nowrap;overflow: hidden;text-overflow: ellipsis;">{{locs[0].Longitude}}</div>
      </div>

      <div flex="10" flex-sm="100" layout="column">
        <div style="color: rgba(0, 0, 0, 0.54); font-size: 12px;">Accuracy</div>
        <div style="color: rgba(0, 0, 0, 1.00); font-size: 20px;">{{locs[0].AccuracyDisplay}}</div>
      </div>

      <div flex="25" layout="column">
        <div style="color: rgba(0, 0, 0, 0.54); font-size: 12px;">Heading</div>
        <div style="color: rgba(0, 0, 0, 1.00); font-size: 20px;">{{locs[0].Course}}</div>
      </div>

      <div flex="25" layout="column">
        <div style="color: rgba(0, 0, 0, 0.54); font-size: 12px;">Number of Satellites</div>
        <div style="color: rgba(0, 0, 0, 1.00); font-size: 20px;">{{locs[0].SatNum}}</div>
      </div>


      <div flex="50" layout="column">
        <div style="color: rgba(0, 0, 0, 0.54); font-size: 12px;">Last location update</div>
        <div style="color: rgba(0, 0, 0, 1.00); font-size: 20px;">{{locs[0].TimeUpdatedString}}</div>
      </div>

      <div flex="100" layout="column">
        <div style="color: rgba(0, 0, 0, 0.54); font-size: 12px;">Historical</div>
        <md-slider flex min="0" max="800" ng-model="historicalShown" ng-disabled="loadingData" aria-label="historicalShown">
        </md-slider>
      </div>

    </div>

  </div>
</div>
