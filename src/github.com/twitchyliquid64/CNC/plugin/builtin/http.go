package builtin


import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/robertkrimen/otto"
  "io/ioutil"
  "net/http"
  "net/url"
  "strings"
)



func function_http_get(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  addr := call.Argument(0).String()

  resp, err := http.Get(addr)
  if err != nil {
  	logging.Warning("plugin-http", err.Error())
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    logging.Error("builtin-http", err.Error())
  }

  val, err := plugin.VM.ToValue(
    struct{
      Code int
      CodeStr string
      Data string
      Addr string
    }{Data: string(body), Addr: addr,
      Code: resp.StatusCode, CodeStr: resp.Status})
  if err != nil {
    logging.Error("builtin-http", err.Error())
  }

  return val
}



func function_http_post(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  addr := call.Argument(0).String()
  contentType := call.Argument(1).String()
  data := call.Argument(2).String()

  resp, err := http.Post(addr, contentType, strings.NewReader(data))
  if err != nil {
  	logging.Warning("plugin-http", err.Error())
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    logging.Error("builtin-http", err.Error())
  }

  val, err := plugin.VM.ToValue(
    struct{
      Code int
      CodeStr string
      Data string
      Addr string
    }{Data: string(body), Addr: addr,
      Code: resp.StatusCode, CodeStr: resp.Status})
  if err != nil {
    logging.Error("builtin-http", err.Error())
  }

  return val
}




func function_http_postValues(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  addr := call.Argument(0).String()
  paramsObj := call.Argument(1).Object()

  urlValues := url.Values{}

  for _, key := range paramsObj.Keys() {
    v, _ := paramsObj.Get(key)
    d, _ := v.ToString()
    urlValues.Add(key, d)
  }

  resp, err := http.PostForm(addr, urlValues)
  if err != nil {
  	logging.Warning("plugin-http", err.Error())
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    logging.Error("builtin-http", err.Error())
  }

  val, err := plugin.VM.ToValue(
    struct{
      Code int
      CodeStr string
      Data string
      Addr string
    }{Data: string(body), Addr: addr,
      Code: resp.StatusCode, CodeStr: resp.Status})
  if err != nil {
    logging.Error("builtin-http", err.Error())
  }

  return val
}
