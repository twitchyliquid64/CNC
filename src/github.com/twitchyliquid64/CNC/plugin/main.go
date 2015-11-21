package plugin

import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/plugin/builtin"
  "github.com/twitchyliquid64/CNC/logging"
  "errors"
  "sync"
)

var pluginByName map[string]*exec.Plugin
var hooksByType map[string]map[string]exec.Hook //maps hooks[name] to plugins by name to hooks
var structureLock sync.Mutex

func Initialise(){
  logging.Info("plugin", "Initialise()")
  structureLock.Lock()
  defer structureLock.Unlock()

  pluginByName = map[string]*exec.Plugin{}
  hooksByType = map[string]map[string]exec.Hook{}

  //dependency injection
  exec.LoadBuiltinFunction = builtin.LoadBuiltinsToVM
  exec.RegisterHookFunction = RegisterHook
}

func RegisterPlugin(plugin *exec.Plugin){
  logging.Info("plugin", "RegisterPlugin() ", plugin.Name)
  structureLock.Lock()
  defer structureLock.Unlock()

  pluginByName[plugin.Name] = plugin
}

func DeregisterPlugin(plugin *exec.Plugin){
  logging.Info("plugin", "DeregisterPlugin() ", plugin.Name)
  structureLock.Lock()
  defer structureLock.Unlock()

  delete(pluginByName, plugin.Name)
  removeAllHooksOfPlugin(plugin)
}

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

//populates hooksByType with a hook for that specific plugin
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


//called to trigger all hooks with that name.
func Dispatch(hookName string, data interface{})bool{
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
