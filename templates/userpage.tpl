<md-content class="content" flex ng-show="main.focus == 'users'" ng-controller="userController as userC">
  <md-data-table-toolbar>
    <h2 class="md-title">Users</h2>

    <div class="md-toolbar-tools">
      <span flex></span>

      <md-button class="ng-icon-button" ng-click="userC.newUser();main.activate('user-edit')" aria-label="Add User">
        <md-icon md-font-library="material-icons">person_add</md-icon>
      </md-button>
    </div>
  </md-data-table-toolbar>

  <div layout="row" layout-sm="column" layout-align="space-around" ng-show="showLoading">
    <md-progress-circular md-mode="indeterminate"></md-progress-circular>
  </div>

  <md-data-table-container ng-hide="showLoading">
    <table md-data-table md-row-select="selected" md-progress="deferred">
      <thead>
        <tr>
          <th name="Username"></th>
          <th name="Name"></th>
          <th name="Actions"></th>
        </tr>
      </thead>
      <tbody>
        <tr ng-repeat="user in users">
          <td>{{user.Username}}</td>
          <td>{{user.Firstname}} {{user.Lastname}}</td>
          <td>
            <md-button class="ng-icon-button" ng-click="userC.editUser(user.Username);main.activate('user-edit')" aria-label="Edit User">
              <md-icon md-font-library="material-icons">mode_edit</md-icon>
            </md-button>
            <md-button class="ng-icon-button" ng-click="userC.deleteUser(user.Username)" aria-label="Delete User">
              <md-icon md-font-library="material-icons">delete</md-icon>
            </md-button>
          </td>
        </tr>
      </tbody>
    </table>
  </md-data-table-container>
</md-content>
