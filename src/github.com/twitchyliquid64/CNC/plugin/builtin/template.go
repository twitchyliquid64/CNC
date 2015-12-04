package builtin

import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/robertkrimen/otto"
  "text/template"
  "bytes"
)


//Sends a normal chat message - called from JS with telegram.sendMsg()
// First param: chatID - typically obtained from a previous message.
// Second param: A string of text to send as a message.
func function_template_render(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  templateContent := call.Argument(0).String()
  templateNamespace, err := call.Argument(1).Export()

  if err != nil {
    logging.Error("builtin-template", "Process namespace: " + err.Error())
    return otto.Value{}
  }

  tmpl, err := template.New("main").Parse(templateContent)
  if err != nil {
    logging.Error("builtin-template", "Parse: " + err.Error())
    return otto.Value{}
  }

  var output bytes.Buffer
  err = tmpl.Execute(&output, templateNamespace)
  if err != nil {
    logging.Error("builtin-template", "Execute: " + err.Error())
    return otto.Value{}
  }

  res, _ := plugin.VM.ToValue(output.String())
  return res
}
