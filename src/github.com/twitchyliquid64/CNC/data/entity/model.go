package entity

import (
  "time"
)

var DEFAULT_APIKEY_SIZE = 10

type Entity struct {
    ID              uint      `gorm:"primary_key"`
    CreatedAt       time.Time
    UpdatedAt       time.Time
    DeletedAt       *time.Time

    Icon            string
    CreatorUserID   int       `sql:"index"`
    Name            string    `sql:"index"`
    Category        string

    LastStatString  string
    LastStatStyle   string //only used if the update web call asserts fancy styling
    LastStatIcon    string //only used if the update web call asserts fancy styling
    LastStatMeta    string //only used if the update web call asserts fancy styling

    APIKey          string    `sql:"index"`
}


//this is the weak entity which maps the many --> many between entities and users.
//The presence of this pivot means the specified user is 'attached' to this entity,
//and can use it.
type EntityPivot struct {
  ID          int         `gorm:"primary_key"`
  UserID      int         `sql:"index"`
  EntityID    int         `sql:"index"`
}


type EntityLocationRecord struct {
  ID          int         `gorm:"primary_key"`
  EntityID    int         `sql:"index"`
  CreatedAt   time.Time   `sql:"index"`
  HasFullInfo bool //set to true if Speed, Course, sats are set

  Latitude    float64
  Longitude   float64
  SpeedKph    float64
  Course      int
  SatNum      int
  Accuracy    int
}


type EntityStatusRecord struct {
  ID          int         `gorm:"primary_key"`
  EntityID    int         `sql:"index"`
  CreatedAt   time.Time   `sql:"index"`

  Status      string
  Voltage     int         //negative -100 is the default
  Signal      int         //negative -100 is the default
  Temperature int         //negative -100 is the default
}


type EntityLogRecord struct {
  ID          int         `gorm:"primary_key"`
  EntityID    int         `sql:"index"`
  CreatedAt   time.Time   `sql:"index"`

  MessageCode int         //custom field, to be used by programmatic magic

  Level       int         //1 = error, 2 = warning, 3 = info, 4 = debug
  Component   string
  Message     string
}
