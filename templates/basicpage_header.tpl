<!DOCTYPE html>
<html lang="en">

  <head>
      <title>{{.Title}}</title>
       <link href="/static/css/v-accordion.min.css" rel="stylesheet">
       <meta charset="utf-8">
       <meta http-equiv="X-UA-Compatible" content="IE=edge">
       <meta name="description" content="">
       <meta name="viewport" content="initial-scale=1, maximum-scale=1, width=device-width" />

       <link rel="stylesheet" href="/static/css/cnc.css">
       <link rel="stylesheet" href="/static/fonts/roboto.css">
       <link rel="stylesheet" href="/static/fonts/material-icons/materialicons.css">
       <link rel="stylesheet" href="/static/css/angular-material.css">
       <link rel="stylesheet" href="/static/css/md-data-table.min.css">
       <script src="/static/js/angular/angular.min.js"></script>


       <!-- Angular Material Dependencies - angular moved to head-->
       <script src="/static/js/angular/angular-animate.min.js"></script>
       <script src="/static/js/angular/angular-route.min.js"></script>
       <script src="/static/js/angular/angular-aria.min.js"></script>
       <script src="/static/js/angular/angular-messages.min.js"></script>
       <script src="/static/js/angular/angular-material.min.js"></script>
       <script src="/static/js/angular/angular-dragdrop.min.js"></script>
       <script src="/static/js/md-data-table.min.js"></script>

       <!-- Base App Dependencies -->

       <!-- Module declaration needs to come first -->
       <script src="/static/js/app/baseApp.js"></script>

       <!-- Other dependencies -->
       <script src="/static/js/moment.min.js"></script>

       <script>
       (function() {

       var app = angular.module('apiApp', [
         'md.data.table',
         'vAccordion',
         'ngMaterial']);

       //material colour scheme
       app.config(function($mdThemingProvider, $mdIconProvider){
         $mdThemingProvider.theme('default')
                             .primaryPalette('{{.PrimaryColour}}')
                             .accentPalette('{{.AccentColour}}');
       });

       })();
       </script>

  </head>

  <body layout="column"  ng-app="apiApp" ng-cloak>
      <md-toolbar layout="row" flex="10">
        <h1 flex><md-icon md-font-library="material-icons" style="font-size: 250%;">{{.Icon}}</md-icon> {{.Title}}</h1>
      </md-toolbar>

    <div layout="row" flex>
      <div flex layout="column" tabIndex="-1" role="main" class="md-whiteframe-z2">

        <md-content md-padding md-margin flex layout-fill style="padding: 12px; ">
