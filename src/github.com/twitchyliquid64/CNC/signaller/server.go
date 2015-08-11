package signaller

import (
	"github.com/exec64/rosella"
	"github.com/twitchyliquid64/CNC/logging"
	"github.com/twitchyliquid64/CNC/config"
)

var gRosellaServer *rosella.Server = nil


func Initialise()error{
	logging.Info("signaller", "Initialise()")
	gRosellaServer = rosella.NewServer()
	gRosellaServer.SetName(config.GetServerName())
	gRosellaServer.SetMotd(config.All().Signaller.MOTD)
	go gRosellaServer.Run()
	return nil
}

func Stop(){
	stopListeners()
}
