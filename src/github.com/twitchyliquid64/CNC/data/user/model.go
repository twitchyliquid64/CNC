package user

import (
  "github.com/jinzhu/gorm"
  "database/sql"
  "time"
)

type User struct {
    gorm.Model

    Username string
    Birth time.Time
    Firstname string
    Lastname string

    MainEmail Email
    MainAddress Address

    Addresses []Address
    Emails []Email
    Permissions []Permission
}

type Permission struct {
  gorm.Model
  Name string `sql:"index:idx_name_code"`
}

type Email struct {
  Address string
}

type Address struct {
    ID       int
    Address1 string         `sql:"not null;unique"` // Set field as not nullable and unique
    Address2 string         `sql:"type:varchar(100);unique"`
    Post     sql.NullString `sql:"not null"`
}
