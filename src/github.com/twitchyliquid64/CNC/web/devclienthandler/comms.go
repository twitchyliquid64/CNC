package devclienthandler

import (
  "github.com/twitchyliquid64/CNC/logging"
  "encoding/json"
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
  STATUS_AUTHENTICATED string = "AUTH OK"
  STATUS_READY string = "READY"
)


type Status struct{
  Status string
}
func (m *Status)Typ()string{
  return "status"
}
