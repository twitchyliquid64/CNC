package web

import (
  "github.com/twitchyliquid64/CNC/data/entity"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/registry"
  "github.com/twitchyliquid64/CNC/data"
  "github.com/hoisie/web"
  "encoding/json"
  "strconv"
  "time"
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



func getNumEntityEventsQueued(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-entity", "getNumEntityEventsQueued() called unauthorized, aborting")
    return
  }

  entID, err := strconv.Atoi(ctx.Params["entityID"])
  if err != nil {
      logging.Error("web-entity", err)
  }

  ent, err := entity.GetNumEntityEventsQueued(int(entID), data.DB)
  if err != nil {
      logging.Error("web-entity", err)
  }

  ctx.ResponseWriter.Write([]byte(strconv.Itoa(ent)))
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



// Called by entities to get an event from the event queue.
//
//
func getEntityQueueAPI(ctx *web.Context)(output interface{}, code int) {
  apiKey := ctx.Params["key"]
  ent, err := entity.GetEntityByKey(apiKey, data.DB)
  if err != nil || ent.ID == 0{
    logging.Error("entity", err.Error())
    return err, 400
  }

  event, err := entity.GetPendingEntityEvent(int(ent.ID), data.DB)
  if err != nil {
    return err, 400
  }

  entity.PublishEventQueueUpdate(ent.ID, false)
  return event, 200
}


// Called by entities to get an event from the event queue, long polling till an event is available.
//
//
func getEntityQueueAPI_longpoll(ctx *web.Context)(output interface{}, code int) {
  apiKey := ctx.Params["key"]
  ent, err := entity.GetEntityByKey(apiKey, data.DB)
  if err != nil || ent.ID == 0{
    logging.Error("entity", err.Error())
    return err, 400
  }

  timeoutSecs := atoiOrDefault(ctx.Params["timeout"], 40)

  //subscribe to the channel for updates from entities
  updateMsgs := make(chan entity.EntityUpdate, 10)
  entity.SubscribeUpdates(updateMsgs)
  defer entity.UnsubscribeUpdates(updateMsgs)

  //wait for message or timeout
  gotMsg := false
  exitLoop := false

  //skip the whole loop thing if there are already things in the database
  numEvents, _ := entity.GetNumEntityEventsQueued(int(ent.ID), data.DB)
  if numEvents > 0 {
    gotMsg = true
    exitLoop = true
  }

  for !exitLoop{
    select {
    case <-time.After(time.Duration(timeoutSecs) * time.Second):
        exitLoop = true
      case msg := <-updateMsgs:
        if msg.EntityID == ent.ID && msg.Type == entity.Updatetype_EventQueue_Increment {
          gotMsg = true
          exitLoop = true
        }
    }
  }

  if gotMsg{
    event, err := entity.GetPendingEntityEvent(int(ent.ID), data.DB)
    if err != nil {
      return err, 400
    }

    entity.PublishEventQueueUpdate(ent.ID, false)
    return event, 200
  }
  return map[string]string{"error": "timeout"}, 400
}

// Called by any entity to insert an event into the queue of any entity.
//
//
func insertEntityEventAPI(ctx *web.Context)(output interface{}, code int) {
  apiKey := ctx.Params["key"]
  ent, err := entity.GetEntityByKey(apiKey, data.DB)
  if err != nil || ent.ID == 0{
    logging.Error("entity", err.Error())
    return err, 400
  }

  id := atoiOrDefault(ctx.Params["id"], 0)
  typ := ctx.Params["type"]
  d := ctx.Params["data"]
  iData := atoiOrDefault(ctx.Params["int"], 0)

  _, err = entity.NewEntityEvent(id, typ, d, iData, data.DB)
  if err == nil {
    entity.PublishEventQueueUpdate(uint(id), true)
  }
  return err, 200
}

// Called by entities to update their status.
//
//
func updateEntityStatusHandlerAPI(ctx *web.Context)(output interface{}, code int) {
  apiKey := ctx.Params["key"]
  ent, err := entity.GetEntityByKey(apiKey, data.DB)
  if err != nil || ent.ID == 0{
    logging.Error("entity", err.Error())
    return err, 400
  }

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
  updatePkt := entity.PublishStatusUpdate(ent.ID, ent.LastStatString, ent.LastStatStyle, ent.LastStatMeta, ent.LastStatIcon)

  registry.DispatchEvent("entity_ON_UPDATE-"+strconv.Itoa(int(updatePkt.EntityID)), updatePkt)

  //save the data in a new statusRecord
  rec := entity.EntityStatusRecord{}
  rec.EntityID = int(ent.ID)
  rec.Status = ent.LastStatString
  rec.Voltage = atoiOrDefault(ctx.Params["voltage"], -100)
  rec.Signal = atoiOrDefault(ctx.Params["signal"], -100)
  rec.Temperature = atoiOrDefault(ctx.Params["temp"], -100)

  err = data.DB.Save(&rec).Error
  if err != nil {
    return err, 400
  } else {
    return map[string]interface{}{"success": true}, 200
  }
}

func atoiOrDefault(input string, defalt int)int{
  val, err := strconv.Atoi(input)
  if err != nil{
    return defalt
  }
  return val
}

func getFloatOrDefault(input string, defalt float64)float64{
  r, err := strconv.ParseFloat(input, 64)
  if err != nil {
    return defalt
  }
  return r
}


// Called by plugins to transmit a record of location.
// HTTP Parameters:
// required: key==entity API key
// option: lat, lon, kph (speed), course, acc (accuracy), sat (number of satellites acquired)
func updateEntityLocationHandlerAPI(ctx *web.Context)(output interface{}, code int) {
  apiKey := ctx.Params["key"]
  ent, err := entity.GetEntityByKey(apiKey, data.DB)
  if err != nil || ent.ID == 0{
    logging.Error("entity", err.Error())
    return err, 400
  }

  rec := entity.EntityLocationRecord{}
  rec.EntityID = int(ent.ID)
  rec.Latitude = getFloatOrDefault(ctx.Params["lat"], -100)
  rec.Longitude = getFloatOrDefault(ctx.Params["lon"], -100)
  rec.SpeedKph = getFloatOrDefault(ctx.Params["kph"], -100)
  rec.Course = atoiOrDefault(ctx.Params["course"], -100)
  rec.Accuracy = atoiOrDefault(ctx.Params["acc"], -100)
  rec.SatNum = atoiOrDefault(ctx.Params["sat"], -100)
  if(rec.Course != -100 && rec.SpeedKph != -100) {
    rec.HasFullInfo = true
  }

  err = data.DB.Save(&rec).Error
  if err != nil {
    return err, 400
  } else {
    updatePkt := entity.PublishLocationUpdate(uint(rec.EntityID), rec.Latitude, rec.Longitude, rec.SpeedKph, rec.Accuracy, rec.Course, rec.SatNum)
    registry.DispatchEvent("entity_ON_LOCATION-"+strconv.Itoa(int(updatePkt.EntityID)), updatePkt)

    return map[string]interface{}{"success": true}, 200
  }
}




// Passes back a JSON array of all location updates for a given entity.
//
//
func getEntityLocationsHandlerAPI(ctx *web.Context)(interface{}, int) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-entity", "getEntityLocations() called unauthorized, aborting")
    return nil, 403
  }

  eID, err := strconv.Atoi(ctx.Params["id"])
  if err != nil{
    logging.Error("web-entity", "getEntityLocations() called without 'id' parameter, aborting")
    return nil, 400
  }
  limit := atoiOrDefault(ctx.Params["limit"], 1)

  records := entity.GetLocations(eID, limit, data.DB)
  return records, 200
}
