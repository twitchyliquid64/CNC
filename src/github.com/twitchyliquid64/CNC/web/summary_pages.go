package web

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/hoisie/web"
)

func dashboardMainPage(ctx *web.Context) {
  t := templates.Lookup("dashboardindex")
	if t == nil {
		logging.Error("web", "No template found.")
	}
	t.Execute(ctx.ResponseWriter, nil)
}
