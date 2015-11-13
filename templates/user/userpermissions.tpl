<md-content class="content" flex ng-show="main.focus == 'user-permissions'" ng-controller="userpermissionControlller as userC">
  <md-data-table-toolbar>
    <h2 class="md-title">Users <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon> Edit <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon> Permissions</h2>
  </md-data-table-toolbar>

  <div layout="row" layout-sm="column" layout-align="space-around" ng-show="showLoading">
    <md-progress-circular md-mode="indeterminate"></md-progress-circular>
  </div>


  <md-list ng-hide="showLoading">
    <md-subheader>User: {{user.Firstname}} {{user.Lastname}}({{user.Username}})</md-subheader>
    <md-list-item class="md-2-line" ng-repeat="item in currentPermissionsDisplay">
      <md-icon md-font-library="material-icons">vpn_key</md-icon>
      <div class="md-list-item-text" ng-class="md-offset">
        <h3> {{item.key}}</h3>
        <p>Permission<md-button class="md-icon-button md-accent" aria-label="Delete" ng-click="userC.delPerm(item.key)">
          <md-icon md-font-library="material-icons">close</md-icon>
        </md-button></p>
      </div>
    </md-list-item>
    <md-subheader class="md-no-sticky" ng-hide="(currentPermissionsDisplay.length > 0)">This user has no permissions.</md-subheader>
  </md-list>

  <md-autocomplete
      ng-hide="showLoading"
      ng-disabled="false"
      md-no-cache="false"
      md-selected-item="userC.selectedItem"
      md-search-text="userC.searchText"
      md-items="item in userC.recommendedPermissions"
      md-selected-item-change="userC.addPerm(item.key)"
      md-item-text="item.key"
      md-min-length="0"
      placeholder=" + Add a permission">
    <md-item-template>
      <md-icon md-font-library="material-icons">vpn_key</md-icon>
      <span md-highlight-text="userC.searchText" md-highlight-flags="^i">{{item.display}}</span>
    </md-item-template>
    <md-not-found>
      No matches found for "{{userC.searchText}}".
    </md-not-found>
  </md-autocomplete>

</md-content>
