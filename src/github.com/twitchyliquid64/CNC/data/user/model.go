package user

import (
  "github.com/jinzhu/gorm"
  //"database/sql"
  "time"
)

const PERM_ADMIN = "ADMIN"

type User struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time

    Username string `sql:"not null;unique"`
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
  ID int      `gorm:"primary_key"`
  UserID  int `sql:"index;unique"`
  Name string `sql:"index"`
}

func (m Permission)Init(DB gorm.DB) {
  //TODO: FIX
  //DB.Model(&Permission{}).AddForeignKey("user_id", "permissions(id)", "CASCADE", "RESTRICT")
}

type Email struct {
  ID int `gorm:"primary_key"`
  UserID  int `sql:"index"`
  Address string `sql:"not null;unique"`
}

type Address struct {
    ID       int
    UserID  int `sql:"index"`
    Address1 string         `sql:"not null;unique"` // Set field as not nullable and unique
    Address2 string         `sql:"unique"`
    Postcode int
}
