package session

import (
  "github.com/jinzhu/gorm"
)


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
