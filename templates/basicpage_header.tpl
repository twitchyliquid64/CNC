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

  </head>

  <body layout="column"  ng-app="apiApp" ng-cloak>
      <md-toolbar layout="row" flex="10">
        <h1 flex><md-icon md-font-library="material-icons" style="font-size: 250%;">{{.Icon}}</md-icon> {{.Title}}</h1>
      </md-toolbar>

    <div layout="row" flex>
      <div flex layout="column" tabIndex="-1" role="main" class="md-whiteframe-z2">

        <md-content md-padding md-margin flex layout-fill style="padding: 12px; ">
