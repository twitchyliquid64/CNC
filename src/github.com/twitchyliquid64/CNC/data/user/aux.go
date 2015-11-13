package user

import (
  //"github.com/twitchyliquid64/CNC/logging"
  "github.com/jinzhu/gorm"
)

func DeleteByUsername(db gorm.DB, username string) error {
  return db.Where("username = ?", username).Delete(&User{}).Error
}


var USERID_AND_THEN_NAME_FILTER = "user_id = ? and name = ?"
var USERID_AND_THEN_METHODTYPE_FILTER = "user_id = ? and method_type = ?"
