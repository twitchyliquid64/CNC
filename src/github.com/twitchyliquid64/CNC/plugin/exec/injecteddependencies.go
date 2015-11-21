package exec

//it is done this way to avoid circular dependencies

var LoadBuiltinFunction func(plugin *Plugin)error = nil
var RegisterHookFunction func(plugin *Plugin, hook Hook)error = nil
