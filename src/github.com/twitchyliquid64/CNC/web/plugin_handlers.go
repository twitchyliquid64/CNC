package web

import (
  pluginController "github.com/twitchyliquid64/CNC/plugin"
  pluginExec "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/logging"
  pluginData "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/data"
  "github.com/hoisie/web"
  "encoding/json"
  "strconv"
)


// Passes back a JSON object representing that particular
// plugin.
//
func getPluginHandlerAPI(ctx *web.Context){
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "getPluginHandlerAPI() called unauthorized, aborting")
    return
  }

  pluginID, _ := strconv.Atoi(ctx.Params["pluginid"])
  databaseObj := pluginData.Get(data.DB, pluginID)

  d, err := json.Marshal(databaseObj)
  if err != nil {
    logging.Error("web-plugin", err)
  }
  ctx.ResponseWriter.Write(d)
}



// Passes back a JSON array of all plugins
//
//
func getAllPluginsHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "getAllPlugins() called unauthorized, aborting")
    return
  }

  //get plugindata.Plugin's out of the currently running plugins
  var plugins []pluginData.Plugin
  temp := pluginController.GetAll()
  if temp == nil {
    temp = []*pluginExec.Plugin{}
  }
  for _, p := range temp {
    var tempPlugin pluginData.Plugin
    tempPlugin = p.Model
    tempPlugin.Resources = nil
    plugins = append(plugins, tempPlugin)
  }

  //turn them into JSON, combining the list from the running plugins,
  //and the list of disabled plugins.
  d, err := json.Marshal(
    struct{
      Running []pluginData.Plugin
      Disabled []pluginData.Plugin
    }{
    Running: plugins,
    Disabled: pluginData.GetAllDisabledNoResources(data.DB),
  })
  if err != nil {
    logging.Error("web-plugin", err)
  }
  ctx.ResponseWriter.Write(d)
}

func changePluginStateAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "newPlugin() called unauthorized, aborting")
    return
  }

  pluginID, _ := strconv.Atoi(ctx.Params["pluginid"])
  startPlugin := ctx.Params["state"] == "true"
  databaseObj := pluginData.Get(data.DB, pluginID)

  if startPlugin {
    databaseObj.Enabled = true
    data.DB.Save(&databaseObj)
    pluginController.StartPluginBasedFromDB(pluginData.Get(data.DB, pluginID))
  } else { //stop plugin
    plugin := pluginController.FindByName(databaseObj.Name)
    pluginController.DeregisterPlugin(plugin)
    plugin.Stop()
    databaseObj.Enabled = false
    data.DB.Save(&databaseObj)
  }
}

// API endpoint called to create a new plugin.
// Checks if the session's user is an admin.
//
func newPluginHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "newPlugin() called unauthorized, aborting")
    return
  }

  decoder := json.NewDecoder(ctx.Request.Body)
  var pl pluginData.Plugin
  err := decoder.Decode(&pl)
  if err != nil {
      logging.Error("web-plugin", "newPluginHandlerAPI() failed to decode JSON:", err)
      ctx.Abort(500, "JSON error")
      return
  }

  err = pluginData.Create(pl, data.DB)
  if err == nil {
      ctx.ResponseWriter.Write([]byte("GOOD"))
  } else {
      ctx.ResponseWriter.Write([]byte(err.Error()))
      logging.Error("web-plugin", err)
  }
}




// API endpoint called to create a new resource.
// Checks if the session's user is an admin.
//
func newResourceHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "newResourceHandlerAPI() called unauthorized, aborting")
    return
  }

  decoder := json.NewDecoder(ctx.Request.Body)
  var res pluginData.Resource
  err := decoder.Decode(&res)
  if err != nil {
      logging.Error("web-plugin", "newResourceHandlerAPI() failed to decode JSON:", err)
      ctx.Abort(500, "JSON error")
      return
  }
  res.Data = []byte(res.JSONData) //hack so that we can pass the data in as a string on clientside.
  res.JSONData = ""

  err = data.DB.Create(&res).Error
  if err == nil {
      ctx.ResponseWriter.Write([]byte("GOOD"))
  } else {
      ctx.ResponseWriter.Write([]byte(err.Error()))
      logging.Error("web-plugin", err)
  }
}



