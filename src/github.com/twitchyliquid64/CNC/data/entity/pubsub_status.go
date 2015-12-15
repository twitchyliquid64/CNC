package entity



import (
  "time"
  "sync"
)

type EntityStatusUpdate struct {
  EntityID uint
  Content string
  Style string
  StyleMeta string
  Icon string
  Created int64
}

var updateSubscribers  = map[chan EntityStatusUpdate]bool{}
var updateSubStructLock sync.Mutex

func SubscribeUpdates(in chan EntityStatusUpdate){ //DO NOT LOG WITHIN THIS MSG - DEADLOCK
  updateSubStructLock.Lock()
  defer updateSubStructLock.Unlock()

  updateSubscribers[in] = true
}

func UnsubscribeUpdates(in chan EntityStatusUpdate){ //DO NOT LOG WITHIN THIS MSG - DEADLOCK
  updateSubStructLock.Lock()
  defer updateSubStructLock.Unlock()

  delete(updateSubscribers, in)
}

func PublishUpdate(eID uint, content, style, styleMeta, icon string){ //DO NOT LOG WITHIN THIS MSG - DEADLOCK
  pkt := EntityStatusUpdate{
    EntityID: eID,
    Content: content,
    Style: style,
    StyleMeta: styleMeta,
    Icon: icon,
    Created: time.Now().Unix(),
  }

  updateSubStructLock.Lock()
  defer updateSubStructLock.Unlock()
  for ch, _ := range updateSubscribers {
    select { //prevents blocking if a channel is full
      case ch <- pkt:
      default:
    }
  }
}
