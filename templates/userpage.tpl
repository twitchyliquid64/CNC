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
          <th name="Name" order-by="Firstname"></th>
          <th name="Username" order-by="Username"></th>
          <th name="Permissions"></th>
          <th name="Email" order-by="Email"></th>
        </tr>
      </thead>
      <tbody>
        <tr md-auto-select ng-repeat="user in users">
          <td>{{user.Firstname}} {{user.Lastname}}</td>
          <td>{{user.Username}}</td>
          <td><md-chips ng-model="user.perms" readonly="true"></md-chips></td>
          <td>{{user.MailEmail.Address}}</td>
        </tr>
      </tbody>
    </table>
  </md-data-table-container>
</md-content>
