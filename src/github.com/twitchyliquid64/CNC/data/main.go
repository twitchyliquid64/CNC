package data

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/config"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "database/sql"
)


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

  _, err = gorm.Open("postgres", dbConn)

  if err != nil {
    logging.Error("data", "Error launching DB engine")
    logging.Error("data", "Error: ", err)
  }
}
