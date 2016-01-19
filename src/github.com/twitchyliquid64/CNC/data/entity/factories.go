package entity

import (
  //"github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/util"
  "github.com/jinzhu/gorm"
)

func GetAll(db gorm.DB)[]Entity{
  var entities = make([]Entity, 0)
  db.Find(&entities)

  return entities
}

func GetLocations(id, limit int, db gorm.DB)[]EntityLocationRecord{
  var ret []EntityLocationRecord
  db.Where(&EntityLocationRecord{EntityID:  id}).Order("created_at desc").Limit(limit).Find(&ret)
  return ret
}

func GetEntityById(id uint, db gorm.DB)(ret Entity,err error) {
  err = db.Where(&Entity{ID:  id}).First(&ret).Error
  return
}

func GetEntityByKey(key string, db gorm.DB)(ret Entity,err error) {
  err = db.Where("api_key = ?", key).First(&ret).Error
  return
}

func GetNumEntityEventsQueued(id int, db gorm.DB)(ret int,err error) {
  err = db.Where(&EntityEvent{ID:  id}).Count(&ret).Error
  return
}


func NewEntity(ent *Entity, usrID uint, db gorm.DB)(*Entity,error){
  ent.CreatorUserID = int(usrID)
  ent.APIKey = util.RandAlphaKey(DEFAULT_APIKEY_SIZE)
  ent.LastStatString = ""
  return ent, db.Save(ent).Error
}
