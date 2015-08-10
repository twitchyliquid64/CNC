package main

import (
	"github.com/twitchyliquid64/CNC/signaller"
	"github.com/twitchyliquid64/CNC/logging"
	"github.com/twitchyliquid64/CNC/config"
	"time"
)


func main() {
	if config.Load("config.json") != nil {
		logging.Fatal("main", "Configuration error")
	}
	logging.Info("main", "Started server '", config.GetServerName(), "'")
	
	signaller.Initialise()
	signaller.StartOnAddr(config.All().Signaller.SockAddr)
	
	time.Sleep(time.Second * 99999)
}
