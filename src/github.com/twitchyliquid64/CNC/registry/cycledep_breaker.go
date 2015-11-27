package registry


var dispatchMethod func(string, interface{})bool = nil

func SetupDispatchMethod(in func(string, interface{})bool) {
  dispatchMethod = in
}

func DispatchEvent(typ string, data interface{})bool{
  if dispatchMethod != nil {
    return dispatchMethod(typ, data)
  }
  return false
}
