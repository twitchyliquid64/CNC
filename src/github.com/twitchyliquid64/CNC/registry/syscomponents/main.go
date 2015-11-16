package syscomponents

import (
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
}

func Register(component SysComponent) {
  componentListLock.Lock()
  defer componentListLock.Unlock()

  sysComponent = append(sysComponent, component)
}
