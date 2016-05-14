package data

import (
  "github.com/twitchyliquid64/CNC/data/stmdata"
  "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "strconv"
)

const currentDataVersion = 2
const dataVersionKey = "DataVersion"

func getDbVersionData(data *stmdata.Stmdata, db gorm.DB) bool {
  db.Where("name = ? AND id = ?", dataVersionKey, 0).First(&data)
  return data.ID != 0
}

func getDbVersion(db gorm.DB) int {
  if data, exists := stmdata.Get(0, dataVersionKey, db); exists {
    version, err := strconv.Atoi(string(data.Data))
    if err != nil {
      panic("Cant determine version")
    }

    return version
  }

  return 1
}

func setDbVersion(version int, db gorm.DB) {
  stmdata.Set(0, dataVersionKey, strconv.Itoa(version), db)
}

func upgradeDb(db gorm.DB) {
  switch originalVersion := getDbVersion(db); {
  case originalVersion <= 1:
    upgrade_v2(db)
    //fallthrough
  //case originalVersion <= 2:
    //upgrade_v3(db)
    //fallthrough
  }

  logging.Info("Setting Current Version")
  setDbVersion(currentDataVersion, db)
}

func upgrade_v2(db gorm.DB) {
  // Remove IsExecutable and IsTemplate from Resource
  // Add Type column

  logging.Info("Migrating 1 => 2")

  db.Exec(`ALTER TABLE "resources" ADD COLUMN res_type char(3)`)

  // If something is executable, it stays that way. Otherwise it becomes a template
  db.Exec("UPDATE resources SET res_type = ? WHERE is_executable = ?", plugin.ResJavascriptCode, true)
  db.Exec("UPDATE resources SET res_type = ? WHERE is_executable = ?", plugin.ResTemplate, false)

  db.Exec(`ALTER TABLE "resources" ALTER COLUMN res_type SET NOT NULL`)

  db.Model(&plugin.Resource{}).DropColumn("is_executable")
  db.Model(&plugin.Resource{}).DropColumn("is_template")
}
