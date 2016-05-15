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
  trackingSetup()

  dbConn, err := sql.Open("postgres", "postgres://" + config.All().Database.Username +
                                     ":" + config.All().Database.Password +
                                     "@" + config.All().Database.Address +
                                     "/" + config.All().Database.Name +
                                     "?sslmode=require")
	if err != nil {
		logging.Error("data", "Error opening DB connection")
    logging.Error("data", "Error: ", err)
    tracking_notifyFault(err)
	}

  DB, err = gorm.Open("postgres", dbConn)
  DB.LogMode(true)

  if err != nil {
    logging.Error("data", "Error launching DB engine")
    logging.Error("data", "Error: ", err)
  }

  checkStructures(DB)

  //make sure that objects in the config BaseObjects are
  //existing, creating them if nessesary.
  for _, usr := range config.All().BaseObjects.AdminUsers{

    tmp := user.User{}

    DB.Where(&user.User{Username:  usr.Username}).First(&tmp)

    if tmp.Username != usr.Username{ //if the user was not found
      logging.Info("data", "Creating admin user: " + usr.Username)
      DB.Create(&user.User{Username: usr.Username,
                        Permissions: []user.Permission{ user.Permission{Name: user.PERM_ADMIN},},
                        AuthMethods: []user.AuthenticationMethod{ user.AuthenticationMethod{
                            MethodType: user.AUTH_PASSWD,
                            Value: usr.Password,
                          }},
                      })
    }
  }

  logging.Info("data", "Initialisation finished.")
}
