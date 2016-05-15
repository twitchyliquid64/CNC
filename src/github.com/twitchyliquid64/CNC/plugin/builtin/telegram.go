package builtin


import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/messenger"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/util"
  "github.com/Syfaro/telegram-bot-api"
  "github.com/robertkrimen/otto"
)

// Called when JS code executes telegram.onChatJoined()
// format for data returned to javascript can be found at:
// https://godoc.org/github.com/Syfaro/telegram-bot-api#Message
func function_telegram_onChatJoined(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{

  callback := util.GetFunc(call.Argument(0), plugin.VM)
  hook := TelegramHook{P: plugin, Callback: &callback, HookType: ON_CHAT_JOINED}
  plugin.RegisterHook(&hook)
  return otto.Value{}
}


// Called when JS code executes telegram.onChatMsg()
// format for data returned to javascript can be found at:
// https://godoc.org/github.com/Syfaro/telegram-bot-api#Message
func function_telegram_onChatMsg(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  callback := util.GetFunc(call.Argument(0), plugin.VM)
  hook := TelegramHook{P: plugin, Callback: &callback, HookType: ON_CHAT_MSG}
  plugin.RegisterHook(&hook)
  return otto.Value{}
}


// Called when JS code executes telegram.onChatLeft()
// format for data returned to javascript can be found at:
// https://godoc.org/github.com/Syfaro/telegram-bot-api#Message
func function_telegram_onChatLeft(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  callback := util.GetFunc(call.Argument(0), plugin.VM)
  hook := TelegramHook{P: plugin, Callback: &callback, HookType: ON_CHAT_LEFT}
  plugin.RegisterHook(&hook)
  return otto.Value{}
}


type TelegramHookTypes int
const (
  ON_CHAT_JOINED TelegramHookTypes = iota
  ON_CHAT_LEFT
  ON_CHAT_MSG
)

type TelegramHook struct {
  HookType TelegramHookTypes
  P *exec.Plugin
  Callback *otto.Value
}

func (h *TelegramHook)Destroy(){
  logging.Info("Telegramhook-" + h.Name(), "Destroy() called")
}
func (h *TelegramHook)Name()string{
  switch h.HookType {
    case ON_CHAT_MSG:
      return "telegram_CHAT_MSG"
    case ON_CHAT_LEFT:
      return "telegram_CHAT_LEFT"
    case ON_CHAT_JOINED:
      return "telegram_CHAT_JOINED"
  }
  return "telegram_OTHER" //should never happen
}

func (h *TelegramHook)Dispatch(data interface{}){
  msg := data.(tgbotapi.Message)

  val, err := h.P.VM.ToValue(msg)
  if err != nil {
    logging.Error("builtin-telegram", err.Error())
  }
  select {
  case h.P.PendingInvocations <- &exec.JSInvocation{Callback: h.Callback, Parameters: []interface{} {val}}:
    default:
  }
}

//Sends a normal chat message - called from JS with telegram.sendMsg()
// First param: chatID - typically obtained from a previous message.
// Second param: A string of text to send as a message.
func function_telegram_sendMsg(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  msg := call.Argument(1).String()
  chatID, _ := call.Argument(0).ToInteger()
  messenger.SendSimpleMessage(int(chatID), msg)
  return otto.Value{}
}
