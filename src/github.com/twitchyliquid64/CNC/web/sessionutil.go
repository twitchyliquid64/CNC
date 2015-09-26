package web

import (
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/data"
  "github.com/hoisie/web"
)


func getSessionByCookie(ctx *web.Context)(loggedIn bool, u *user.User, s *session.Session) {
  key := getSessionKey(ctx)
  if key == ""{
    return false, &user.User{}, &session.Session{}
  }

  s = session.GetSession(key, data.DB)
  if s != nil {
    u = user.GetUser(s.UserID, data.DB)
  }
  return s.IsValid(), u, s
}


func getSessionKey(ctx *web.Context)string{
	for _, cookie := range ctx.Request.Cookies(){
		if cookie.Name == "sid"{
			return cookie.Value
		}
	}
	return ""
}


func deleteSessionKey(ctx *web.Context){
	for _, cookie := range ctx.Request.Cookies(){
		if cookie.Name == "sid"{
			cookie.MaxAge = -1
		}
	}
}
