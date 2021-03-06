package builtin

import (
  "github.com/twitchyliquid64/CNC/web/pluginhandler"
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/util"
  "github.com/robertkrimen/otto"
)

const WEB_HOOK_PREFIX = "web_"
const HANDLER_ID_LENGTH = 12


// Called when JS code executes web.handle()
//
//
func function_web_handle(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  patternRegex := call.Argument(0).String()
  callback := util.GetFunc(call.Argument(1), plugin.VM)
  addHTTPException, _ := call.Argument(2).ToBoolean()

  hookID := util.RandAlphaKey(HANDLER_ID_LENGTH)
  hook := WebHook{P: plugin, Callback: &callback, HookID: hookID, Pattern: patternRegex}
  plugin.RegisterHook(&hook)
  if pluginhandler.AddHook(hook.Name(), patternRegex, addHTTPException) {
    return otto.TrueValue()
  } else {
    return otto.FalseValue()
  }
}



type Req interface{
  Write(string)
  Done()
  URL()string
  Parameter(string)string
  PostBody()string
  LoggedIn()bool
  User()*user.User
  Session()*session.Session
}

type WebHook struct {
  Pattern string
  HookID string
  P *exec.Plugin
  Callback *otto.Value
}

func (h *WebHook)Destroy(){
  logging.Info(h.Name(), "hook.Destroy() called")
  pluginhandler.RemoveHook(h.Name())
}
func (h *WebHook)Name()string{
  return WEB_HOOK_PREFIX + h.HookID
}
func (h *WebHook)Dispatch(data interface{}){
  rObj, ok := data.(Req)
  if !ok{
    logging.Error("builtin-web", "Failed to coerce interface to Request interface, aborting dispatch")
    return
  }

  v, err := h.P.VM.Call("new Object", nil)
  if err != nil {
    logging.Error("builtin-web", err.Error())
  }
  obj := v.Object()

  obj.Set("write", func(in otto.FunctionCall)otto.Value{
    rObj.Write(in.Argument(0).String())
    return otto.Value{}
  })

  obj.Set("done", func(in otto.FunctionCall)otto.Value{
    rObj.Done()
    return otto.Value{}
  })

  obj.Set("url", rObj.URL())
  obj.Set("data", rObj.PostBody())
  obj.Set("parameter", func(in otto.FunctionCall)otto.Value{
    ret, _ := otto.ToValue(rObj.Parameter(in.Argument(0).String()))
    return ret
  })

  obj.Set("isLoggedIn", func(in otto.FunctionCall)otto.Value{
    if rObj.LoggedIn() {
      return otto.TrueValue()
    } else {
      return otto.FalseValue()
    }
  })
  obj.Set("user", func(in otto.FunctionCall)otto.Value{
    ret, err := h.P.VM.ToValue(rObj.User())
    if err != nil {
      logging.Error("builtin-web", err.Error())
    }
    return ret
  })
  obj.Set("session", func(in otto.FunctionCall)otto.Value{
    ret, _ := h.P.VM.ToValue(rObj.Session())
    return ret
  })


  select {
  case h.P.PendingInvocations <- &exec.JSInvocation{Callback: h.Callback, Parameters: []interface{} { obj }}:
    default:
  }
}
