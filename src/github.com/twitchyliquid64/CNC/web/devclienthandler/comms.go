package devclienthandler

import (
  pluginData "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/logging"
  "encoding/json"
  "time"
)

type Packet struct {
  Type string
  Subdata []byte
}

func (p *Packet)Serialize()[]byte{
  d, err := json.Marshal(p)
  if err != nil{
    logging.Error("ws-devclient", err.Error())
    return []byte("")
  }
  return d
}

type Subpacket interface{
  Typ()string
}

func newPacket(pkt Subpacket)*Packet{
  d, err := json.Marshal(pkt)
  if err != nil{
    logging.Error("ws-devclient", err.Error())
    return nil
  }

  return &Packet{
    Type: pkt.Typ(),
    Subdata: d,
  }
}


type TextMsg struct{
  Fatal bool
  Message string
}
func (m *TextMsg)Typ()string{
  return "txtmsg"
}


type FatalError struct{
  Error string
}
func (m *FatalError)Typ()string{
  return "ferror"
}
func decodeFatalError(data []byte)*FatalError{
  var t FatalError
  err := json.Unmarshal(data, &t)
  if err != nil{
    return nil
  }
  return &t
}


const (
  REQUEST_PLUGININFO string = "plugininfo"
)
type DataRequest struct{
  DataType string
  ID int
}
func (m *DataRequest)Typ()string{
  return "dataRequest"
}
func decodeDataRequest(data []byte)*DataRequest{
  var t DataRequest
  err := json.Unmarshal(data, &t)
  if err != nil{
    return nil
  }
  return &t
}




type PluginInfo struct{
  P pluginData.Plugin //only on the server, have type Plugin on client
}
func (m *PluginInfo)Typ()string{
  return "plugininfo"
}
func decodePluginInfo(data []byte)*PluginInfo{
  var t PluginInfo
  err := json.Unmarshal(data, &t)
  if err != nil{
    return nil
  }
  return &t
}







type Plugin struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time

    Name string `sql:"not null;unique;index"`
    Icon string
    Description string
    Enabled bool

    HasCrashed bool
    ErrorStr string

    Resources []Resource
}


type Resource struct {
  ID int      `gorm:"primary_key"`
  PluginID int `sql:"index"`
  Name string `sql:"index"`
  Data []byte
  IsExecutable bool
  IsTemplate bool
  JSONData string `sql:"-"` //only used for JSON deserialisation - not a DB field
}



type PluginList struct{
  Plugins []pluginData.Plugin
}
func (m *PluginList)Typ()string{
  return "plist"
}

const (
  STATUS_AUTHENTICATED string = "AUTH OK"
  STATUS_READY string = "READY"
)


type Status struct{
  Status string
}
func (m *Status)Typ()string{
  return "status"
}
