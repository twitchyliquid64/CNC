package config



type Config struct {
	Name string							//canonical name to help identify this server
	Database struct{				//information needed to connect to PostGresQL database
		Username string
		Password string
		Name string
		Address string
	}
	Signaller struct {			//Listener address for the old IRC-based signalling backend - commented out ATM.
		SockAddr []string
		MOTD string
	}
	TLS struct {						//Relative file addresses of the .pem files needed for TLS.
		PrivateKey string
		Cert string
	}
	BaseObjects struct {		//List of objects to be autocreated on startup if they dont exist.
		AdminUsers []struct {	//If the username does not exist, it will be created with the given password and admin privs.
			Username string
			Password string
		}
	}
	Web struct{							//Details needed to get the website part working.
		Domain string					//Domain should be in the form example.com
		Listener string				//Address:port (address can be omitted) where the HTTPS listener will bind.
		SimpleHTTPGateway struct {
			Enable bool
			Listener string			//Address:port where the HTTP listener should bind.
		}
	}
	Messenger struct{
		TelegramIntegration struct{
			Token string
			BotUsername string
		}
	}
}
