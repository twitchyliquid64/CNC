package messenger

import (
  "github.com/twitchyliquid64/CNC/registry/syscomponents"
)

var trackerObj TelegramComponent

type TelegramComponent struct{
  Disabled bool
  err error
}

func (d *TelegramComponent)Name() string{
  return "Telegram"
}
func (d *TelegramComponent)IconStr() string{
  return "send"
}
func (d *TelegramComponent)IsNominal()bool{
  return d.err == nil
}
func (d *TelegramComponent)IsDisabled()bool{
  return d.Disabled
}
func (d *TelegramComponent)IsFault()bool{
  return d.err != nil
}
func (d *TelegramComponent)Error()string{
  if d.err == nil{
    return ""
  }
  return d.err.Error()
}
func (d *TelegramComponent)SetError(e error){
  d.err = e
}

func trackingSetup(disabled bool){
  trackerObj = TelegramComponent{Disabled: disabled}
  syscomponents.Register(&trackerObj)
}

func tracking_notifyFault(err error){
  syscomponents.SetError(trackerObj.Name(), err)
}
