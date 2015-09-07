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

  checkStructures()

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

//called during initialisation. Should make sure the schema is intact and up to date.
func checkStructures() {
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
}
