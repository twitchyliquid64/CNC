package exec

import (
  "github.com/twitchyliquid64/CNC/logging"
)

// Called by all factory methods to initialise
//
//
func initialise(plugin *Plugin) {
  logging.Info("plugin", "plugin.initialise()")
  plugin.State = STATE_RUNNING
  go func(){
    firstRun(plugin)

    if plugin.State == STATE_RUNNING{
      plugin.run() //start the mainloop if there was no error
    }
  }()
}

// Called in initialisation to run the code and handle errors
//
//
func firstRun(plugin *Plugin) {
  logging.Info("plugin", "plugin.firstRun()")
  _, err := plugin.VM.Run(plugin.Code)
  if err != nil{
    plugin.State = STATE_CODE_ERROR
    plugin.Error = err
  }
}
