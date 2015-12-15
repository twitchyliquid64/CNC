package web

import (
  "github.com/twitchyliquid64/CNC/data/entity"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/data"
  "github.com/hoisie/web"
  "encoding/json"
  "strconv"
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




// Called to get the details for a specific entity ID, passing back all info in JSON.
//
//
func getEntityHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-entity", "getEntity() called unauthorized, aborting")
    return
  }

  entID, err := strconv.Atoi(ctx.Params["entityID"])
  if err != nil {
      logging.Error("web-entity", err)
  }

  ent, err := entity.GetEntityById(uint(entID), data.DB)
  if err != nil {
      logging.Error("web-entity", err)
  }

  d, err := json.Marshal(ent)
  if err != nil {
    logging.Error("web-entity", err)
  }
  ctx.ResponseWriter.Write(d)
}



// Called to update the details for a specific entity, recieving all info as JSON.
//
//
func updateEntityHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-entity", "updateEntity() called unauthorized, aborting")
    return
  }

  decoder := json.NewDecoder(ctx.Request.Body)
  var ent entity.Entity
  err := decoder.Decode(&ent)
  if err != nil {
      logging.Error("web-entity", "updateEntity() failed to decode JSON:", err)
      ctx.Abort(500, "JSON error")
      return
  }

  err = data.DB.Save(&ent).Error
  if err == nil {
      ctx.ResponseWriter.Write([]byte("GOOD"))
  } else {
      ctx.ResponseWriter.Write([]byte(err.Error()))
      logging.Error("web-entity", err)
  }
}



// Called by plugins to update their status.
//
//
func updateEntityStatusHandlerAPI(ctx *web.Context)*APIResult {
  apiKey := ctx.Params["key"]
  ent, err := entity.GetEntityByKey(apiKey, data.DB)
  if err != nil || ent.ID == 0{
    logging.Error("entity", err.Error())
    return &APIResult{Code: 400,Error: err}
  }
  logging.Info("entity", ent)

  //update entity data struct - shows the latest status
  ent.LastStatString = ctx.Params["status"]

  //apply styling to the last status if provided
  ent.LastStatIcon = ctx.Params["styleicon"]
  statusStyling := ctx.Params["style"]
  if statusStyling != "" && statusStyling != "false" {
    ent.LastStatStyle = statusStyling
    ent.LastStatMeta = ctx.Params["stylemeta"]
  } else {
    ent.LastStatStyle = ""
    ent.LastStatMeta = ""
  }
  data.DB.Save(&ent)
  entity.PublishUpdate(ent.ID, ent.LastStatString, ent.LastStatStyle, ent.LastStatMeta, ent.LastStatIcon)

  //save the data in a new statusRecord
  rec := entity.EntityStatusRecord{}
  rec.EntityID = int(ent.ID)
  rec.Status = ent.LastStatString
  rec.Voltage = atoiOrDefault(ctx.Params["voltage"], -100)
  rec.Signal = atoiOrDefault(ctx.Params["signal"], -100)
  rec.Temperature = atoiOrDefault(ctx.Params["temp"], -100)

  err = data.DB.Save(&rec).Error
  if err != nil {
    return &APIResult{Code: 400,Error: err}
  } else {
    return &APIResult{Code: 200,Data: map[string]interface{}{"success": true}}
  }
}

func atoiOrDefault(input string, defalt int)int{
  val, err := strconv.Atoi(input)
  if err != nil{
    return defalt
  }
  return val
}
