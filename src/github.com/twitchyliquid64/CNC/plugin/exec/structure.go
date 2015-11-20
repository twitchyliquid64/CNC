package exec

import (
  "github.com/robertkrimen/otto"
)

type States int
const (
  STATE_INIT States = iota
  STATE_CODE_ERROR
  STATE_RUN_ERROR
  STATE_RUNNING
  STATE_DISABLED
  STATE_STOPPED
)

type Hook interface{
  Name()string            //identifies the class of hook - eg: OnLogMsg. Should never change.
  Destroy()
  Plugin()*Plugin
  MethodName()string      //the method to call in javascript code when the hook fires
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
}
