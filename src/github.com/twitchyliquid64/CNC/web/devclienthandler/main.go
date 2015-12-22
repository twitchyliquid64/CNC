package devclienthandler

import (
  "github.com/twitchyliquid64/CNC/data/session"
  pluginData "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/data"
  "golang.org/x/net/websocket"
  "time"
)


func Main_ws(ws *websocket.Conn){
  usrname := ws.Request().URL.Query().Get("user")
  passwd := ws.Request().URL.Query().Get("pass")
  pName := ws.Request().URL.Query().Get("plugin")
  var p pluginData.Plugin

  isValidLogin, usr := user.CheckAuthByPassword(usrname, passwd, data.DB)

  if isValidLogin && usr.IsAdmin() {
    skey := session.CreateSession(int(usr.ID), "devclient", data.DB)
    logging.Info("ws-devclient", "User '", usrname, "' has authenticated. Session: " + skey)
  } else {
    ws.Write(newPacket(&FatalError{Error: "Username / password incorrect or not an admin."}).Serialize())
    return
  }

  ws.Write(newPacket(&Status{Status: STATUS_AUTHENTICATED}).Serialize())

  if pName == ""{//just want a list of plugins
    ws.Write(newPacket(&Status{Status: STATUS_READY}).Serialize())
    pList := pluginData.GetAllEnabled(data.DB)
    ws.Write(newPacket(&PluginList{Plugins: pList}).Serialize())
    return
  } else {//fetch the plugin details
    p = pluginData.GetByName(data.DB, pName)
    if p.ID == 0{//plugin was not found
      ws.Write(newPacket(&FatalError{Error: "Could not find plugin '" + pName + "'"}).Serialize())
      return
    }
    ws.Write(newPacket(&Status{Status: STATUS_READY}).Serialize())
  }


  time.Sleep(20 * time.Second)
}
