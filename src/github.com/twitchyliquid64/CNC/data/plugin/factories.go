package plugin


import (
  //"github.com/twitchyliquid64/CNC/logging"
  "github.com/jinzhu/gorm"
)


func GetAllDisabledNoResources(db gorm.DB)[]Plugin{
  var plugins = make([]Plugin, 0)
  db.Where("enabled = ?", false).Find(&plugins)
  return plugins
}

func GetByName(db gorm.DB, name string)Plugin{
  var ret Plugin
  db.Where("name = ?", name).First(&ret)
  return ret
}

func GetAllEnabled(db gorm.DB)[]Plugin{
  var plugins = make([]Plugin, 0)
  db.Where("enabled = ?", true).Find(&plugins)

  for i := 0; i < len(plugins); i++ {
    LoadResources(&(plugins[i]), db)
  }

  return plugins
}

func Get(db gorm.DB, pluginID int)Plugin{
  var plugin Plugin
  db.Find(&plugin, pluginID)
  LoadResources(&plugin, db)
  return plugin
}

func LoadResources(p *Plugin, db gorm.DB){
  db.Model(&p).Related(&p.Resources)
}

func Create(p Plugin, db gorm.DB)error{
  return db.Create(&p).Error
}

func GetResource(db gorm.DB, resourceID int)Resource{
  var resource Resource
  db.Find(&resource, resourceID)
  return resource
}
