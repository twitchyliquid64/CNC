<!DOCTYPE html>
<html lang="en">

  <head>
      <title>CNC Dashboard</title>
      {!{template "headcontent"}!}
  </head>

  <body layout="column" ng-app="baseApp">
      <md-toolbar layout="row" flex="5" class="md-whiteframe-z1">
        <h1>CNC</h1>
      </md-toolbar>

    <div layout="row" flex>
      <md-sidenav class="site-sidenav md-sidenav-left md-whiteframe-z2"
                    md-component-id="left"
                    md-is-locked-open="$mdMedia('gt-sm')">

        <md-list><!--Put ng-repeat in the md-list -->
          <md-list-item>
              <md-button>
                <md-icon></md-icon>
                LOLCAKES
              </md-button>
          </md-list-item>
        </md-list>

      </md-sidenav>
    </div>

    {!{template "tailcontent"}!}
  </body>
</html>
