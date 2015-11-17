package data

import (
  "github.com/twitchyliquid64/CNC/registry/syscomponents"
)

var trackerObj DatabaseComponent

type DatabaseComponent struct{
  err error
}

func (d *DatabaseComponent)Name() string{
  return "Database"
}
func (d *DatabaseComponent)IconStr() string{
  return "list"
}
func (d *DatabaseComponent)IsNominal()bool{
  return d.err == nil
}
func (d *DatabaseComponent)IsDisabled()bool{
  return false
}
func (d *DatabaseComponent)IsFault()bool{
  return d.err != nil
}
func (d *DatabaseComponent)Error()string{
  if d.err == nil{
    return ""
  }
  return d.err.Error()
}
func (d *DatabaseComponent)SetError(e error){
  d.err = e
}

func trackingSetup(){
  trackerObj = DatabaseComponent{}
  syscomponents.Register(&trackerObj)
}

func tracking_notifyFault(err error){
  syscomponents.SetError(trackerObj.Name(), err)
}
