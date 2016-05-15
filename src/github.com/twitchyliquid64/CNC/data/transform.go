package data

import (
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/data/stmdata"
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/entity"
  "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "strconv"
  "fmt"
)

const currentDataVersion = 2
const dataVersionKey = "DataVersion"

//called during initialisation. Should make sure the schema is intact and up to date.
func checkStructures(db gorm.DB) {
  version := getDbVersion(db);
  if version < currentDataVersion {
    logging.Info("data", fmt.Sprintf("Migrating from %v to %v", version, currentDataVersion))

    logging.Info("data", "Auto migrating tables")
    autoMigrateTables()

    logging.Info("data", "Migrating DB")
    upgradeDb(db)

    logging.Info("data", "Setting Current Version")
    setDbVersion(currentDataVersion, db)

    logging.Info("data", "Migrations finished")
  } else {
    logging.Info("data", "DB up to date - no migration needed")
  }
}


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

func autoMigrateTables() {
  logging.Info("data", "Checking structure: Users")
  DB.AutoMigrate(&user.User{})
  logging.Info("data", "Checking structure: Permissions")
  DB.AutoMigrate(&user.Permission{})
  user.Permission{}.Init(DB)
  logging.Info("data", "Checking structure: Emails")
  DB.AutoMigrate(&user.Email{})
  logging.Info("data", "Checking structure: Addresses")
  DB.AutoMigrate(&user.Address{})
  logging.Info("data", "Checking structure: AuthenticationMethods")
  DB.AutoMigrate(&user.AuthenticationMethod{})
  logging.Info("data", "Checking structure: Sessions")
  DB.AutoMigrate(&session.Session{})

  logging.Info("data", "Checking structure: Entity")
  DB.AutoMigrate(&entity.Entity{})
  logging.Info("data", "Checking structure: EntityPivot")
  DB.AutoMigrate(&entity.EntityPivot{})
  logging.Info("data", "Checking structure: EntityLocationRecord")
  DB.AutoMigrate(&entity.EntityLocationRecord{})
  logging.Info("data", "Checking structure: EntityStatusRecord")
  DB.AutoMigrate(&entity.EntityStatusRecord{})
  logging.Info("data", "Checking structure: EntityLogRecord")
  DB.AutoMigrate(&entity.EntityLogRecord{})
  logging.Info("data", "Checking structure: EntityEvent")
  DB.AutoMigrate(&entity.EntityEvent{})

  logging.Info("data", "Checking structure: Plugin")
  DB.AutoMigrate(&plugin.Plugin{})
  logging.Info("data", "Checking structure: Resource")
  DB.AutoMigrate(&plugin.Resource{})
  logging.Info("data", "Checking structure: Stmdata")
  DB.AutoMigrate(&stmdata.Stmdata{})
}
