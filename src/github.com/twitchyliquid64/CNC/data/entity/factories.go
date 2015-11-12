package entity

import (
  //"github.com/twitchyliquid64/CNC/logging"
  "github.com/jinzhu/gorm"
)

func GetAll(db gorm.DB)[]Entity{
  var entities = make([]Entity, 0)
  db.Find(&entities)

  return entities
}
