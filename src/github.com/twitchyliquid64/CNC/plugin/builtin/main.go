package builtin

import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/robertkrimen/otto"
)

//Called to load up the VM with pointers to API functions.
//Injected into plugin.exec to avoid circular dependency.
//
func LoadBuiltinsToVM(plugin *exec.Plugin)error{
  //logging
  plugin.VM.Set("onLogMessage", func(in otto.FunctionCall)otto.Value{return function_onLogMessage(plugin, in)})
  plugin.VM.Set("log", func(in otto.FunctionCall)otto.Value{return function_onLog(plugin, in)})

  //telegram
  tgram, _ := plugin.VM.Object(`telegram = {}`)
  tgram.Set("onChatJoined", func(in otto.FunctionCall)otto.Value{return function_telegram_onChatJoined(plugin, in)})
  tgram.Set("onChatLeft", func(in otto.FunctionCall)otto.Value{return function_telegram_onChatLeft(plugin, in)})
  tgram.Set("onChatMsg", func(in otto.FunctionCall)otto.Value{return function_telegram_onChatMsg(plugin, in)})
  tgram.Set("sendMsg", func(in otto.FunctionCall)otto.Value{return function_telegram_sendMsg(plugin, in)})
  plugin.VM.Set("telegram", tgram)

  //slack
  sl, _ := plugin.VM.Object(`slack = {}`)
  sl.Set("join", func(in otto.FunctionCall)otto.Value{return function_slack_join(plugin, in)})
  sl.Set("send", func(in otto.FunctionCall)otto.Value{return function_slack_send(plugin, in)})

  //entities
  ent, _ := plugin.VM.Object(`entities = {}`)
  ent.Set("onStatusUpdate", func(in otto.FunctionCall)otto.Value{return function_entities_onStatusUpdate(plugin, in)})
  ent.Set("onLocationUpdate", func(in otto.FunctionCall)otto.Value{return function_entities_onLocationUpdate(plugin, in)})
  plugin.VM.Set("entities", ent)

  //web
  web, _ := plugin.VM.Object(`web = {}`)
  web.Set("handle", func(in otto.FunctionCall)otto.Value{return function_web_handle(plugin, in)})
  plugin.VM.Set("web", web)


  //websockets
  ws, _ := plugin.VM.Object(`websockets = {}`)
  ws.Set("register", func(in otto.FunctionCall)otto.Value{return function_websockets_register(plugin, in)})
  plugin.VM.Set("websockets", ws)

  //plugin
  pl, _ := plugin.VM.Object(`plugin = {}`)
  pl.Set("getResources", func(in otto.FunctionCall)otto.Value{return function_plugin_getResources(plugin, in)})
  pl.Set("getResource", func(in otto.FunctionCall)otto.Value{return function_plugin_getResource(plugin, in)})
  pl.Set("ready", func(in otto.FunctionCall)otto.Value{return function_plugin_ready(plugin, in)})
  pl.Set("disable", func(in otto.FunctionCall)otto.Value{return function_plugin_disable(plugin, in)})
  pl.Set("getIcon", func(in otto.FunctionCall)otto.Value{return function_plugin_geticon(plugin, in)})
  pl.Set("setIcon", func(in otto.FunctionCall)otto.Value{return function_plugin_seticon(plugin, in)})
  pl.Set("delay", func(in otto.FunctionCall)otto.Value{return function_plugin_delay(plugin, in)})
  plugin.VM.Set("plugin", pl)

  //data
  d, _ := plugin.VM.Object(`data = {}`)
  d.Set("get", func(in otto.FunctionCall)otto.Value{return function_data_get(plugin, in)})
  d.Set("set", func(in otto.FunctionCall)otto.Value{return function_data_set(plugin, in)})
  plugin.VM.Set("data", d)

  //template
  template, _ := plugin.VM.Object(`template = {}`)
  template.Set("render", func(in otto.FunctionCall)otto.Value{return function_template_render(plugin, in)})
  template.Set("renderWeb", func(in otto.FunctionCall)otto.Value{return function_template_renderWeb(plugin, in)})
  plugin.VM.Set("template", template)

  //gmail
  gmail, _ := plugin.VM.Object(`gmail = {}`)
  gmail.Set("setup", func(in otto.FunctionCall)otto.Value{return function_gmail_setup(plugin, in)})
  gmail.Set("sendMessage", func(in otto.FunctionCall)otto.Value{return function_gmail_sendMessage(plugin, in)})
  plugin.VM.Set("gmail", gmail)

  //http
  http, _ := plugin.VM.Object(`http = {}`)
  http.Set("get", func(in otto.FunctionCall)otto.Value{return function_http_get(plugin, in)})
  http.Set("post", func(in otto.FunctionCall)otto.Value{return function_http_post(plugin, in)})
  http.Set("postValues", func(in otto.FunctionCall)otto.Value{return function_http_postValues(plugin, in)})
  plugin.VM.Set("http", http)

  //browser
  browser, _ := plugin.VM.Object(`browser = {}`)
  browser.Set("new", func(in otto.FunctionCall)otto.Value{return function_browser_new(plugin, in)})
  plugin.VM.Set("browser", browser)


  //cron
  cr, _ := plugin.VM.Object(`cron = {}`)
  cr.Set("schedule", func(in otto.FunctionCall)otto.Value{return function_cron_schedule(plugin, in)})
  plugin.VM.Set("cron", cr)

  //twilio (SMS)
  twilio, _ := plugin.VM.Object(`twilio = {}`)
  twilio.Set("sendSMS", func(in otto.FunctionCall)otto.Value{return function_twilio_sendSMS(plugin, in)})
  plugin.VM.Set("twilio", twilio)

  //aux
  plugin.VM.Set("testendpoint_good", func(in otto.FunctionCall)otto.Value{return function_onTestEndpointGood(plugin, in)})
  plugin.VM.Set("onTestDispatchTriggered", func(in otto.FunctionCall)otto.Value{return function_onTestDispatchTriggered(plugin, in)})

  return nil
}
