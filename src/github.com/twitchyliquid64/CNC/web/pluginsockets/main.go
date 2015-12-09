package pluginsockets

import (
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/registry"
  "golang.org/x/net/websocket"
  "github.com/hoisie/web"
	"net/http"
  "time"
  "sync"
)


const PLUGIN_TIMEOUT = 8 * time.Second

type Socket struct {
  IsFinished bool
  ID uint64

  Req *http.Request
  Ws *websocket.Conn

  Finish chan bool

  hasCheckedSession bool //set if Session, isLoggedIn, and User are populated
  isLoggedIn bool
  user *user.User
  session *session.Session
}


var totalSockReqCount uint64 = 0
var totalLock sync.Mutex
func getId()uint64{
  totalLock.Lock()
  defer totalLock.Unlock()
  totalSockReqCount += 1
  return totalSockReqCount-1
}

func (r *Socket)Write(d string){
  if !r.IsFinished{
    websocket.Message.Send(r.Ws, d)
  }
}
func (r *Socket)Close(){
  if !r.IsFinished{
    r.Ws.Close()
    r.IsFinished = true
  }
}
func (r *Socket)Addr()string{
  return r.Ws.RemoteAddr().String()
}

func (r *Socket)GetID()uint64{
  return r.ID
}

func (r *Socket)URL()string{
  return r.Req.URL.String()
}
func (r *Socket)Parameter(key string)string{
  return r.Req.URL.Query().Get(key)
}




func (r *Socket)LoggedIn()bool{
  if !r.hasCheckedSession{
    r.loadSessionData()
  }
  return r.isLoggedIn
}
func (r *Socket)User()*user.User{
  if !r.hasCheckedSession{
    r.loadSessionData()
  }
  return r.user
}
func (r *Socket)Session()*session.Session{
  if !r.hasCheckedSession{
    r.loadSessionData()
  }
  return r.session
}

func (r *Socket)loadSessionData(){
  r.hasCheckedSession = true
  r.isLoggedIn, r.user, r.session = getSessionByCookie(r.Req)
}


type SocketEvent struct {
  EventType string
  Data string
  Socket *Socket
}
func (r SocketEvent)Sock()interface{}{
  return r.Socket
}
func (r SocketEvent)Event()string{
  return r.EventType
}
func (r SocketEvent)GetData()string{
  return r.Data
}



func newReqStruct(ws *websocket.Conn, req *http.Request)*Socket {

  return &Socket{
    IsFinished: false,
    Req: req,
    Ws: ws,
    Finish: make(chan bool, 1),
    hasCheckedSession: false,
    isLoggedIn: false,
    ID: getId(),
  }
}

func Handle(ctx *web.Context, dummy string){
  websocket.Handler(HandleWebsocket).ServeHTTP(ctx, ctx.Request)
}

func HandleWebsocket(ws *websocket.Conn) {
  req := ws.Request()
  logging.Info("web-pluginsocks", "Req from: ", req.URL)

  hookRec := findMatch(req.URL.Path)

  if hookRec.OnOpen == "" {
    ws.Write([]byte("No websocket handler for the specified request"))
  } else {
    r := newReqStruct(ws, req)
    registry.DispatchEvent(hookRec.OnOpen, SocketEvent{EventType: "OPEN", Socket: r})

    for !r.IsFinished {
      var reply string
      if err := websocket.Message.Receive(ws, &reply); err != nil {
          r.IsFinished = true
          registry.DispatchEvent(hookRec.OnClose, SocketEvent{EventType: "CLOSE", Socket: r})
      } else {
        registry.DispatchEvent(hookRec.OnMessage, SocketEvent{EventType: "MSG", Socket: r, Data: reply})
      }
    }
  }
}
