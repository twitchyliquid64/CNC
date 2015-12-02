package pluginhandler

import (
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


func newReqStruct(w http.ResponseWriter, req *http.Request)*Request {
  buf := new(bytes.Buffer)
  buf.ReadFrom(req.Body)

  return &Request{
    Body: buf.String(),
    IsFinished: false,
    W: w,
    Req: req,
    Finish: make(chan bool, 1),
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
