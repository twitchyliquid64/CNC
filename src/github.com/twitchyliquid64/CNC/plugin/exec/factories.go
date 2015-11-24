package exec

import (
  pluginData "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/robertkrimen/otto"
)

const MAX_INVOCATION_QUEUE = 10

// Simple method to construct a plugin and get it running in the system.
//
//
func BuildPlugin(name, code string)*Plugin {
  logging.Info("plugin", "BuildPlugin()")
  p := &Plugin{
    Name: name,
    Code: code,
    State: STATE_INIT,
    VM: otto.New(),
    PendingInvocations: make(chan *JSInvocation, MAX_INVOCATION_QUEUE),
    IsCurrentlyInExecution: false,
  }
  p.VM.Interrupt = make(chan func(), 1)

  initialise(p)
  return p
}


// Simple method to construct a plugin and get it running in the system.
//
//
func BuildPluginFromDatabase(name string, plugin pluginData.Plugin, res []pluginData.Resource)*Plugin {
  logging.Info("plugin", "BuildPluginFromDatabase()")

  code := ""
  for _, resource := range res {
    if resource.IsJavascriptCode() {
      code += "\n" + resource.Data
    }
  }

  p := &Plugin{
    Name: name,
    Code: code,
    State: STATE_INIT,
    VM: otto.New(),
    PendingInvocations: make(chan *JSInvocation, MAX_INVOCATION_QUEUE),
    IsCurrentlyInExecution: false,
    Model: plugin,
    Resources: res,
  }
  p.VM.Interrupt = make(chan func(), 1)

  initialise(p)
  return p
}
