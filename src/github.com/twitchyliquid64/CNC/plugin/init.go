package plugin

import (

  pluginData "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/plugin/builtin"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/data"
  "sync"
)

var isInitialised bool = false
var pluginByName map[string]*exec.Plugin
var hooksByType map[string]map[string]exec.Hook //maps hooks[name] to plugins by name to hooks
var structureLock sync.Mutex

func Initialise(loadFromDatabase bool){
  logging.Info("plugin", "Initialise()")
  structureLock.Lock()
  defer structureLock.Unlock()

  pluginByName = map[string]*exec.Plugin{}
  hooksByType = map[string]map[string]exec.Hook{}
  isInitialised = true

  //dependency injection
  exec.LoadBuiltinFunction = builtin.LoadBuiltinsToVM
  exec.RegisterHookFunction = RegisterHook
  if loadFromDatabase{
    startEnabledPluginsFromDatabase()
  }
}



func startEnabledPluginsFromDatabase(){ //assumes lock is held
  plugins := pluginData.GetAllEnabledNoTrim(data.DB)
  for _, plugin := range plugins {
    startPluginBasedFromDB(plugin)
  }
}

func startPluginBasedFromDB(plugin pluginData.Plugin){//assumes lock is held
  logging.Info("plugin", "Starting plugin ", plugin.Name)
  createdPluginObj := exec.BuildPluginFromDatabase(plugin.Name, plugin, plugin.Resources)
  structureLock.Unlock()
  err := RegisterPlugin(createdPluginObj) //if there was an error is was already posted to logging
  structureLock.Lock()
  if err != nil{
    createdPluginObj.Stop()
  }
}
