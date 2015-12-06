package pluginhandler

import (
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/registry"
	"net/http"
  "bytes"
  "time"
)


const PLUGIN_TIMEOUT = 5 * time.Second

type Request struct {
  Body string
  IsFinished bool
  W http.ResponseWriter
  Req *http.Request
  Finish chan bool

  hasCheckedSession bool //set if Session, isLoggedIn, and User are populated
  isLoggedIn bool
  user *user.User
  session *session.Session
}

func (r *Request)Write(d string){
  r.W.Write([]byte(d))
}
func (r *Request)Done(){
  r.IsFinished = true
  r.Finish <- true
}
func (r *Request)URL()string{
  return r.Req.URL.String()
}
func (r *Request)Parameter(key string)string{
  return r.Req.URL.Query().Get(key)
}
func (r *Request)PostBody()string{
  return r.Body
}

func (r *Request)LoggedIn()bool{
  if !r.hasCheckedSession{
    r.loadSessionData()
  }
  return r.isLoggedIn
}
func (r *Request)User()*user.User{
  if !r.hasCheckedSession{
    r.loadSessionData()
  }
  return r.user
}
func (r *Request)Session()*session.Session{
  if !r.hasCheckedSession{
    r.loadSessionData()
  }
  return r.session
}

func (r *Request)loadSessionData(){
  r.hasCheckedSession = true
  r.isLoggedIn, r.user, r.session = getSessionByCookie(r.Req)
}



func newReqStruct(w http.ResponseWriter, req *http.Request)*Request {
  buf := new(bytes.Buffer)
  buf.ReadFrom(req.Body)

  return &Request{
    Body: buf.String(),
    IsFinished: false,
    W: w,
    Req: req,
    Finish: make(chan bool, 1),
    hasCheckedSession: false,
    isLoggedIn: false,
  }
}

func HandleHTTP(w http.ResponseWriter, req *http.Request) {
  logging.Info("web-plugin", "Req from: ", req.URL)

  hook := findMatch(req.URL.Path)

  if hook == "" {
    http.Error(w, "No handler for the specified path", 500)
  } else {
    r := newReqStruct(w, req)
    registry.DispatchEvent(hook, r)

    select {
      case <- r.Finish:
      case <-time.After(PLUGIN_TIMEOUT):
    }

    if !r.IsFinished {
      http.Error(w, "Plugin Timeout", 500)
    }
  }
}
