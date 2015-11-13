<md-content class="content" flex ng-show="main.focus == 'user-edit'" ng-controller="usereditController as userC">
  <md-data-table-toolbar>
    <h2 class="md-title" flex="50">Users <md-icon md-font-library="material-icons">keyboard_arrow_right</md-icon> <span ng-show="isNewUserMode">New</span><span ng-hide="isNewUserMode">Edit</span></h2>

    <div class="md-toolbar-tools">
      <span flex></span>

      <md-button class="ng-icon-button" ng-disabled="!user.Username" ng-hide="isNewUserMode" ng-click="userC.editPermsUser(user.Username);main.activate('user-permissions')" aria-label="Edit Permissions">
        <md-icon md-font-library="material-icons">security</md-icon>
      </md-button>
    </div>

  </md-data-table-toolbar>


  <p>User Details</p>
  <md-input-container flex layout-fill>
    <label>Username</label>
    <input ng-model="user.Username" type="text">
  </md-input-container>

  <md-content layout="row" layout-sm="column">
    <md-input-container flex>
      <label>First Name</label>
      <input ng-model="user.Firstname" type="text">
    </md-input-container>
    <md-input-container flex>
      <label>Last Name</label>
      <input ng-model="user.Lastname" type="text">
    </md-input-container>
  </md-content>

  <p>Contact Details</p>
  <md-content layout="row" layout-sm="column" flex>
    <md-input-container flex="60">
      <label>Email</label>
      <input ng-model="user.MainEmail.Address" type="email">
    </md-input-container>
    <md-input-container flex>
      <label>Mobile Phone</label>
      <input ng-model="user.Mobile" type="phone">
    </md-input-container>
  </md-content>

  <p>Residential Address</p>
  <md-input-container flex>
    <label>Address line 1</label>
    <input ng-model="user.MainAddress.Address1">
  </md-input-container>
  <md-input-container md-no-float>
    <label>Address line 2</label>
    <input ng-model="user.MainAddress.Address2">
  </md-input-container>
  <div layout layout-sm="column">
    <md-input-container flex>
      <label>City</label>
      <input ng-model="user.MainAddress.City">
    </md-input-container>
    <md-input-container flex>
      <label>State</label>
      <input ng-model="user.MainAddress.State">
    </md-input-container>
    <md-input-container flex>
      <label>Postal Code</label>
      <input ng-model="user.MainAddress.Postcode" type="number">
    </md-input-container>
  </div>

  <md-content ng-show="isNewUserMode">
    <p>Default Authentication</p>
    <md-input-container flex>
      <label>Password</label>
      <input ng-model="user.Password" type="password">
    </md-input-container>
  </md-content>

  <button flex="" layout-fill=""
  class="md-raised md-primary md-button md-scope"
  ng-click="userC.process()"
  aria-label="Create User" tabindex="0" aria-disabled="true">
   <i class="material-icons" style="vertical-align: middle;">save</i>
  <span ng-show="isNewUserMode" style="vertical-align: middle;">Create</span>
  <span ng-hide="isNewUserMode" style="vertical-align: middle;">Update</span>
  </button>


</md-content>
