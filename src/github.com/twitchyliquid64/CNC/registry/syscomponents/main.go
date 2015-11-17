package syscomponents

import (
	"github.com/twitchyliquid64/CNC/logging"
  "encoding/json"
  "sync"
)

// Keeps a list of all the internal subsystems in this program, along with their status. Used for tracing the status
// of internals and keeping track of failures / errors etc.
//

var sysComponent = []SysComponent{}
var componentListLock sync.Mutex

type SysComponent interface{
  Name() string     //Returns the name of the component
  IconStr() string  //Returns a string describing the material-icon to display for this component
  IsNominal()bool   //Returns True if everything is operating correctly
  IsDisabled()bool  //Returns True if the component is disabled by configuration
  IsFault()bool     //Returns True on complete failure (not just temporal error)
  Error()string     //Returns a description of a fault
  SetError(error)   //Called to report an error to the system component
}

func Register(component SysComponent) {
  componentListLock.Lock()
  defer componentListLock.Unlock()

  sysComponent = append(sysComponent, component)
}

func SetError(componentName string, e error){
  componentListLock.Lock()
  defer componentListLock.Unlock()

  for _, component := range sysComponent {
    if component.Name() == componentName {
      component.SetError(e)
      break
    }
  }
}

func GetJSON()string{
  var data []map[string]interface{}
  componentListLock.Lock()
  defer componentListLock.Unlock()

  for _, component := range sysComponent {
    temp := map[string]interface{}{}
    temp["Name"] = component.Name()
    temp["Icon"] = component.IconStr()
    if component.IsDisabled() {
      temp["State"] = "Disabled"
    } else if component.IsNominal() {
      temp["State"] = "OK"
    } else if component.IsFault() {
      temp["State"] = "Fault"
    } else {
      temp["State"] = "?"
    }
    if component.IsFault() {
      temp["Error"] = component.Error()
    }

    data = append(data, temp)
  }

  output, err := json.Marshal(data)
  if err != nil{
    logging.Error("registry-syscomponent", "JSON error: ", err)
  }

  return string(output)
}
