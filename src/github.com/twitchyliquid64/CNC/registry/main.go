package registry

// This package holds information about internal state of the system - both data oriented (current system components
// and their state) and channel oriented (hooks, publishers, subscribers)
//
// This is different from the data package in that it stores non-persistant state, not useful data.

import (
  "github.com/twitchyliquid64/CNC/logging"
)


func Initialise(){
  logging.Info("registry", "Initialise()")
}
