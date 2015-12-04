package stmdata


import (
  "time"
)

type Stmdata struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    PluginID int `sql:"not null;index:idx_plugin_name"`

    Name string `sql:"not null;index:idx_plugin_name"`
    Data []byte
    Content string `sql:"-"`
}
