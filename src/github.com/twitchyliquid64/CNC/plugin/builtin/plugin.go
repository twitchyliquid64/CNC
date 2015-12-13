package builtin

import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  pluginModel "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/registry"
  "github.com/twitchyliquid64/CNC/data"
  "github.com/twitchyliquid64/CNC/util"
  "github.com/robertkrimen/otto"
)



// Called when JS code executes plugin.ready(<methodname>)
//
//
func function_plugin_ready(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  callback := util.GetFunc(call.Argument(0), plugin.VM)

  go func(){
    plugin.PendingInvocations <- &exec.JSInvocation{Callback: &callback}
  }()

  return otto.Value{}
}



// Called when JS code executes plugin.disable()
//
//
func function_plugin_disable(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  registry.DeregisterPluginMethod(plugin)
  go func(){
    plugin.Stop()
  }()

  panic("plugin.exit()")
  return otto.Value{}
}

// Called when JS code executes plugin.getIcon()
//
//
func function_plugin_geticon(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  res, _ := plugin.VM.ToValue(plugin.Model.Icon)
  return res
}

// Called when JS code executes plugin.setIcon()
//
//
func function_plugin_seticon(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  i := call.Argument(0).String()
  data.DB.Save(&pluginModel.Plugin{ID: plugin.Model.ID, Icon: i})
  plugin.Model.Icon = i
  return otto.Value{}
}



// Called when JS code executes plugin.getResources()
//
//
func function_plugin_getResources(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{

  var output []map[string]interface{}
  for _, resource := range plugin.Resources {
    temp := map[string]interface{}{}
    temp["name"] = resource.Name
    temp["isJs"] = resource.IsExecutable
    temp["isTemplate"] = resource.IsTemplate
    temp["data"] = string(resource.Data)
    output = append(output, temp)
  }

  val, err := plugin.VM.ToValue(output)
  if err != nil {
    logging.Error("builtin-plugin", err.Error())
    return otto.Value{}
  }else {
    return val
  }
}




// Called when JS code executes plugin.getResource(<resource name>)
//
//
func function_plugin_getResource(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  resourceName := call.Argument(0).String()

  var resource *pluginModel.Resource = nil
  for _, rCandidate := range plugin.Resources {
    if rCandidate.Name == resourceName {
      resource = &rCandidate
      break
    }
  }

  //if it was not found
  if resource == nil {
    return otto.Value{}
  }

  temp := map[string]interface{}{}
  temp["name"] = resource.Name
  temp["isJs"] = resource.IsExecutable
  temp["isTemplate"] = resource.IsTemplate
  temp["data"] = string(resource.Data)

  val, err := plugin.VM.ToValue(temp)
  if err != nil {
    logging.Error("builtin-plugin", err.Error())
    return otto.Value{}
  }else {
    return val
  }
}
