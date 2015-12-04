package builtin

import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/data/stmdata"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/data"
  "github.com/robertkrimen/otto"
)





// Called when JS code executes data.get(<key>)
//
//
func function_data_get(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  key := call.Argument(0).String()

  obj, exists := stmdata.Get(int(plugin.Model.ID), key, data.DB)
  if !exists {
    return otto.Value{} //undefined
  }

  obj.Content = string(obj.Data)
  obj.Data = []byte("")

  val, err := plugin.VM.ToValue(obj)
  if err != nil {
    logging.Error("builtin-data", err.Error())
    return otto.Value{}
  }else {
    return val
  }
}





// Called when JS code executes data.set(<key>, <data>)
//
//
func function_data_set(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  key := call.Argument(0).String()
  dat := call.Argument(1).String()

  err := stmdata.Set(int(plugin.Model.ID), key, dat, data.DB)
  if err != nil {
    logging.Error("builtin-data", err.Error())
    val, _ := plugin.VM.ToValue(map[string]interface{}{"error": err.Error()})
    return val
  } else {
    return otto.TrueValue()
  }
}
