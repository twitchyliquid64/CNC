package main

import (
	"github.com/twitchyliquid64/CNC/signaller"
	"github.com/twitchyliquid64/CNC/logging"
	"github.com/twitchyliquid64/CNC/config"
	"os/signal"
	"syscall"
	"time"
	"os"
)


func run(stopSignal chan bool) {
	logging.Info("main", "Starting server")
	if config.Load("config.json") != nil {
		logging.Fatal("main", "Configuration error")
	}
	logging.Info("main", "Configuration read, now starting '", config.GetServerName(), "'")

	signaller.Initialise()
	for _, a := range config.All().Signaller.SockAddr{
		signaller.StartListener(a)
	}

	for {
		select {
		case <- stopSignal:
			logging.Info("main", "Got stop signal, finalizing now")
			signaller.Stop()
			return
		default:
			time.Sleep(time.Millisecond * 400)
		}
	}
}


func main() {
	processSignal := make(chan os.Signal, 1)
	signal.Notify(processSignal, os.Interrupt, os.Kill, syscall.SIGHUP)
	chanStop := make(chan bool)
	shouldRun := true

	go func(){//goroutine to monitor OS signals
		for{
			s := <- processSignal //wait for signal from OS
			logging.Info("main", "Got OS signal: ", s)
			if s != syscall.SIGHUP{
				shouldRun = false
			}
			chanStop <- true
		}
	}()

	for shouldRun{
		run(chanStop)//will run until signalled to stop from above goroutine
	}
	time.Sleep(time.Millisecond * 100)
}
