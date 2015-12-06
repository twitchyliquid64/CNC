package pluginhandler


import (
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/data"
  "net/http"
)

const COOKIE_KEY_NAME = "sid"

func getSessionByCookie(req *http.Request)(loggedIn bool, u *user.User, s *session.Session) {
  key := getSessionKey(req)
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


func getSessionKey(req *http.Request)string{
	for _, cookie := range req.Cookies(){
		if cookie.Name == COOKIE_KEY_NAME{
			return cookie.Value
		}
	}
	return ""
}
