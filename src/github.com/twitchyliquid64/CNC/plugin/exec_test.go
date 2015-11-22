package plugin

import (
  "github.com/twitchyliquid64/CNC/plugin/builtin"
  "github.com/twitchyliquid64/CNC/plugin/exec"
  //"github.com/twitchyliquid64/CNC/logging"
  "testing"
  "time"
)

func TestBasicExec(t *testing.T) {
  resetAndInit()
  builtin.TestEndpointGood_called = false

  code := `
  var i = 0;
  i = 3;
  i = i * 4;
  log(i);
  testendpoint_good();
  `

  p1 := exec.BuildPlugin("Test1", code)
  RegisterPlugin(p1)
  time.Sleep(time.Millisecond * 100)
  if p1.State != exec.STATE_RUNNING {
    t.Error("Incorrect state")
  }

  DeregisterPlugin(p1)
  p1.Stop()

  if !builtin.TestEndpointGood_called {
    t.Error("Did not run builtin successfully")
  }
}

func TestSyntaxError(t *testing.T) {
  resetAndInit()

  code := `var i = 4fdfgfgfgfd ' LOLOLOL KEK`

  p1 := exec.BuildPlugin("Test1", code)
  RegisterPlugin(p1)
  time.Sleep(time.Millisecond * 100)
  if p1.State != exec.STATE_CODE_ERROR {
    t.Error("Incorrect state")
  }

  DeregisterPlugin(p1)
  p1.Stop()
}


func TestInterrupt(t *testing.T) {
  resetAndInit()

  code := `
  var i = 0;
  while (true){
    i = i + 1;
  }`

  p1 := exec.BuildPlugin("Test1", code)
  RegisterPlugin(p1)
  time.Sleep(time.Millisecond * 100)
  DeregisterPlugin(p1)
  p1.Stop()
  time.Sleep(time.Millisecond * 100)
  if p1.IsCurrentlyInExecution {
    t.Error("Interrupt failed")
  }
  if p1.State != exec.STATE_STOPPED {
    t.Error("Wrong state:", p1.State)
  }
}

func TestDispatch(t *testing.T) {
  resetAndInit()
  builtin.TestEndpointGood_called = false

  code := `
  function testDispatchHandler() {
  log('Kek, it actually got called')
    testendpoint_good();
  }
  onTestDispatchTriggered("testDispatchHandler");
  `

  p1 := exec.BuildPlugin("Test1", code)
  RegisterPlugin(p1)
  time.Sleep(time.Millisecond * 100)
  if p1.Error != nil{
    t.Error("code error:", p1.Error)
  }
  if p1.State != exec.STATE_RUNNING{
    t.Error("Wrong State")
  }
  Dispatch("dispatchtest", nil)
  time.Sleep(time.Millisecond * 100)
  if !builtin.TestEndpointGood_called {
    t.Error("Did not run builtin successfully")
  }
}
