package web

import (
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/data"
  "github.com/hoisie/web"
)

const COOKIE_KEY_NAME = "sid"

func getSessionByCookie(ctx *web.Context)(loggedIn bool, u *user.User, s *session.Session) {
  key := getSessionKey(ctx)
  if key == ""{
    return false, &user.User{}, &session.Session{}
  }

  s = session.GetSession(key, data.DB)
  if s != nil {
    u = user.GetUser(s.UserID, data.DB)
    return s.IsValid(), u, s
  }else{
    return false, &user.User{}, &session.Session{}
  }
}


func getSessionKey(ctx *web.Context)string{
	for _, cookie := range ctx.Request.Cookies(){
		if cookie.Name == COOKIE_KEY_NAME{
			return cookie.Value
		}
	}
	return ""
}


func deleteSessionKey(ctx *web.Context){
	for _, cookie := range ctx.Request.Cookies(){
		if cookie.Name == COOKIE_KEY_NAME{
			cookie.MaxAge = -1
		}
	}
}
