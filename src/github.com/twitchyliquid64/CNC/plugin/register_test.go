package plugin

import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/logging"
  "testing"
)

func resetAndInit() {
  pluginByName = nil
  hooksByType = nil
  Initialise()
}

func TestRegistration(t *testing.T) {
  resetAndInit()
  if pluginByName==nil || hooksByType==nil || len(pluginByName)>0 || len(hooksByType)>0{
    t.Error("Initialisation failed")
  }

  testp1 := exec.Plugin{Name: "test1"} //mocked without internals as we are just testing registration

  RegisterPlugin(&testp1)
  if len(pluginByName) != 1{
    t.Error("Structure save error 1")
  }
  if pluginByName["test1"].Name != "test1"{
    t.Error("Structure save error 2")
  }
}

func TestDeregistration(t *testing.T) {//assumes TestRegistration passes
  resetAndInit()
  testp1 := exec.Plugin{Name: "test1"} //mocked without internals as we are just testing registration

  RegisterPlugin(&testp1)
  DeregisterPlugin(&exec.Plugin{Name: "test1"})
  if len(pluginByName) != 0{
    t.Error("Structure delete error 1")
  }
}

func TestMuliPluginRegistration(t *testing.T) {//Assumes TestDeregistration passes
  resetAndInit()
  testp1 := exec.Plugin{Name: "test1"} //mocked without internals as we are just testing registration
  testp2 := exec.Plugin{Name: "test2"} //mocked without internals as we are just testing registration
  testp3 := exec.Plugin{Name: "test3"} //mocked without internals as we are just testing registration

  RegisterPlugin(&testp1)
  RegisterPlugin(&testp2)
  if len(pluginByName) != 2{
    t.Error("Structure save error 3")
  }
  if _, ok := pluginByName["test3"]; ok {
    t.Error("test3 should not be present")
  }

  DeregisterPlugin(&testp1)
  if len(pluginByName) != 1{
    t.Error("Structure delete error 2")
  }

  RegisterPlugin(&testp3)
  if len(pluginByName) != 2{
    t.Error("Structure save error 4")
  }
  if _, ok := pluginByName["test1"]; ok {
    t.Error("test1 should not be present")
  }
  if _, ok := pluginByName["test2"]; !ok {
    t.Error("test2 should be present")
  }
  if _, ok := pluginByName["test3"]; !ok {
    t.Error("test3 should be present")
  }

  RegisterPlugin(&testp1)
  if len(pluginByName) != 3{
    t.Error("Structure save error 5")
  }
}




type MockHook struct {
  TestName string
}
func (h *MockHook)Destroy(){
  logging.Info("mockHook", "Destroy() called")
}
func (h *MockHook)Name()string{
  return h.TestName
}
func (h *MockHook)Dispatch(data interface{}){}


func TestHookRegistration(t *testing.T) {
  resetAndInit()

  if len(hooksByType) != 0{
    t.Error("Expected hooksByType to be 0 on start")
  }

  testp1 := exec.Plugin{Name: "Test1"}
  testm1 := MockHook{TestName: "Testhook"}
  RegisterPlugin(&testp1)
  testp1.RegisterHook(&testm1)

  //we should now have a single hook with key "Testhook1"
  if len(hooksByType) != 1{
    t.Error("Expected hooksByType to be 1 after a registration")
    t.FailNow()
  }
  if _, ok := hooksByType["Testhook"]; !ok {
    t.Error("Testhook1 should be present, instead:", hooksByType)
    t.FailNow()
  }

  //test adding another hook of the same type with a different plugin
  testp2 := exec.Plugin{Name: "Test2"}
  testm2 := MockHook{TestName: "Testhook"}
  RegisterPlugin(&testp2)
  testp2.RegisterHook(&testm2)
  pluginsWithHook := hooksByType["Testhook"]
  if len(pluginsWithHook) != 2 {
    t.Error("Expected two plugins registered with the same hook type")
  }
  if _, ok := pluginsWithHook["Test1"]; !ok{
    t.Error("Expected Test1 to be present")
  }
  if _, ok := pluginsWithHook["Test2"]; !ok{
    t.Error("Expected Test2 to be present")
  }

}

func TestHookDeregistration(t *testing.T) {
  resetAndInit()

  //setup first plugin
  testp1 := exec.Plugin{Name: "Test1"}
  testm1 := MockHook{TestName: "Testhook"}
  RegisterPlugin(&testp1)
  testp1.RegisterHook(&testm1)

  //setup second plugin
  testp2 := exec.Plugin{Name: "Test2"}
  testm2 := MockHook{TestName: "Testhook"}
  testm3 := MockHook{TestName: "Testhook2"}
  RegisterPlugin(&testp2)
  testp2.RegisterHook(&testm2)
  testp2.RegisterHook(&testm3)

  //unregister first plugin and see if the hooks got removed
  DeregisterPlugin(&testp1)
  if len(hooksByType) != 2 || len(pluginByName) != 1 {
    t.Error("Test assertion failure")
    t.FailNow()
  }
  if _, ok := hooksByType["Testhook"]["Test1"]; ok{
    t.Error("Expected Test1 to not have hook Testhook")
    t.FailNow()
  }
  if _, ok := hooksByType["Testhook"]["Test2"]; !ok{
    t.Error("Expected Test2 to have hook Testhook")
    t.FailNow()
  }

  //unregister second and see that there are now no hooks on either.
  //the two hooksByType entries should still be there though, just the
  //value should be empty maps.
  DeregisterPlugin(&testp2)
  if len(hooksByType) != 2 || len(pluginByName) != 0 {
    t.Error("Test assertion failure")
    t.FailNow()
  }
  if len(hooksByType["Testhook"]) != 0 || len(hooksByType["Testhook2"]) != 0 {
    t.Error("Delete failed")
  }

}
