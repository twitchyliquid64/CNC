package web

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/hoisie/web"
  "encoding/json"
)

func apiHandler(handlerFunc func(*web.Context)(interface{}, int))func(ctx *web.Context) {
  return func(ctx *web.Context){
    result, code := handlerFunc(ctx)
    body, err := getJsonBody(result)
    if err != nil {
      logging.Error("web-apicore", err)
    }

    ctx.Abort(code, body)
  }
}

func getJsonBody(body interface{}) (string, error) {
  if body == nil {
    return "{}", nil
  }

  itemToSerialise := body
  if err, isError:= body.(error); isError {
    itemToSerialise = map[string]interface{}{"error": err.Error()}
  }

  d, err := json.Marshal(itemToSerialise)
  return string(d), err
}
