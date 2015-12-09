package builtin

import (
  "github.com/twitchyliquid64/CNC/web/pluginsockets"
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/util"
  "github.com/robertkrimen/otto"
)


const WEBSOCK_PREFIX = "ws_event_"
const WSHANDLER_ID_LENGTH = 12


// Called when JS code executes websocket.handle()
//
//
func function_websockets_register(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  patternRegex  := call.Argument(0).String()
  onOpenMethod  := call.Argument(1).String()
  onCloseMethod := call.Argument(2).String()
  onMsgMethod   := call.Argument(3).String()

  hookID := util.RandAlphaKey(WSHANDLER_ID_LENGTH)
  hook := WebsocketHook{P: plugin,
                        OnOpen: onOpenMethod,
                        OnClose: onCloseMethod,
                        OnMsg: onMsgMethod,
                        HookID: hookID,
                        Pattern: patternRegex}
  plugin.RegisterHook(&hook)

  if pluginsockets.AddHook(hook.Name(), hook.Name(), hook.Name(), patternRegex) { //use the same hook dispatch for all three event types
    return otto.TrueValue()
  } else {
    return otto.FalseValue()
  }
}

type WebsocketHook struct {
  Pattern string
  HookID string
  P *exec.Plugin

  OnOpen string
  OnClose string
  OnMsg string
}

type WebSock interface{
  Write(string)
  URL()string
  Parameter(string)string

  LoggedIn()bool
  User()*user.User
  Session()*session.Session
  GetID()uint64
  Addr()string
  Close()
}

type SocketEvent interface {
  Event() string
  GetData() string
  Sock() interface{}
}





func (h *WebsocketHook)Destroy(){
  logging.Info(h.Name(), "hook.Destroy() called")
  pluginsockets.RemoveHook(h.Name())
}

func (h *WebsocketHook)Name()string{
  return WEBSOCK_PREFIX + h.HookID
}


func (h *WebsocketHook)Dispatch(data interface{}){
  event := data.(SocketEvent)
  sock := event.Sock().(WebSock)

  jsObj := h.genSocketObj(event, sock)
  if event.Event() == "OPEN" {
    logging.Info(h.Name(), "Dispatch() OPEN")
    h.P.PendingInvocations <- &exec.JSInvocation{MethodName: h.OnOpen, Data: jsObj}
  } else if event.Event() == "MSG" {
    logging.Info(h.Name(), "Dispatch() MSG")
    jsObj.Set("data", event.GetData())
    h.P.PendingInvocations <- &exec.JSInvocation{MethodName: h.OnMsg, Data: jsObj}
  } else if event.Event() == "CLOSE" {
    logging.Info(h.Name(), "Dispatch() CLOSE")
    h.P.PendingInvocations <- &exec.JSInvocation{MethodName: h.OnClose, Data: jsObj}
  }
}


func (h *WebsocketHook)genSocketObj(e SocketEvent, s WebSock)*otto.Object {
  v, err := h.P.VM.Call("new Object", nil)
  if err != nil {
    logging.Error("builtin-ws", err.Error())
  }
  obj := v.Object()

  obj.Set("write", func(in otto.FunctionCall)otto.Value{
    s.Write(in.Argument(0).String())
    return otto.Value{}
  })

  obj.Set("url", s.URL())
  obj.Set("id", s.GetID())
  obj.Set("addr", s.Addr())
  obj.Set("close", func(in otto.FunctionCall)otto.Value{
    s.Close()
    return otto.Value{}
  })
  obj.Set("parameter", func(in otto.FunctionCall)otto.Value{
    ret, _ := otto.ToValue(s.Parameter(in.Argument(0).String()))
    return ret
  })

  obj.Set("isLoggedIn", func(in otto.FunctionCall)otto.Value{
    if s.LoggedIn() {
      return otto.TrueValue()
    } else {
      return otto.FalseValue()
    }
  })
  obj.Set("user", func(in otto.FunctionCall)otto.Value{
    ret, err := h.P.VM.ToValue(s.User())
    if err != nil {
      logging.Error("builtin-ws", err.Error())
    }
    return ret
  })
  obj.Set("session", func(in otto.FunctionCall)otto.Value{
    ret, _ := h.P.VM.ToValue(s.Session())
    return ret
  })

  return obj
}
