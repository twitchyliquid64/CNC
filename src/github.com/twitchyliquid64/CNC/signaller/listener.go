package signaller


import (
	"github.com/twitchyliquid64/CNC/logging"
	"github.com/twitchyliquid64/CNC/config"
	"crypto/tls"
  "net"
	"time"
	"sync"
)

const LISTENER_LOOP_TIMEOUT = 400//miliseconds

//global variables
var gListeners []*Listener = nil
var gListenerWaitGroup sync.WaitGroup

type Listener struct {
  Addr string
  CloseSignal chan bool
  tlsSocket net.Listener
	baseSocket *net.TCPListener
}

//Loops around, accepting new connections on the network socket and pumping them to the server.
//Aborts itself when CloseSignal is sent.
func (l *Listener)Run() {
	gListenerWaitGroup.Add(1)
	defer gListenerWaitGroup.Done()

	for {
		select {
			case <- l.CloseSignal:
				logging.Info("signaller", "Listener shutting down: ", l.Addr)
				l.tlsSocket.Close()
				return
			default:
				l.baseSocket.SetDeadline(time.Now().Add(time.Millisecond * LISTENER_LOOP_TIMEOUT))
				conn, err := l.tlsSocket.Accept()
				if err != nil {
					if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
						continue
					}
					logging.Error("signaller", "Error accepting connection.")
					logging.Error("signaller", err.Error())
					continue
				}
				logging.Info("signaller", "Got new connection from: ", l.Addr, " for ", conn.RemoteAddr())
				gRosellaServer.HandleConnection(conn)
		}
	}
}


func StartListener(addr string)error{
	netAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		logging.Error("signaller", "Invalid address.")
		logging.Error("signaller", err.Error())
		return err
	}

	listener, err := net.ListenTCP("tcp", netAddr)
	if err != nil {
		logging.Error("signaller", "Could not open network listener.")
		logging.Error("signaller", err.Error())
		return err
	}
	listener.SetDeadline(time.Now().Add(time.Millisecond * LISTENER_LOOP_TIMEOUT))

	tlsListener := tls.NewListener(listener, config.TLS())//wrap network socket in TLS
	logging.Info("signaller", "Listening on: ", addr)
	listObj := &Listener{
		Addr: addr,
		CloseSignal: make(chan bool, 0),
		tlsSocket: tlsListener,
		baseSocket: listener,
	}

	go listObj.Run()
	gListeners = append(gListeners, listObj)
	return nil
}

func stopListeners(){
	for _, listener := range gListeners{
		listener.CloseSignal <- true
	}
	gListenerWaitGroup.Wait()
}
