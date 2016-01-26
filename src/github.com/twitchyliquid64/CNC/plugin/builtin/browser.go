package builtin

import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/headzoo/surf/agent"
  "github.com/robertkrimen/otto"
  "github.com/headzoo/surf"
  "bytes"
)

// Called when JS code executes browser.open()
//
//
func function_browser_new(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  bow := surf.NewBrowser()

  v, err := plugin.VM.Call("new Object", nil)
  if err != nil {
    logging.Error("builtin-browser", err.Error())
  }
  obj := v.Object()

  obj.Set("open", func(in otto.FunctionCall)otto.Value{
    err := bow.Open(in.Argument(0).String())
    if err == nil {
        return otto.Value{}
    }

    ret, _ := otto.ToValue(err.Error())
    return ret
  })






  obj.Set("find", func(in otto.FunctionCall)otto.Value{
    expr := bow.Find(in.Argument(0).String())

    v, err := plugin.VM.Call("new Object", nil)
    if err != nil {
      logging.Error("builtin-browser", err.Error())
    }
    obj := v.Object()

    obj.Set("text", func(in otto.FunctionCall)otto.Value{
      ret, _ := otto.ToValue(expr.Text())
      return ret
    })
    obj.Set("html", func(in otto.FunctionCall)otto.Value{
      content, e := expr.Html()
      if e != nil {
        logging.Error("builtin-browser", "selector.HTML(): " + err.Error())
      }
      ret, _ := otto.ToValue(content)
      return ret
    })

    return obj.Value()
  })








  obj.Set("form", func(in otto.FunctionCall)otto.Value{
    form, err := bow.Form(in.Argument(0).String())
    if err != nil {
      logging.Error("builtin-browser", err.Error())
      return otto.Value{}
    }

    v, err := plugin.VM.Call("new Object", nil)
    if err != nil {
      logging.Error("builtin-browser", err.Error())
    }
    obj := v.Object()

    obj.Set("set", func(in otto.FunctionCall)otto.Value{
      e := form.Input(in.Argument(0).String(), in.Argument(1).String())
      if e != nil{
        logging.Error("builtin-browser", "form.set(): " + err.Error())
        ret, _ := otto.ToValue(e.Error())
        return ret
      }
      return otto.TrueValue()
    })

    obj.Set("submit", func(in otto.FunctionCall)otto.Value{
      e := form.Submit()
      if e != nil{
        logging.Error("builtin-browser", "form.submit(): " + err.Error())
        ret, _ := otto.ToValue(e.Error())
        return ret
      }
      return otto.TrueValue()
    })

    return obj.Value()
  })





  obj.Set("getCookies", func(in otto.FunctionCall)otto.Value{
    c := bow.SiteCookies()
    ret, e := plugin.VM.ToValue(c)
    if e != nil {
      logging.Error("builtin-browser", "getCookies(): " + e.Error())
    }
    return ret
  })






  obj.Set("getCookies", func(in otto.FunctionCall)otto.Value{
    c := bow.SiteCookies()
    ret, e := plugin.VM.ToValue(c)
    if e != nil {
      logging.Error("builtin-browser", "getCookies(): " + e.Error())
    }
    return ret
  })



  obj.Set("title", func(in otto.FunctionCall)otto.Value{
    ret, _ := otto.ToValue(bow.Title())
    return ret
  })
  obj.Set("body", func(in otto.FunctionCall)otto.Value{
    ret, _ := otto.ToValue(bow.Body())
    return ret
  })
  obj.Set("bodyRaw", func(in otto.FunctionCall)otto.Value{
    buf := new(bytes.Buffer)
    bow.Download(buf)
    ret, _ := otto.ToValue(buf.String())
    return ret
  })


  obj.Set("setUserAgent", func(in otto.FunctionCall)otto.Value{
    bow.SetUserAgent(in.Argument(0).String())
    return otto.Value{}
  })
  obj.Set("setChromeAgent", func(in otto.FunctionCall)otto.Value{
    bow.SetUserAgent(agent.Chrome())
    return otto.Value{}
  })
  obj.Set("setFirefoxAgent", func(in otto.FunctionCall)otto.Value{
    bow.SetUserAgent(agent.Firefox())
    return otto.Value{}
  })


  return obj.Value()
}
