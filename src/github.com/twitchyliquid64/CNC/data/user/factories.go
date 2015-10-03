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

//loads all Permissions - called in factory methods to populate key fields
func loadBasicWeakEntities(usr *User, db gorm.DB){
  db.Model(&usr).Related(&usr.Permissions)
}

//called for login verification only - no need to get authentication methods otherwise
//should be called be user if auth methods are required
func LoadAuthMethods(usr *User, db gorm.DB){
  db.Model(&usr).Related(&usr.AuthMethods)
}

//should be called prior to using Address or Email fields.
func LoadEphemeral(usr *User, db gorm.DB){
  db.Model(&usr).Related(&usr.Emails)
  db.Model(&usr).Related(&usr.Addresses)

  //DB engine (gorm) does not keep track of which elements are 'MainAddress' / 'MainEmail'
  //hence, we need to find the first one and use that to populate the field.
  if (len(usr.Emails) > 0) {
    usr.MainEmail = usr.Emails[0]
    usr.Emails = usr.Emails[1:]
  }
  if (len(usr.Addresses) > 0) {
    usr.MainAddress = usr.Addresses[0]
    usr.Addresses = usr.Addresses[1:]
  }
}
