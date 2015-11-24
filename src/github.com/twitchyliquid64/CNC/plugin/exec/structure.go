package exec

import (
  "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/robertkrimen/otto"
)

type States int
const (
  STATE_INIT States = iota
  STATE_CODE_ERROR
  STATE_RUN_ERROR
  STATE_RUNNING
  STATE_STOPPED
)

type Hook interface{
  Name()string            //identifies the class of hook - eg: OnLogMsg. Should never change.
  Destroy()
  Dispatch(interface{})   //Called by plugin.Dispatch to actually dispatch an invocation request on its respective plugin
}

//Represents a pending invocation - typically enqueued
type JSInvocation struct{
  Data *otto.Object
  MethodName string
}

type Plugin struct {
  Name string             //Should never change and must be unique
  State States

  Hooks []Hook

  Code string
  VM *otto.Otto
  Error error             //populated when STATE_RUN_ERROR or STATE_CODE_ERROR
  IsCurrentlyInExecution bool

  //this channel is closed when the mainloop should stop
  PendingInvocations chan *JSInvocation

  //Populated when this plugin is based off one saved in the DB, and
  //updates should be made accordingly.
  Model plugin.Plugin
  Resources []plugin.Resource //these structures should not omit any data.
}

func (p *Plugin)RegisterHook(h Hook){
  p.Hooks = append(p.Hooks, h)
  err := RegisterHookFunction(p, h)
  if err != nil {
    logging.Error("plugin", "RegisterHook() ", err)
  }
}
