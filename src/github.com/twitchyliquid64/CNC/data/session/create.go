package session

import (
  "github.com/twitchyliquid64/CNC/util"
  "github.com/jinzhu/gorm"
  "time"
)

//Called when a user logged in to create a new session
//pass the UserID (PK in the DB) and a string describing what
//they are logging in to (ie: web, IRC etc)
func CreateSession(userID int, component string, db gorm.DB)string {
  s := Session{
    Key: util.RandAlphaKey(DEFAULT_KEY_SIZE),
    UserID: userID,
    Component: component,
    Expires: time.Now().Add(DEFAULT_EXPIRY),
    Revoked: false,
  }
  db.Create(&s)
  return s.Key
}
