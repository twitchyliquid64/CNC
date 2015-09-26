package user

import (
  //"github.com/twitchyliquid64/CNC/logging"
  "github.com/jinzhu/gorm"
)

//Checks if a usename and password pair are valid, returning true and the user
//if it is.
func CheckAuthByPassword(username, password string, db gorm.DB)(bool,User) {
  success, tmp := GetByUsername(username, db)

  if !success {
    return false, tmp
  }

  for _, authmethod := range tmp.AuthMethods {
    if authmethod.MethodType == AUTH_PASSWD {
      if authmethod.Value == password {
        return true, tmp
      }
    }
  }

  return false, User{}
}
