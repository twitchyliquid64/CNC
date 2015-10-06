package web

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/config"
  "github.com/hoisie/web"
)



func Initialise() {
  logging.Info("web", "Registering page handlers")
  registerCoreHandlers()
  registerUserHandlers()
  registerSummaryHandlers()

  logging.Info("web", "Registering templates")
  registerCoreTemplates()
  registerUserTemplates()
  registerSummaryTemplates()
}

func registerCoreHandlers() {
  web.Get("/login", loginMainPage, config.All().Web.Domain)
  web.Get("/dev/reload", templateReloadHandler, config.All().Web.Domain)
}

func registerUserHandlers() {
  web.Post("/login", loginHandler, config.All().Web.Domain)
  web.Get("/users", getUsersHandlerAPI, config.All().Web.Domain)
  web.Get("/user", getUserHandlerAPI, config.All().Web.Domain)
  web.Get("/user/delete", deleteUserHandlerAPI, config.All().Web.Domain)
  web.Get("/logout", logoutHandler, config.All().Web.Domain)
  web.Post("/users/new", newUserHandlerAPI, config.All().Web.Domain)
  web.Post("/users/edit", updateUserHandlerAPI, config.All().Web.Domain)
  web.Get("/user/permission/add", addPermissionUserHandlerAPI, config.All().Web.Domain)
  web.Get("/user/permission/delete", deletePermissionUserHandlerAPI, config.All().Web.Domain)
}

func registerSummaryHandlers(){
  web.Get("/", dashboardMainPage, config.All().Web.Domain)
}

func registerCoreTemplates(){
  logError(registerTemplate("bannertop.tpl", "bannertop"), "Template load error: ")
  logError(registerTemplate("headcontent.tpl", "headcontent"), "Template load error: ")
  logError(registerTemplate("tailcontent.tpl", "tailcontent"), "Template load error: ")
}

func registerUserTemplates(){
  logError(registerTemplate("test.tpl", "test"), "Template load error: ")
  logError(registerTemplate("login.tpl", "login"), "Template load error: ")
  logError(registerTemplate("userpage.tpl", "userpage"), "Template load error: ")
  logError(registerTemplate("usercreateeditpage.tpl", "usercreateeditpage"), "Template load error: ")
  logError(registerTemplate("userpermissions.tpl", "userpermissions"), "Template load error: ")
}

func registerSummaryTemplates(){
  logError(registerTemplate("dashboardindex.tpl", "dashboardindex"), "Template load error: ")
}


func logError(e error, prefix string){
  if e != nil{
    logging.Error("web", prefix, e.Error())
  }
}
