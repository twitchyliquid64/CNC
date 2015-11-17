package web

import (
  "github.com/twitchyliquid64/CNC/registry/syscomponents"
)

var trackerObj WebComponent

type WebComponent struct{
  Disabled bool
  err error
}

func (d *WebComponent)Name() string{
  return "Web Application"
}
func (d *WebComponent)IconStr() string{
  return "language"
}
func (d *WebComponent)IsNominal()bool{
  return d.err == nil
}
func (d *WebComponent)IsDisabled()bool{
  return d.Disabled
}
func (d *WebComponent)IsFault()bool{
  return d.err != nil
}
func (d *WebComponent)Error()string{
  if d.err == nil{
    return ""
  }
  return d.err.Error()
}
func (d *WebComponent)SetError(e error){
  d.err = e
}

func trackingSetup(disabled bool){
  trackerObj = WebComponent{Disabled: disabled}
  syscomponents.Register(&trackerObj)
}

func tracking_notifyFault(err error){
  syscomponents.SetError(trackerObj.Name(), err)
}
