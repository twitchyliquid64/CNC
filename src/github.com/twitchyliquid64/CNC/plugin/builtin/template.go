package builtin

import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/robertkrimen/otto"
  "text/template"
  "bytes"
)



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


func setMapDefaultIfNotExisting(mp map[string]interface{}, key, defaultValue string)map[string]interface{}{
  if _, ok := mp[key]; !ok {
    mp[key] = defaultValue
  }
  return mp
}


//Wraps content in a properly styled webpage.
//
//
func function_template_renderWeb(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  templateContent := call.Argument(0).String()
  out, err := call.Argument(1).Export()
  options := out.(map[string]interface{})
  options = setMapDefaultIfNotExisting(options, "PrimaryColour", "indigo")
  options = setMapDefaultIfNotExisting(options, "AccentColour", "amber")
  options = setMapDefaultIfNotExisting(options, "Title", "CNC")
  options = setMapDefaultIfNotExisting(options, "Icon", "wifi")
  outputSoFar := ""

  if err != nil {
    logging.Error("builtin-template", "Options Export: " + err.Error())
    return otto.Value{}
  }


  t, err := template.ParseFiles("templates/basicpage_header.tpl", "templates/basicpage_footer.tpl")
  if err != nil {
    logging.Error("builtin-template", "Template load: " + err.Error())
    return otto.Value{}
  }

  var output1 bytes.Buffer
  var output2 bytes.Buffer
  err = t.Lookup("basicpage_header.tpl").Execute(&output1, options)
  if err != nil {
    logging.Error("builtin-template", "Execute header: " + err.Error())
    return otto.Value{}
  }
  err = t.Lookup("basicpage_footer.tpl").Execute(&output2, options)
  if err != nil {
    logging.Error("builtin-template", "Execute footer: " + err.Error())
    return otto.Value{}
  }
  outputSoFar += output1.String()
  outputSoFar += templateContent
  outputSoFar += output2.String()


  res, _ := plugin.VM.ToValue(outputSoFar)
  return res
}
