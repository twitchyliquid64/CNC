package builtin

import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/robertkrimen/otto"
)

//Called to load up the VM with pointers to API functions.
//Injected into plugin.exec to avoid circular dependency.
//
func LoadBuiltinsToVM(plugin *exec.Plugin)error{
  //logging
  plugin.VM.Set("onLogMessage", func(in otto.FunctionCall)otto.Value{return function_onLogMessage(plugin, in)})
  plugin.VM.Set("log", func(in otto.FunctionCall)otto.Value{return function_onLog(plugin, in)})

  //telegram
  tgram, _ := plugin.VM.Object(`telegram = {}`)
  tgram.Set("onChatJoined", func(in otto.FunctionCall)otto.Value{return function_telegram_onChatJoined(plugin, in)})
  tgram.Set("onChatLeft", func(in otto.FunctionCall)otto.Value{return function_telegram_onChatLeft(plugin, in)})
  tgram.Set("onChatMsg", func(in otto.FunctionCall)otto.Value{return function_telegram_onChatMsg(plugin, in)})
  tgram.Set("sendMsg", func(in otto.FunctionCall)otto.Value{return function_telegram_sendMsg(plugin, in)})
  plugin.VM.Set("telegram", tgram)


  //gmail
  gmail, _ := plugin.VM.Object(`gmail = {}`)
  gmail.Set("setup", func(in otto.FunctionCall)otto.Value{return function_gmail_setup(plugin, in)})
  gmail.Set("sendMessage", func(in otto.FunctionCall)otto.Value{return function_gmail_sendMessage(plugin, in)})
  plugin.VM.Set("gmail", gmail)

  //twilio (SMS)
  twilio, _ := plugin.VM.Object(`twilio = {}`)
  twilio.Set("sendSMS", func(in otto.FunctionCall)otto.Value{return function_twilio_sendSMS(plugin, in)})
  plugin.VM.Set("twilio", twilio)
  //aux
  plugin.VM.Set("testendpoint_good", func(in otto.FunctionCall)otto.Value{return function_onTestEndpointGood(plugin, in)})
  plugin.VM.Set("onTestDispatchTriggered", func(in otto.FunctionCall)otto.Value{return function_onTestDispatchTriggered(plugin, in)})

  return nil
}
