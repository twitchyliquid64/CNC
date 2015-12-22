package devclienthandler

import (
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/data"
  "golang.org/x/net/websocket"
  "time"
)


func Main_ws(ws *websocket.Conn){
  usrname := ws.Request().URL.Query().Get("user")
  passwd := ws.Request().URL.Query().Get("pass")

  isValidLogin, usr := user.CheckAuthByPassword(usrname, passwd, data.DB)

  if isValidLogin && usr.IsAdmin() {
    skey := session.CreateSession(int(usr.ID), "devclient", data.DB)
    logging.Info("ws-devclient", "User '", usrname, "' has authenticated. Session: " + skey)
  } else {
    ws.Write(newPacket(&FatalError{Error: "Username / password incorrect or not an admin."}).Serialize())
    return
  }

  ws.Write(newPacket(&Status{Status: STATUS_AUTHENTICATED}).Serialize())
  ws.Write(newPacket(&FatalError{Error: "Nothing more is implemented."}).Serialize())

  time.Sleep(20 * time.Second)
}