// API endpoint called to edit the general properties of an existing plugin.
// Checks if the session's user is an admin.
//
func editPluginHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "editPlugin() called unauthorized, aborting")
    return
  }

  //decode JSON to pl
  decoder := json.NewDecoder(ctx.Request.Body)
  var pl pluginData.Plugin
  err := decoder.Decode(&pl)
  if err != nil {
      logging.Error("web-plugin", "editPluginHandlerAPI() failed to decode JSON:", err)
      ctx.Abort(500, "JSON error")
      return
  }

  //to preserve integrity, shut down the plugin if it is currently running.
  existingDatabaseObj := pluginData.Get(data.DB, int(pl.ID))
  if existingDatabaseObj.Enabled { //must of been running, shut it down.
    plugin := pluginController.FindByName(existingDatabaseObj.Name)
    pluginController.DeregisterPlugin(plugin)
    plugin.Stop()
  }

  //FINALLY, save the changes.
  err = data.DB.Save(&pl).Error
  if err == nil {
    if pl.Enabled { //was saved successfully, so we should start it again if pl.Enabled == true
      pluginController.StartPluginBasedFromDB(pluginData.Get(data.DB, int(pl.ID)))
    }
    ctx.ResponseWriter.Write([]byte("GOOD"))
  } else {
      ctx.ResponseWriter.Write([]byte(err.Error()))
      logging.Error("web-plugin", err)
  }
}


// Passes back a JSON object representing that particular
// resource.
//
func getResourceHandlerAPI(ctx *web.Context){
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "getResourceHandlerAPI() called unauthorized, aborting")
    return
  }

  resourceID, _ := strconv.Atoi(ctx.Params["resourceid"])
  databaseObj := pluginData.GetResource(data.DB, resourceID)

  d, err := json.Marshal(databaseObj)
  if err != nil {
    logging.Error("web-plugin", err)
  }
  ctx.ResponseWriter.Write(d)
}




// API endpoint called to edit a resource.
// Checks if the session's user is an admin.
//
func editResourceHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "editResource() called unauthorized, aborting")
    return
  }

  //decode JSON to pl
  decoder := json.NewDecoder(ctx.Request.Body)
  var res pluginData.Resource
  err := decoder.Decode(&res)
  if err != nil {
      logging.Error("web-plugin", "editResourceHandlerAPI() failed to decode JSON:", err)
      ctx.Abort(500, "JSON error")
      return
  }
  res.Data = []byte(res.JSONData) //hack so that we can pass the data in as a string on clientside.
  res.JSONData = ""

  err = data.DB.Save(&res).Error
  if err == nil {
    ctx.ResponseWriter.Write([]byte("GOOD"))
  } else {
      ctx.ResponseWriter.Write([]byte(err.Error()))
      logging.Error("web-plugin", err)
  }
}


// API endpoint called to delete a resource using a resourceID.
// Checks if the session's user is an admin.
//
func deleteResourceHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "deleteResource() called unauthorized, aborting")
    return
  }
  err := data.DB.Where("id = ?", ctx.Params["resourceid"]).Delete(&pluginData.Resource{}).Error
  if err == nil {
    ctx.ResponseWriter.Write([]byte("GOOD"))
  } else {
      ctx.ResponseWriter.Write([]byte(err.Error()))
      logging.Error("web-plugin", err)
  }
}


// API endpoint called to delete a plugin using a pluginID.
// Checks if the session's user is an admin.
//
// TODO: Make method stop the plugin if it is running first.
func deletePluginHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "deletePlugin() called unauthorized, aborting")
    return
  }
  err := data.DB.Where("id = ?", ctx.Params["pluginid"]).Delete(&pluginData.Plugin{}).Error
  if err == nil {
    ctx.ResponseWriter.Write([]byte("GOOD"))
  } else {
      ctx.ResponseWriter.Write([]byte(err.Error()))
      logging.Error("web-plugin", err)
  }
}
