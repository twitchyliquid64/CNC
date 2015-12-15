package web

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/hoisie/web"
  "encoding/json"
)

type APIResult struct {
  Code int
  Data interface{}
  Error error
}


func apiHandler(handlerFunc func(*web.Context)(interface{}, int))func(ctx *web.Context) {
  return func(ctx *web.Context){
    result, code := handlerFunc(ctx)
    err, isError := result.(error)
    if isError{
      d, err := json.Marshal(map[string]interface{}{"error": err.Error()})
      if err != nil {
        logging.Error("web-apicore", err)
      }
      ctx.Abort(code, string(d))
    }else{
      d, err := json.Marshal(result)
      if err != nil {
        logging.Error("web-apicore", err)
      }
      ctx.Abort(code, string(d))
    }
  }
}
