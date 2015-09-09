package config



type Config struct {
	Name string							//canonical name to help identify this server
	Database struct{
		Username string
		Password string
		Name string
		Address string
	}
	Signaller struct {
		SockAddr []string
		MOTD string
	}
	TLS struct {
		PrivateKey string
		Cert string
	}
	BaseObjects struct {
		AdminUsers []struct {
			Username string
			Password string
		}
	}
	Web struct{
		Domain string
		Listener string
	}
}
