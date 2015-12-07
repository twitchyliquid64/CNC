package exec

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/data"
)

// Called by all factory methods to initialise
//
//
func initialise(plugin *Plugin) {
  plugin.State = STATE_RUNNING
  go func(){
    defer func() {
        if caught := recover(); caught != nil {
          if caught != "plugin.exit()" {
            logging.Error("plugin-mainloop", "Panic: ", caught)
          }
          plugin.IsCurrentlyInExecution = false
          plugin.State = STATE_STOPPED
          if plugin.Model.ID != 0 {
            if caught == "plugin.exit()" {
              plugin.Model.Enabled = false
            } else {
              plugin.Model.ErrorStr = "Mainloop Panic"
              plugin.Model.HasCrashed = true
            }
            plugin.Model.Resources = nil

            data.DB.Save(&(plugin.Model))
          }
          return
        }
        return
    }()

    LoadBuiltinFunction(plugin)//populates the namespace with API methods

    if plugin.Model.ID != 0 {
      plugin.Model.ErrorStr = ""
      plugin.Model.HasCrashed = false
      plugin.Model.Resources = nil
      data.DB.Save(&(plugin.Model))
    }

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
  plugin.IsCurrentlyInExecution = true
  _, err := plugin.VM.Run(plugin.Code)
  plugin.IsCurrentlyInExecution = false
  if err != nil{
    plugin.State = STATE_CODE_ERROR
    plugin.Error = err

    if plugin.Model.ID != 0 {
      plugin.Model.ErrorStr = err.Error()
      plugin.Model.HasCrashed = true
      plugin.Model.Resources = nil
      data.DB.Save(&(plugin.Model))
      logging.Error("plugin-"+plugin.Name, "Code Error: " + err.Error())
    }

  }
}
