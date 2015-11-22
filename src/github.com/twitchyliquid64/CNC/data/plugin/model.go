package plugin

import (
  "time"
)

type Plugin struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time

    Name string `sql:"not null;unique;index"`
    Icon string
    Description string
    Enabled bool

    HasCrashed bool
    ErrorStr string

    Resources []Resource
}


type Resource struct {
  ID int      `gorm:"primary_key"`
  PluginID int `sql:"index"`
  Name string `sql:"index"`
  Data string `sql:"type:text"`
  IsExecutable bool
  IsTemplate bool
}
