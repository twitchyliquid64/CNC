package web

import (
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/hoisie/web"
)

type mainPageData struct {
  User *user.User
  Session *session.Session
}

func (inst mainPageData)IsAdmin()bool{
  for _, perm := range inst.User.Permissions {
    if perm.Name == user.PERM_ADMIN {
      return true
    }
  }
  return false
}

func dashboardMainPage(ctx *web.Context) {
  isLoggedIn, user, session := getSessionByCookie(ctx)

  if !isLoggedIn {
    ctx.Redirect(302, "/login")
    return
  }

  //all code from this point on can assume that
  //user and session are both populated

  t := templates.Lookup("dashboardindex")
	if t == nil {
		logging.Error("web", "No template found.")
	}
	t.Execute(ctx.ResponseWriter, mainPageData{User: user, Session: session})
}
