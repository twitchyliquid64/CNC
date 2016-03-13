package builtin

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/robertkrimen/otto"
  "github.com/bluele/slack"
)

//Instructs the bot to join a slack channel - called from JS with slack.join()
// First param: Bot Token
// Second param: Channel name without #
func function_slack_join(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  tok := call.Argument(0).String()
  chanName := call.Argument(1).String()

  //do shit
  api := slack.New(tok)
  err := api.JoinChannel(chanName)
  if err != nil {
    logging.Error("builtin-slack", err)
  }

  return otto.Value{}
}

//Instructs the bot to send a message in a channel
// First param: Bot Token
// Second param: Channel name without #
// Third param: Message
func function_slack_send(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  tok := call.Argument(0).String()
  chanName := call.Argument(1).String()
  msg := call.Argument(2).String()

  out, err := call.Argument(3).Export()
  var options map[string]interface{}
  var ok bool
  options, ok = out.(map[string]interface{})
  if !ok {
    options = map[string]interface{}{}
  }

  var opt slack.ChatPostMessageOpt

  if _, present := options["icon"]; present {
    opt.IconUrl = options["icon"].(string)
  }
  if _, present := options["parse"]; present {
    opt.Parse = options["parse"].(string)
  }
  if _, present := options["pretext"]; present {
    opt.Attachments = append(opt.Attachments, &slack.Attachment{ Pretext: options["pretext"].(string)})
  }

  //logging.Info("builtin-slack", tok, ":", chanName, ":", msg)
  //do shit
  api := slack.New(tok)
  c, err := api.FindChannelByName(chanName)
  if err != nil {
    logging.Error("builtin-slack", "Error on channel lookup: ", err)
    return otto.Value{}
  }
  err = api.ChatPostMessage(c.Id, msg, &opt)
  if err != nil {
    logging.Error("builtin-slack", "Error on send: ", err)
  }

  return otto.Value{}
}
