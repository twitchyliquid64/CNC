package data

import (
  "github.com/twitchyliquid64/CNC/logging"
  "golang.org/x/net/websocket"
)

func SqlQueryServer(ws *websocket.Conn){
  for {
    var data []byte
    err := websocket.Message.Receive(ws, &data)
    if err != nil{
      logging.Warning("data-websock", "Recieve error: ", err.Error())
      return
    }

    //process message
  }
}
