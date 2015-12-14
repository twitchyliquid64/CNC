(function () {
    angular.module('baseApp').directive('entityStatus', function(){
      return {
          restrict: 'A',
          template: '<span ng-hide="entity.LastStatString==\'\'" ng-class="{amber: entity.LastStatStyle==\'amber\', red: entity.LastStatStyle==\'red\', green: entity.LastStatStyle==\'green\'}">\
                      <span ng-show="entity.LastStatIcon!=\'\'"><md-icon md-font-library="material-icons">{{entity.LastStatIcon}}</md-icon></span>\
                      {{entity.LastStatString}}\
                      <md-progress-linear md-mode="determinate" ng-show="entity.LastStatStyle==\'progress-linear\'" value="{{entity.LastStatMeta}}">\
                     </span>\
                     <span ng-show="entity.LastStatString==\'\'">No status to show</span>\
          '
      };
  });
})();
