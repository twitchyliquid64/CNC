package plugin


import (
  //"github.com/twitchyliquid64/CNC/logging"
  "github.com/jinzhu/gorm"
)


func GetAll(db gorm.DB)[]Plugin{
  var plugins = make([]Plugin, 0)
  db.Find(&plugins)

  for i := 0; i < len(plugins); i++ {
    LoadResources(&(plugins[i]), db, true)
  }

  return plugins
}

func Get(db gorm.DB, pluginID int, trimResourceData bool)Plugin{
  var plugin Plugin
  db.Find(&plugin, pluginID)
  LoadResources(&plugin, db, trimResourceData)
  return plugin
}

func LoadResources(p *Plugin, db gorm.DB, trimResourceData bool){
  db.Model(&p).Related(&p.Resources)
  if trimResourceData {
    for i := 0; i < len(p.Resources); i++ {
      p.Resources[i].Data = ""
    }
  }
}

func Create(p Plugin, db gorm.DB)error{
  return db.Create(&p).Error
}
