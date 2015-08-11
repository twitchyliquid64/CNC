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
	for _, a := range config.All().Signaller.SockAddr{
		signaller.StartListener(a)
	}

	time.Sleep(time.Second * 9999)
	signaller.Stop()
}
