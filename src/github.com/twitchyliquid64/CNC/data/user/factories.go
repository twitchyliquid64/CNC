package user

import (
  //"github.com/twitchyliquid64/CNC/logging"
  "github.com/jinzhu/gorm"
)


func GetAll(db gorm.DB)[]User{
  var users = make([]User, 0)
  db.Find(&users)

  for i := 0; i < len(users); i++ {
    loadBasicWeakEntities(&(users[i]), db)
  }

  return users
}


func GetByUsername(username string, db gorm.DB)(success bool, ret User) {
  db.Where(&User{Username:  username}).First(&ret)

  if ret.Username != username {
    return false, User{}
  }

  loadBasicWeakEntities(&ret, db)
  return true, ret
}


func GetUser(id int, db gorm.DB)*User {
  var usr User
  db.First(&usr, uint(id))

  if usr.ID == uint(id){
    loadBasicWeakEntities(&usr, db)
    return &usr
  } else { //DB fetch did not work
    return nil
  }
}

//loads all AuthMethods and Permissions
func loadBasicWeakEntities(usr *User, db gorm.DB){
  db.Model(&usr).Related(&usr.AuthMethods)
  db.Model(&usr).Related(&usr.Permissions)
}
