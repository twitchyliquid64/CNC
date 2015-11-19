package web

import (
  "github.com/twitchyliquid64/CNC/logging"
  "golang.org/x/net/websocket"
  "github.com/hoisie/web"
  "io"
)

// Echo the data received on the WebSocket.
func ws_EchoServer(ws *websocket.Conn) {
    io.Copy(ws, ws)
}

func ws_LogServer(ws *websocket.Conn){
  isLoggedIn, u, _ := getSessionByCookie(&web.Context{Request: ws.Request()})

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-sockets", "logServer() called unauthorized, aborting")
    return
  }

  //transmit the log messages stored in history
  for _, msg := range logging.GetBacklog() {
    err := websocket.JSON.Send(ws, msg)
    if err != nil{
      return
    }
  }

  logMessages := make(chan logging.LogMessage, 10)

  logging.Subscribe(logMessages)
  defer logging.Unsubscribe(logMessages)

  for msg := range logMessages {
    err := websocket.JSON.Send(ws, msg)
    if err != nil{
      return
    }
  }
}
