package web

import (
  "io"
   "golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func ws_EchoServer(ws *websocket.Conn) {
    io.Copy(ws, ws)
}
