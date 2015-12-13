package util

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/robertkrimen/otto"
)

func GetFunc(argument otto.Value, vm *otto.Otto) otto.Value {
  if (argument.IsFunction()) {
    return argument;
  }

  mname := argument.String()
  logging.Warning("plugin-builtin", "Using a string to name parameters is now deprecated. Please use callbacks directly: " + mname)

  method, err := vm.Get(mname)
  if (err != nil){
    logging.Error("plugin-builtin", "Error getting function \"" + mname + "\": ")
  }

  return method;
}
