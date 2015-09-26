package user

import (
  //"github.com/twitchyliquid64/CNC/logging"
  "github.com/jinzhu/gorm"
)

func CheckAuthByPassword(username, password string, db gorm.DB)(bool,User) {
  tmp := User{}

  //TODO: Split db call for user into separate method
  db.Where(&User{Username:  username}).First(&tmp)
  db.Model(&tmp).Related(&tmp.AuthMethods)

  for _, authmethod := range tmp.AuthMethods {
    if authmethod.MethodType == AUTH_PASSWD {
      if authmethod.Value == password {
        return true, tmp
      }
    }
  }

  return false, User{}
}
