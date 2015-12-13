package pluginhandler

import (
  "github.com/twitchyliquid64/CNC/logging"
  "regexp"
  "sync"
)

var allowedRegexs = map[string]bool{} //contains regex's which are allowed to be served via HTTP
var arLock sync.Mutex


func CheckHTTPAllowed(url string)bool{
  arLock.Lock()
  defer arLock.Unlock()

  for regex, _ := range allowedRegexs {
    if m, _ := regexp.MatchString(regex, url); m {
      return true
    }
  }
  return false
}

func deleteHTTPIfExists(regex string){
  arLock.Lock()
  defer arLock.Unlock()

  if _, exists := allowedRegexs[regex]; !exists{
    return
  }

  logging.Info("web-plugin", "Removing exception to HTTP -> HTTPS re-route for ", regex)
  delete(allowedRegexs, regex)
}

func addHTTPAllowRuleForRegex(regex string){
  logging.Info("web-plugin", "Adding exception to HTTP -> HTTPS re-route for ", regex)
  allowedRegexs[regex] = true
}
