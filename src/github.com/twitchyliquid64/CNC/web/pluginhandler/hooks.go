package pluginhandler

import (
  "github.com/twitchyliquid64/CNC/logging"
  "regexp"
  "sync"
)

var hookMap = map[string]string{} //maps URL regex's to hook names to dispatch.
var mapLock sync.Mutex



func findMatch(url string)string { //returns the hook name if a match was found else returns ""
  mapLock.Lock()
  defer mapLock.Unlock()

  for hook, regex := range hookMap {
    if m, _ := regexp.MatchString(regex, url); m {
      //logging.Info("web-plugin", "Matched ", regex, " with ", hook)
      return hook
    }
  }
  return ""
}

func RemoveHook(hook string){
  mapLock.Lock()
  defer mapLock.Unlock()

  if _, exists := hookMap[hook]; !exists{
    logging.Warning("web-plugin", "Cannot remove hook which does not exist: ", hook, " - ", hookMap)
    return
  }else {
    deleteHTTPIfExists(hookMap[hook]) //delete the regex from the HTTP allow rule cache
  }

  //logging.Info("web-plugin", "Hook deleted for ", hookMap[hook], " :: ", hook)
  delete(hookMap, hook)
}


func AddHook(hook, regex string, addHTTPException bool)bool{
  mapLock.Lock()
  defer mapLock.Unlock()

  if _, exists := hookMap[hook]; exists{
    return false
  }
  hookMap[hook] = regex
  //logging.Info("web-plugin", "Hook added for ", regex, " :: ", hook)
  if addHTTPException{
    addHTTPAllowRuleForRegex(regex)
  }

  return true
}
