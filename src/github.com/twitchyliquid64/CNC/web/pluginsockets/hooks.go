package pluginsockets

import (
  "github.com/twitchyliquid64/CNC/logging"
  "regexp"
  "sync"
)

var hookMap = map[string]HookRecord{} //maps hook name to the hooks which should be fired when a record is recieved
var mapLock sync.Mutex

type HookRecord struct {
  Regex string
  OnOpen string     //hook name - typically all same (and hook.dispatch() filters and fires to different methods)
  OnMessage string  //hook name
  OnClose string    //hook name
}


func findMatch(url string)HookRecord { //returns the hook name if a match was found else returns ""
  mapLock.Lock()
  defer mapLock.Unlock()

  for _, hookRecord := range hookMap {
    if m, _ := regexp.MatchString(hookRecord.Regex, url); m {
      //logging.Info("web-plugin", "Matched ", regex, " with ", hook)
      return hookRecord
    }
  }
  return HookRecord{}
}

func RemoveHook(hook string){ //shpuld only be called with the onOpen hook
  mapLock.Lock()
  defer mapLock.Unlock()

  if _, exists := hookMap[hook]; !exists{
    logging.Warning("web-pluginsocks", "Cannot remove hook which does not exist: ", hook, " - ", hookMap)
    return
  }
  delete(hookMap, hook)
}


func AddHook(onOpen, onMessage, onClose, regex string)bool{
  mapLock.Lock()
  defer mapLock.Unlock()

  temp := HookRecord{
    Regex: regex,
    OnOpen: onOpen,
    OnMessage: onMessage,
    OnClose: onClose,
  }

  if _, exists := hookMap[onOpen]; exists{
    return false
  }
  hookMap[onOpen] = temp
  return true
}
