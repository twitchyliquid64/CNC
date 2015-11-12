package web

import (
  "github.com/twitchyliquid64/CNC/data/entity"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/data"
  "github.com/hoisie/web"
  "encoding/json"
)

// Passes back a JSON array of all entities
//
//
func getAllEntitiesHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-entity", "getAllEntities() called unauthorized, aborting")
    return
  }

  entities := entity.GetAll(data.DB)

  d, err := json.Marshal(entities)
  if err != nil {
    logging.Error("web-entity", err)
  }
  ctx.ResponseWriter.Write(d)
}
