package builtin

import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/robertkrimen/otto"
)



// Called when JS code executes onLogMessage()
//
//
func function_onLogMessage(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  return otto.Value{}
}


// Called when JS code executes log()
//
//
func function_onLog(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  logging.Info("plugin-" + plugin.Name, call.Argument(0).String())
  return otto.Value{}
}
