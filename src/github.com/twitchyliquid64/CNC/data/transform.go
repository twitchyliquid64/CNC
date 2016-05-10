package data

import (
  "github.com/twitchyliquid64/CNC/data/stmdata"
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/entity"
  "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/config"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "database/sql"
)

const currentDataVersion = 2
const dataVersionKey = "DataVersion"

func getDbVersionData(data *StmData, db *gorm.DB) bool {
  db.Where("name = ? AND id = ?", dataVersionKey, 0).First(&data)
  return data.ID != 0
}

func getDbVersion(db *gorm.DB) {
  var data StmData;
  if getDbVersion(&data, db) {
    return int(data.Content)
  }

  return 1;
}

func setDbVersion(version int, db *gorm.DB) {
  var data StmData
  if getDbVersionData(&data, db) {
    data.Data = []byte(version)
    db.Save(&data)
  } else {
    data.Name = dataVersionKey
    data.PluginID = 0
    data.Data = []byte(version)
  }
}

func upgradeDb(db *gorm.DB) {
  switch originalVersion := getDbVersion(); {
  case originalVersion <= 1:
    upgrade_v2(db)
    fallthrough
  //case originalVersion <= 2:
    //upgrade_v3(db)
    //fallthrough
  }

  logging.Info("Setting Current Version")
  setDbVersion(currentDataVersion, db)
}

func upgrade_v2(db *gorm.DB) {
  // Remove IsExecutable and IsTemplate from Resource
  // Add Type column

  logging.Info("Migrating 1 => 2")

  // If something is executable, it stays that way. Otherwise it becomes a template
  db.Exec("UPDATE resource SET Type = ? WHERE IsExecutable = ?", ResJavascriptCode, true)
  db.Exec("UPDATE resource SET Type = ? WHERE IsExecutable = ?", ResTemplate, false)

  db.DropColumn("IsExecutable")
  db.DropColumn("IsTemplate")
}
