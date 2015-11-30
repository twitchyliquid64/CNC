package builtin

import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/robertkrimen/otto"
  "github.com/twitchyliquid64/CNC/config"
  "github.com/twitchyliquid64/twiliocore"
  "errors"
)


func getTwilioObj()*twiliocore.TwilioClient {
  if config.All().Twilio.Enable {
    return twiliocore.NewClient(config.All().Twilio.AccountSID, config.All().Twilio.AuthToken)
  }
  return nil
}


func sendMsg(fromNum, toNum, msg string)(*twiliocore.Message,error){
  obj := getTwilioObj()
  if obj == nil{
    return nil, errors.New("Configuration error")
  }
  return obj.NewMessage(fromNum, toNum, msg, "")
}



func function_twilio_sendSMS(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  from := call.Argument(0).String()
  to   := call.Argument(1).String()
  msg  := call.Argument(2).String()

  resp, err := sendMsg(from, to, msg)
  if err != nil{
    logging.Error("builtin-twilio", err.Error())
    out, err := plugin.VM.ToValue(resp)
    if err != nil {
      logging.Error("builtin-twilio", err.Error())
      out, _ = plugin.VM.ToValue(err.Error())
    }
    return out
  } else {
    out, _ := plugin.VM.ToValue(err.Error())
    return out
  }
}
