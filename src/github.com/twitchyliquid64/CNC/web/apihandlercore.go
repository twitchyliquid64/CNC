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


func apiHandler(handlerFunc func(*web.Context)*APIResult)func(ctx *web.Context) {
  return func(ctx *web.Context){
    result := handlerFunc(ctx)
    if result.Error == nil{
      d, err := json.Marshal(result.Data)
      if err != nil {
        logging.Error("web-apicore", err)
      }
      ctx.ResponseWriter.Write(d)
    }else{//write error
      d, err := json.Marshal(map[string]interface{}{"error": result.Error.Error(), "success": false})
      if err != nil {
        logging.Error("web-apicore", err)
      }
      ctx.Abort(result.Code, string(d))
    }
  }
}
