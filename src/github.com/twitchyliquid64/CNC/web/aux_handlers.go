package web

import (
  "github.com/twitchyliquid64/CNC/registry/syscomponents"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/hoisie/web"
)


func getSysComponentsStatusAPIHandler(ctx *web.Context) {
  isLoggedIn, _, _ := getSessionByCookie(ctx)

  if !isLoggedIn{
    logging.Warning("web-entity", "getSysComponents() called unauthorized, aborting")
    return
  }

  ctx.ResponseWriter.Write([]byte(syscomponents.GetJSON()))
}
