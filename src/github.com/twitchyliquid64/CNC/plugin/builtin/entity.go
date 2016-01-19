package builtin


import (
  "github.com/twitchyliquid64/CNC/data/entity"
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/util"
  "github.com/robertkrimen/otto"
  "strconv"
)


// Called when JS code executes entities.onStatusUpdate(<entityID>)
//
//
func function_entities_onStatusUpdate(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  callback := util.GetFunc(call.Argument(1), plugin.VM)
  entityID, _ := call.Argument(0).ToInteger()
  hook := EntityHook{EID: entityID, P: plugin, Callback: &callback, HookType: ON_STATUS_UPDATE}
  plugin.RegisterHook(&hook)
  return otto.Value{}
}

// Called when JS code executes entities.onLocationUpdate(<entityID>)
//
//
func function_entities_onLocationUpdate(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  callback := util.GetFunc(call.Argument(1), plugin.VM)
  entityID, _ := call.Argument(0).ToInteger()
  hook := EntityHook{EID: entityID, P: plugin, Callback: &callback, HookType: ON_LOCATION_UPDATE}
  plugin.RegisterHook(&hook)
  return otto.Value{}
}



type EntityHookTypes int
const (
  ON_STATUS_UPDATE EntityHookTypes = iota
  ON_LOCATION_UPDATE
)

type EntityHook struct {
  EID int64
  HookType EntityHookTypes
  P *exec.Plugin
  Callback *otto.Value
}



func (h *EntityHook)Destroy(){
  logging.Info("Entityhook-" + h.Name(), "Destroy() called")
}
func (h *EntityHook)Name()string{
  switch h.HookType {
  case ON_STATUS_UPDATE:
      return "entity_ON_UPDATE-" + strconv.Itoa(int(h.EID))
  case ON_LOCATION_UPDATE:
      return "entity_ON_LOCATION-" + strconv.Itoa(int(h.EID))
  }
  return "entity_UNKNOWN" //should never happen
}

func (h *EntityHook)Dispatch(data interface{}){
  msg := data.(entity.EntityUpdate)

  val, err := h.P.VM.ToValue(msg)
  if err != nil {
    logging.Error("builtin-entity", err.Error())
  }
  select {
  case h.P.PendingInvocations <- &exec.JSInvocation{Callback: h.Callback, Parameters: []interface{} {val}}:
    default:
  }
}
