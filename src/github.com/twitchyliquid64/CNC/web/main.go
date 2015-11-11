package web

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/config"
  "github.com/hoisie/web"
)

func Run() {
  logging.Info("web", "Initialising server")
  web.RunTLS(config.All().Web.Listener, config.TLS())
  //web.Run(config.All().Web.Listener)
}
