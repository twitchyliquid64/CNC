package user

import (
  //"github.com/twitchyliquid64/CNC/logging"
  "github.com/jinzhu/gorm"
)

func GetByUsername(username string, db gorm.DB)(success bool, ret User) {
  db.Where(&User{Username:  username}).First(&ret)

  if ret.Username != username {
    return false, User{}
  }

  db.Model(&ret).Related(&ret.AuthMethods)

  return true, ret
}


func GetUser(id int, db gorm.DB)*User {
  var usr User
  db.First(&usr, uint(id))

  if usr.ID == uint(id){
    return &usr
  } else {
    return nil
  }
}
