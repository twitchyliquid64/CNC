package entity



import (
  "time"
  "sync"
)

type UpdateType string
const (
  Updatetype_Status UpdateType = "status"
  Updatetype_Location UpdateType = "location"
  Updatetype_Log UpdateType = "log"
)


type EntityUpdate struct {
  EntityID uint
  Type UpdateType

  //used when updateType == "status"
  Content string
  Style string
  StyleMeta string
  Icon string

  //used when updateType == "location"
  Latitude float64
  Longitude float64
  SpeedKph float64
  Course      int
  SatNum      int
  Accuracy    int

  //time at which the update occurred in unix seconds
  Created int64
}

var updateSubscribers  = map[chan EntityUpdate]bool{}
var updateSubStructLock sync.Mutex

func SubscribeUpdates(in chan EntityUpdate){
  updateSubStructLock.Lock()
  defer updateSubStructLock.Unlock()

  updateSubscribers[in] = true
}

func UnsubscribeUpdates(in chan EntityUpdate){
  updateSubStructLock.Lock()
  defer updateSubStructLock.Unlock()

  delete(updateSubscribers, in)
}

func PublishStatusUpdate(eID uint, content, style, styleMeta, icon string)EntityUpdate{
  pkt := EntityUpdate{
    EntityID: eID,
    Type: Updatetype_Status,
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

  return pkt
}

func PublishLocationUpdate(eID uint, lat, lon, speed float64, acc, course, sat int){
  pkt := EntityUpdate{
    EntityID: eID,
    Type: Updatetype_Location,

    Latitude: lat,
    Longitude: lon,
    SpeedKph: speed,
    Accuracy: acc,
    Course: course,
    SatNum: sat,

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
