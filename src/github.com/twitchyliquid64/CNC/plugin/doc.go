package plugin


/*
How to create a plugin:
1. Caller needs to use .plugin/exec to create a plugin
2. Caller needs to register that plugin with plugin.RegisterPlugin(plugin)

How to stop a plugin:
1. Caller needs to call DeregisterPlugin(), which will also delete all hooks
2. Caller needs to call .stop() on the plugin

Internally:
  Caller calls exec.BuildPlugin()
  -Sets up structure
  -(goroutine 1)Runs the code the first time
  -(goroutine 1)Sets up a goroutine with a mainloop to execute incoming invocations

  When a hook is created:
  -Method plugin.RegisterHook() called in plugin/exec
  -Adds hook to an array in the plugin structure
  -Method plugin.RegisterHook is called (done this way to avoid circular dependencies)
  -This method adds it to HookByType using the hook type and the plugin name.

  When a hook is invoked: (TODO)
  -Lock is held on the hooksByType structure.
  -All hooks stored in the correct type are iterated and for each hook.
  -For each hook instance of that type, .Dispatch is called. If the hook decides, it can build a
   JSInvocation and add that to the PendingInvocations queue to cause execution.

  When DeregisterPlugin() is invoked:
  -Lock is held on the hooksByType structure.
  -Plugin deleted from pluginByName.
  -removeAllHooksOfPlugin() called, which iterates hooksByType and
   deletes all entries for all hook types with the same plugin name.


*/
