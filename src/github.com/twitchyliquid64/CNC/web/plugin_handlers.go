package web

import (
  pluginController "github.com/twitchyliquid64/CNC/plugin"
  pluginExec "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/logging"
  //"github.com/twitchyliquid64/CNC/data"
  "github.com/hoisie/web"
  "encoding/json"
)

// Passes back a JSON array of all plugins
//
//
func getAllPluginsHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "getAllPlugins() called unauthorized, aborting")
    return
  }

  plugins := pluginController.GetAll()
  if plugins == nil {
    plugins = []*pluginExec.Plugin{}
  }

  d, err := json.Marshal(plugins)
  if err != nil {
    logging.Error("web-plugin", err)
  }
  ctx.ResponseWriter.Write(d)
}
