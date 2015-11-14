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



// Called to create a new entity in the system, given its params by JSON.
//
//
func newEntityHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-entity", "newEntity() called unauthorized, aborting")
    return
  }

  decoder := json.NewDecoder(ctx.Request.Body)
  var ent entity.Entity
  err := decoder.Decode(&ent)
  if err != nil {
      logging.Error("web-entity", "newEntityHandlerAPI() failed to decode JSON:", err)
      ctx.Abort(500, "JSON error")
      return
  }

  _, err = entity.NewEntity(&ent, u.ID, data.DB)
  if err == nil {
      ctx.ResponseWriter.Write([]byte("GOOD"))
  } else {
      ctx.ResponseWriter.Write([]byte(err.Error()))
      logging.Error("web-entity", err)
  }
}
