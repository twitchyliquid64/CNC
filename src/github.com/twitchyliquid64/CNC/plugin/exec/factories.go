package exec

import (
  pluginData "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/robertkrimen/otto"
)

const MAX_INVOCATION_QUEUE = 40

// Simple method to construct a plugin and get it running in the system.
//
//
func BuildPlugin(name, code string)*Plugin {
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
  code := ""
  for _, resource := range res {
    if resource.IsJavascriptCode() {
      code += "\n" + string(resource.Data)
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
