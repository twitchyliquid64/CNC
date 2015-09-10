package web

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/config"
  "github.com/hoisie/web"
)

func registerUserHandlers() {
  web.Get("/", loginHandler, config.All().Web.Domain)
}

func Initialise() {
  logging.Info("web", "Registering page handlers")
  registerUserHandlers()

  logging.Info("web", "Registering templates")
  registerUserTemplates()
}

func registerUserTemplates(){
  logError(registerTemplate("test.tpl", "test"), "Template load error: ")
}

func logError(e error, prefix string){
  if e != nil{
    logging.Error("web", prefix, e.Error())
  }
}
