package stmdata

import (
  "github.com/jinzhu/gorm"
)

func Set(pluginID int, key, data string, db gorm.DB)error{

  var d Stmdata
  db.Where(&Stmdata{PluginID:  pluginID, Name: key}).First(&d)
  if d.ID == 0 { //doesnt exist
    d = Stmdata{
      PluginID: pluginID,
      Name: key,
      Data: []byte(data),
    }
    return db.Create(&d).Error
  } else {
    d.Data = []byte(data)
    return db.Save(&d).Error
  }
}


func Get(pluginID int, key string, db gorm.DB)(Stmdata, bool){
  var d Stmdata
  db.Where(&Stmdata{PluginID:  pluginID, Name: key}).First(&d)

  return d, d.ID != 0 //ID != 0, must exist
}
