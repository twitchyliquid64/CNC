package httpgateway

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/config"
	"net/http"
  "strings"
)

type BasicHTTPHandler struct{}


// Mapping of all the URLs which the gateway can handle without redirecting
//
//
var allowedURLMapping = map[string]map[string]func(w http.ResponseWriter, req *http.Request){
  "/test": nil,
}

// Redirects a request to it's HTTPS equivalent, automagically calculating the correct host and port based on config.
//
//
func redir(w http.ResponseWriter, req *http.Request) {
  portStr := ""
  if(strings.Contains(config.All().Web.Listener, ":")) { //if a port is specified in the listening address
    portStr = config.All().Web.Listener[strings.Index(config.All().Web.Listener, ":"):] //get it
  }
  if(portStr == ":443") {
    portStr = ""//no point, port 443 is implicit anyway for HTTPS
  }

  newURL := "https://" + config.All().Web.Domain + portStr + req.RequestURI
  logging.Info("web-gateway", "Redirecting ", req.URL, " to ", newURL)
	http.Redirect(w, req, newURL, http.StatusMovedPermanently)
}




// Handling entrypoint for all requests to this basic HTTP gateway.
//
//
func (f *BasicHTTPHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  handler, ok := allowedURLMapping[req.URL.String()]

  if !ok{ //no handler for this URL, redirect request to web package (HTTPS address).
    redir(w,req)
  } else {
    logging.Info("web-gateway", "handling ", req.URL.String())
    if handler != nil{
      //TODO: Call handler method
    }
  }
}





// Called during CNC/web initialisation.
//
//
func Init() {
  if(config.All().Web.SimpleHTTPGateway.Enable) {
    
    // sets up a gateway on the address specificed in config, with all req's handled by ServeHTTP above
    go func(){
      err := http.ListenAndServe(config.All().Web.SimpleHTTPGateway.Listener, &BasicHTTPHandler{})
      tracking_notifyFault(err)
    }()

    logging.Info("web-gateway", "Initialised server on ", config.All().Web.SimpleHTTPGateway.Listener)
    trackingSetup(false)//enabled
  } else {
    logging.Warning("web-gateway", "HTTP Gateway disabled.")
    trackingSetup(true) //disabled
  }
}
