package signaller

import (
	"github.com/exec64/rosella"
	"github.com/twitchyliquid64/CNC/logging"
	"github.com/twitchyliquid64/CNC/config"
	"crypto/tls"
)

var gRosellaServer *rosella.Server = nil

func StartOnAddr(addr string)error{
	tlsListener, err := tls.Listen("tcp", addr, config.TLS())
	if err != nil {
		logging.Error("signaller", "Could not open tls listener.")
		logging.Error("signaller", err.Error())
		return err
	}

	logging.Info("signaller", "Listening on: ", addr)

	go func() {//TODO: Add a listen SetDeadline() and a closeChannel to the listener so we can shut it down gracefully.
		for {
			conn, err := tlsListener.Accept()
			if err != nil {
				logging.Error("signaller", "Error accepting connection.")
				logging.Error("signaller", err.Error())
				continue
			}

			gRosellaServer.HandleConnection(conn)
		}
	}()
	return nil
}

func Initialise()error{
	logging.Info("signaller", "Initialise()")
	gRosellaServer = rosella.NewServer()
	gRosellaServer.SetName(config.GetServerName())
	gRosellaServer.SetMotd(config.All().Signaller.MOTD)
	go gRosellaServer.Run()
	return nil
}
