package httpgateway

import (
  "github.com/twitchyliquid64/CNC/registry/syscomponents"
)

var trackerObj GatewayComponent

type GatewayComponent struct{
  Disabled bool
  err error
}

func (d *GatewayComponent)Name() string{
  return "HTTP Gateway"
}
func (d *GatewayComponent)IconStr() string{
  return "http"
}
func (d *GatewayComponent)IsNominal()bool{
  return d.err == nil
}
func (d *GatewayComponent)IsDisabled()bool{
  return d.Disabled
}
func (d *GatewayComponent)IsFault()bool{
  return d.err != nil
}
func (d *GatewayComponent)Error()string{
  if d.err == nil{
    return ""
  }
  return d.err.Error()
}
func (d *GatewayComponent)SetError(e error){
  d.err = e
}

func trackingSetup(disabled bool){
  trackerObj = GatewayComponent{Disabled: disabled}
  syscomponents.Register(&trackerObj)
}

func tracking_notifyFault(err error){
  syscomponents.SetError(trackerObj.Name(), err)
}
