package devclienthandler

import (
  "github.com/twitchyliquid64/CNC/data/session"
  pluginData "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/data"
  "golang.org/x/net/websocket"
  "encoding/json"
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
    pList := pluginData.GetAllEnabled_NoResources(data.DB)
    ws.Write(newPacket(&PluginList{Plugins: pList}).Serialize())
    return
  } else {//fetch the plugin details
    p = pluginData.GetByName(data.DB, pName)
    if p.ID == 0{//plugin was not found
      ws.Write(newPacket(&FatalError{Error: "Could not find plugin '" + pName + "'"}).Serialize())
      return
    }
    pluginData.LoadResources(&p, data.DB)
    ws.Write(newPacket(&Status{Status: STATUS_READY}).Serialize())
  }

  //at this stage: we are authenticated, and an existing plugin is selected. Lets loop and execute commands.
  for {
    var data []byte
    err := websocket.Message.Receive(ws, &data)
    if err != nil{
      return
    }
    processMessage(ws, data, p)
  }
}




func processRequest(ws *websocket.Conn, data []byte, p pluginData.Plugin){
  msg := decodeDataRequest(data)
  switch msg.DataType {
  case REQUEST_PLUGININFO:
    ws.Write(newPacket(&PluginInfo{P: p}).Serialize())
  }
}

func processMessage(ws *websocket.Conn, data []byte, p pluginData.Plugin){
  var pkt Packet
  err := json.Unmarshal(data, &pkt)
  if err != nil{
    logging.Info("ws-devclient", "JSON Error: ", err.Error())
    return
  }

  switch pkt.Type {
  case "dataRequest":
    processRequest(ws, pkt.Subdata, p)

  default:
    logging.Info("ws-devclient", "Unknown type: ", pkt.Type, " ---- ", string(pkt.Subdata))
  }
}
