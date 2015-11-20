package exec

import (
  "github.com/twitchyliquid64/CNC/logging"
  "time"
)

func (p *Plugin)run(){ //should be called in initialise() once everything is setup
  logging.Info("plugin", "plugin.run()")
  for {
    select {
    case invocation, ok := <- p.PendingInvocations:
      if ok && (p.State==STATE_RUNNING){ //channel is open and our state is running
        p.IsCurrentlyInExecution = true
        _, err := p.VM.Call(invocation.MethodName, nil, invocation.Data)
        if err != nil{
          p.Error = err
          p.State = STATE_RUN_ERROR
        }
        p.IsCurrentlyInExecution = false
      }else{
        logging.Info("plugin", "plugin.mainloop() closing down")
        return  //closing down!
      }
    }
  }
}


func (p *Plugin)Stop(){
  logging.Info("plugin", "Now stopping plugin ", p.Name)
  close(p.PendingInvocations)
  time.Sleep(time.Millisecond * 50)
  p.State = STATE_STOPPED
  if p.IsCurrentlyInExecution{
    p.VM.Interrupt <- func(){panic("Interrupt")}
  }
}
