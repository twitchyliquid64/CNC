package web

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/config"
  "github.com/hoisie/web"
)

func registerUserHandlers() {
  logging.Info("web", "Registering page handlers")
  web.Get("/", loginHandler, config.All().Web.Domain)
}

func RegisterHandlers() {
  registerUserHandlers()
}
