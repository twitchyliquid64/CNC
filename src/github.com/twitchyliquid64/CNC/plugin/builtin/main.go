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

  //aux
  plugin.VM.Set("testendpoint_good", func(in otto.FunctionCall)otto.Value{return function_onTestEndpointGood(plugin, in)})
  plugin.VM.Set("onTestDispatchTriggered", func(in otto.FunctionCall)otto.Value{return function_onTestDispatchTriggered(plugin, in)})
  return nil
}
