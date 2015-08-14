package data

import (
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/config"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "database/sql"
)

var DB gorm.DB

func Initialise() {
  logging.Info("data", "Initialise()")

  dbConn, err := sql.Open("postgres", "postgres://" + config.All().Database.Username +
                                     ":" + config.All().Database.Password +
                                     "@" + config.All().Database.Address +
                                     "/" + config.All().Database.Name +
                                     "?sslmode=require")
	if err != nil {
		logging.Error("data", "Error opening DB connection")
    logging.Error("data", "Error: ", err)
	}

  DB, err = gorm.Open("postgres", dbConn)

  if err != nil {
    logging.Error("data", "Error launching DB engine")
    logging.Error("data", "Error: ", err)
  }

  //make sure that objects in the config BaseObjects are
  //existing, creating them if nessesary.
  for _, usr := range config.All().BaseObjects.AdminUsers{
    var tmp user.User
    logging.Info("data", usr.Username)
    DB.FirstOrCreate(&tmp, &user.User{Username: usr.Username})
    logging.Info("data", tmp)
  }
}
