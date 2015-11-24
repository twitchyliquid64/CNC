package plugin

import (

  pluginData "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/logging"
  "errors"
)



//Called with a fully populated plugin (no trimming of resources)
//To create a plugin and get it running.
//
func StartPluginBasedFromDB(plugin pluginData.Plugin){
  structureLock.Lock()
  defer structureLock.Unlock()
  startPluginBasedFromDB(plugin)
}


//Called to register an already-existing plugin with the system. This method
//should probably not be used - use startPluginBasedFromDB() instead.
//
func RegisterPlugin(plugin *exec.Plugin)error{
  logging.Info("plugin", "RegisterPlugin() ", plugin.Name)
  structureLock.Lock()
  defer structureLock.Unlock()

  _, ok := pluginByName[plugin.Name]
  if ok{
    logging.Error("plugin", "Attempted to add plugin '", plugin.Name, "' which already exists!")
    return errors.New("Plugin already exists")
  }

  pluginByName[plugin.Name] = plugin
  return nil
}



// Do this prior to stopping a plugin - remove it from tracking so it is not visible
// to incoming dispatches/hooks.
//
func DeregisterPlugin(plugin *exec.Plugin){
  logging.Info("plugin", "DeregisterPlugin() ", plugin.Name)
  structureLock.Lock()
  defer structureLock.Unlock()

  delete(pluginByName, plugin.Name)
  removeAllHooksOfPlugin(plugin)
}


// If you know a plugin is running and is named something, you can get the
// plugin object. Typically used to .Stop() it.
//
func FindByName(name string)*exec.Plugin{
  structureLock.Lock()
  defer structureLock.Unlock()
  return pluginByName[name]
}


// Called internally during deregistration to remove hooks associated with
// the plugin.
//
func removeAllHooksOfPlugin(plugin *exec.Plugin){//assumes structureLock is held
  for hookType, _ := range hooksByType {
    _, ok := hooksByType[hookType][plugin.Name]
    if ok {
      hooksByType[hookType][plugin.Name].Destroy()
      delete(hooksByType[hookType], plugin.Name)
      logging.Info("plugin", "Found hook ", hookType, " for plugin ", plugin.Name, ", deleting")
    }
  }
}

// Populates hooksByType with a hook for that specific plugin
// Typically called internally.
//
func RegisterHook(plugin *exec.Plugin, hook exec.Hook)error {
  logging.Info("plugin", "RegisterHook() ", hook.Name())
  structureLock.Lock()
  defer structureLock.Unlock()

  _, ok := pluginByName[plugin.Name]
  if !ok{
    return errors.New("Plugin is not registered")
  }

  _, ok = hooksByType[hook.Name()]
  if !ok{
    hooksByType[hook.Name()] = map[string]exec.Hook{plugin.Name: hook}
  }else{
    hooksByType[hook.Name()][plugin.Name] = hook
  }
  return nil
}


// Called to trigger all hooks with that name, passing data
// to the hooks' handler method.
//
func Dispatch(hookName string, data interface{})bool{
  if !isInitialised{return false}

  var foundSome bool = false
  logging.Info("plugin", "Dispatch() called")
  structureLock.Lock()
  defer structureLock.Unlock()

  hookSet, ok := hooksByType[hookName]
  if !ok{
    return false
  }

  for _, hook := range hookSet{
    foundSome = true
    hook.Dispatch(data)
  }
  return foundSome
}


// Called to return a structure of all the plugins
//
//
func GetAll()[]*exec.Plugin {
  structureLock.Lock()
  defer structureLock.Unlock()

  var ret []*exec.Plugin

  for _, plugin := range pluginByName {
    ret = append(ret, plugin)
  }
  return ret
}
