package session

import (
  "github.com/twitchyliquid64/CNC/util"
  "github.com/jinzhu/gorm"
  "time"
)

var DEFAULT_KEY_SIZE = 14
var DEFAULT_EXPIRY = time.Hour * 24 * 20

type Session struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time

    Expires time.Time
    Revoked bool

    UserID  int `sql:"index"`
    Key string `sql:"index"`
    Component string
}


//Returns true if the session has not expired or been revoked
func (inst *Session)IsValid()bool {
  return (inst.Expires.Before(time.Now())) && (!inst.Revoked)
}


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


//Called from a user handler to load a session
//pass the key from the cookie, returns a session obj if valid.
func GetSession(key string, db gorm.DB)*Session{
  s := Session{}

  db.Where(&Session{Key: key}).First(&s)
  if s.Key == key {
    return &s
  } else {
    return nil
  }
}
