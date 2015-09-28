package session

import (
  "github.com/jinzhu/gorm"
  "time"
)

func Delete(s *Session, db gorm.DB) {
  s.Expires = time.Now()
  db.Save(s)
}
