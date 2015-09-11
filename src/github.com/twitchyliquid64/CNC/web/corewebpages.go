package web

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/hoisie/web"
)

func loginMainPage(ctx *web.Context) {
  t := templates.Lookup("login")
	if t == nil {
		logging.Error("web", "No template found.")
	}
	t.Execute(ctx.ResponseWriter, nil)
}
