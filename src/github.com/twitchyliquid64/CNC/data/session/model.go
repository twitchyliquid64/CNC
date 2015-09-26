package session

import (
  "time"
)

const DEFAULT_KEY_SIZE = 14
const DEFAULT_EXPIRY = time.Hour * 24 * 20

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
  return (time.Now().Before(inst.Expires)) && (!inst.Revoked)
}
