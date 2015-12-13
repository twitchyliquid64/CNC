<!DOCTYPE html>
<html lang="en">

  <head>
      <title>CNC Plugin API Reference</title>
       <link href="/static/css/v-accordion.min.css" rel="stylesheet">
      {!{template "headcontent"}!}
  </head>

  <body layout="column"  ng-app="apiApp" ng-cloak>
      <md-toolbar layout="row" flex="10">
        <h1 flex><md-icon md-font-library="material-icons" style="font-size: 250%;">code</md-icon> CNC Plugin API Reference</h1>
      </md-toolbar>

    <div layout="row" flex>
      <div flex layout="column" tabIndex="-1" role="main" class="md-whiteframe-z2">

        <md-content md-padding md-margin flex layout-fill style="padding-left: 12px; padding-right: 12px;">
          <p>This document outlines the available features within plugins.</p>

          <v-accordion class="vAccordion--default" multiple>

            <!-- LOGGING FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">list</md-icon>
                Logging
              </v-pane-header>

              <v-pane-content>
                <p>These functions allow you to write to the system log and to stdout of the server process.</p>

                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      log(
                      <i class="amber">message</i> <sup style="color: #444444;">str</sup> )
                    </v-pane-header>
                    <v-pane-content>
                      Posts an INFO message to the system log (available in the summary tab). The message is tagged with
                      the name of the plugin prefixed with 'PLUGIN-'.
                      <p>EG: </p>
                      <pre>
                      log("Your shit."); //Output: [I] [PLUGIN-pluginnamehere] Your Shit.
                      </pre>
                    </v-pane-content>
                  </v-pane>

                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      console.log(
                      <i class="amber">message</i> <sup style="color: #444444;">str</sup> ... )
                    </v-pane-header>
                    <v-pane-content>
                      Writes data out to standard output. Should not use unless your an idiot or you dont want to dirty the system log.
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>







            <!-- GMAIL FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">mail</md-icon>
                Gmail
              </v-pane-header>

              <v-pane-content>
                <p>These functions allow you to send emails from a Gmail address, given a username and password. If
                  Two-factor authentication is used on the account, you will need to generate an 'app specific password' (google it).</p>

                <p>There are two steps to sending an email. First, you should call <span class="green">setup()</span>
                to initialise the system with the correct credentials. Next, you should immediately call <span class="green">sendMessage()</span>
                with your subject, a string list of recipients, and the contents (optionally html) of your email.</p>

                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      gmail.setup(
                      <i class="amber">username</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">password</i> <sup style="color: #444444;">str</sup> )
                    </v-pane-header>
                    <v-pane-content>
                      Sets up the system with the correct credentials for when you call an email. NOTE call this immediately before you call
                      sendMessage().
                      <p>EG: </p>
                      <pre>
                      gmail.setup('jacob@gmail.com', 'pass');
                      </pre>
                    </v-pane-content>
                  </v-pane>

                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      gmail.sendMessage(
                      <i class="amber">subject</i> <sup style="color: #444444;">str</sup> ,
                      <i class="amber">recipients</i> <sup style="color: #444444;">list of strings</sup> ,
                      <i class="amber">content</i> <sup style="color: #444444;">str</sup> )
                    </v-pane-header>
                    <v-pane-content>
                      Sets up the system with the correct credentials for when you call an email. NOTE call this immediately before you call
                      sendMessage().
                      <p>EG: </p>
                      <pre>
                      gmail.sendMessage('Test Subject', ['barry@gmail.com'], 'KEK BRAAAH');
                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>




            <!-- CRON FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">build</md-icon>
                CRON
              </v-pane-header>

              <v-pane-content>
                <p>This feature allows you to schedule a method to be run at a certain time, or ever x seconds/minutes/hours</p>

                <p>The cron format used when writing <i>schedule descriptor</i> is described in <a href="https://godoc.org/github.com/robfig/cron">https://godoc.org/github.com/robfig/cron</a></p>
                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      cron.schedule(
                      <i class="amber">schedule descriptor</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">method name</i> <sup style="color: #444444;">str</sup> )
                    </v-pane-header>
                    <v-pane-content>
                      Sets up a method to be called at the time specified by <i>schedule descriptor</i>.
                      <p>EG: </p>
                      <pre>
                        function periodicCron(dummy) {
                            log("CRON trigger activated")
                        }

                        cron.schedule("@every 15m", "periodicCron");
                      </pre>
                    </v-pane-content>
                  </v-pane>

                  <v-pane>
                    <v-pane-header style="color: #000077;">
                      <md-icon md-font-library="material-icons">call_missed </md-icon>
                      function methodname (
                      <i class="amber">param1</i> <sup style="color: #444444;">obj</sup> ->
                      (
                        <i class="amber">CronID</i> <sup style="color: #444444;">str</sup>
                      )
                      ) { }
                    </v-pane-header>
                    <v-pane-content>
                      Describes the callback of the method whoes name is passed in cron.schedule().

                      <p>EG: </p>
                      <pre>
                        function periodicCron(dummy) {
                            log("CRON trigger activated")
                        }

                        cron.schedule("@every 15m", "periodicCron");
                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>






            <!-- TWILIO FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">message</md-icon>
                Twilio
              </v-pane-header>

              <v-pane-content>
                <p>This feature allows you to send SMS'es. Be advised this costs the owner (about 6c per SMS).</p>

                <p>The owner will have one or more numbers setup from which can be SMS'ed. One needs to be known to send SMS'es.</p>
                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      twilio.sendSMS(
                      <i class="amber">from number</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">to number</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">message</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      To and from addresses must be in the format: +COUNTRYCODENUMBER.
                      <p>EG: </p>
                      <pre>
                        var fromNumber = "+19284332644";
                        var toAddress = "+61342320983";

                        twilio.sendSMS(fromNumber, toAddress, "YOLO!!! :D");

                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>






            <!-- PLUGIN FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">whatshot</md-icon>
                Self Aware
              </v-pane-header>

              <v-pane-content>
                <p>This feature allows you to access information / control parameters directly associated with the life of the plugin.</p>

                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      plugin.getResources()
                    </v-pane-header>
                    <v-pane-content>
                      Returns a list of the resources associated with the plugin.
                      <p>EG: </p>
                      <pre>
                        var resources = plugin.getResources();
                        for (var i = 0; i < resources.length ; i++) {
                          log(resources[i].name);
                          log(resources[i].data);
                          log(resources[i].isJs);
                          log(resources[i].isTemplate);
                        }
                      </pre>
                    </v-pane-content>
                  </v-pane>

                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      plugin.getResource(
                      <i class="amber">resource name</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      Returns an object representing a single resource.
                      <p>EG: </p>
                      <pre>
                        //assumes a resource called 'main' exists.
                        var resource = plugin.getResource('main');

                        if (resource == undefined) {
                          log("resource does not exist!");
                        } else {
                          log(resource.name);
                          log(resource.data);
                          log(resource.isJs);
                          log(resource.isTemplate);
                        }
                      </pre>
                    </v-pane-content>
                  </v-pane>




                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      plugin.disable()
                    </v-pane-header>
                    <v-pane-content>
                      Halts current execution, releases all resources and kills/disables the plugin.
                      <p>EG: </p>
                      <pre>
                        console.log("starting");
                        plugin.disable();
                        console.log("after (should not be run)");
                      </pre>
                    </v-pane-content>
                  </v-pane>




                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      plugin.getIcon()
                    </v-pane-header>
                    <v-pane-content>
                      Returns the string resource describing the icon associated with the plugin. (EG: 'wifi')
                      <p>EG: </p>
                      <pre>
                        log("Icon: " + plugin.getIcon());
                      </pre>
                    </v-pane-content>
                  </v-pane>

                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      plugin.setIcon(
                        <i class="amber">icon resource name</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      Sets the icon to the new material-icons resource. See <a href="https://www.google.com/design/icons/">here</a> for a list of icons.
                      <p>EG: </p>
                      <pre>
                        log("Icon: " + plugin.getIcon());
                        plugin.setIcon("error");
                        log("Icon: " + plugin.getIcon());
                      </pre>
                    </v-pane-content>
                  </v-pane>


                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      plugin.ready(
                      <i class="amber">method name</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      Queues a method to be called by the system, once all resources are loaded and pending events handled.
                      <p>EG: </p>
                      <pre>
                        function kek(){
                          log('hi');
                        }

                        plugin.ready('kek');
                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>


            <!-- DATA FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">storage</md-icon>
                Data
              </v-pane-header>

              <v-pane-content>
                <p>This feature allows you to persistantly store information (persistant across plugin restarts) in a key value store. Keys and values are both strings.</p>

                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      data.get(
                      <i class="amber">key</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      Returns undefined if the entry does not exist, otherwise returns an object representing the data entry.
                      <p>EG: </p>
                      <pre>
                        var a = data.get('lolz');
                        if (a === undefined) {
                          log('data does not exist');
                        } else {
                          log(a.Content); //the contents
                          log(a.Name);    //the key
                        }
                      </pre>
                    </v-pane-content>
                  </v-pane>

                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      data.set(
                      <i class="amber">key</i> <sup style="color: #444444;">str</sup>,
                      <i class="amber">data</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      Sets a value in the keystore. Make sure that <i>data</i> is a string.
                      <p>EG: </p>
                      <pre>
                        var a = data.set('lolz', 'kek');
                        if (a === true) {
                          log('Save successful')
                        } else {
                          log('Error saving: ' + a.error);
                        }
                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>





            <!-- WEB FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">cloud</md-icon>
                Web Handlers
              </v-pane-header>

              <v-pane-content>
                <p>This feature allows you to specify URLs, prefixed with /p/, which when hit by a HTTP request can be serviced by a method which you specify.</p>

                <p>You need to 'register' the name of the method which will handle requests, along with a regex that will match paths.</p>
                <p>Inside the method, you control the response and access request parameters by accessing properties and calling methods on the parameter passed.</p>
                <p>Please note that if your plugin's dispatch queue is full, requests will be silently dropped and a 'Plugin Timeout' error will be reported to the client.</p>
                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      web.handle(
                      <i class="amber">url path (regex)</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">method</i> <sup style="color: #444444;">callback</sup>
                      ,
                      <i class="amber">allow via HTTP</i> <sup style="color: #444444;">boolean</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      Registers a method to handle web requests to /p/ which match the specified regex. Returns True if the registration was successful.
                      <p>EG: </p>
                      <pre>
                        function handleHi(web) {
                            web.done();
                        }

                        var didBind = web.handle("/p/hi", "handleHi", false);//a request on HTTP would redirect to HTTPS gateway.
                        log("Bound to /p/hi: " + didBind);
                        web.handle("/p/insecure/hi", "handleHi", true); //a request on HTTP would be responded without redirection.
                      </pre>
                    </v-pane-content>
                  </v-pane>


                  <v-pane>
                    <v-pane-header style="color: #000077;">
                      <md-icon md-font-library="material-icons">call_missed </md-icon>
                      function methodname (
                      <i class="amber">param1</i> <sup style="color: #444444;">obj</sup> ->
                      (
                        <i class="amber">url</i> <sup style="color: #444444;">str</sup> ,
                        <i class="amber">data</i> <sup style="color: #444444;">str</sup> ,
                        <i class="green">done()</i> <sup style="color: #444444;">method</sup> ,
                        <i class="green">write( <i class="amber">data</i> <sup style="color: #444444;">str</sup> )</i><sup style="color: #444444;">method</sup>
                        <i class="green">isLoggedIn()</i> <sup style="color: #444444;">method</sup> ,
                        <i class="green">user()</i> <sup style="color: #444444;">method</sup> ,
                        <i class="green">session()</i> <sup style="color: #444444;">method</sup> ,

                      ) { }
                    </v-pane-header>
                    <v-pane-content>
                      Describes the callback of the method whoes name is passed in web.handle().
                      <p>Some notes:</p>
                      <ul>
                        <li>Use <i class="green">param1.write(data)</i> to write content to the browser.</li>
                        <li>Use <i class="green">param1.done()</i> to finish the request. Do NOT write after calling done.</li>
                        <li><i class="green">param1.url</i> is the full request URL.</li>
                        <li><i class="green">param1.data</i> contains the POST data of the request, if any.</li>
                        <li><i class="green">param1.parameter(paramname<sup style="color: #444444;">str</sup>)</i> returns the value of the GET/POST parameter, if any.</li>
                        <li><i class="green">param1.isLoggedIn()</i> returns true if the browser that issued the request has a valid cookie, which corresponds with a web session for a user on CNC.</li>
                        <li><i class="green">param1.user()</i> returns the user object of the logged in user, if any. Format as described in data/user/model.go.</li>
                        <li><i class="green">param1.session()</i> returns the session object of the logged in user, if any. Format as described in data/session/model.go.</li>
                      </ul>

                      <p>EG: </p>
                      <pre>
                        function handleHi(web) {
                            log(web.url);
                            log(web.data);
                            web.write(web.parameter("name") + "&lt;br&gt;&lt;br&gt;");

                            if (web.isLoggedIn()) {
                                web.write("Hello " + web.user().Firstname + " " + web.user().Lastname + "(" + web.user().Username + ")!");
                                web.write("&lt;br&gt;Your session was created at " + JSON.stringify(web.session().CreatedAt));
                            } else {
                                web.write("You are not logged in");
                            }
                            web.done();
                        }

                        var didBind = web.handle("/p/hi", "handleHi");
                        log("Bound to /p/hi: " + didBind);
                      </pre>
                    </v-pane-content>
                  </v-pane>

                </v-accordion>

              </v-pane-content>
            </v-pane>



            <!-- WEBSOCKET FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">wifi</md-icon>
                Websocket Handlers
              </v-pane-header>

              <v-pane-content>
                <p>This feature allows you to specify URLs, prefixed with /ws/p/, which when hit by a Websocket request can be serviced by methods which you specify.</p>

                <p>You need to 'register' the name of the methods which will handle requests, along with a regex that will match paths for the websocket requests you recieve.</p>
                <p>You will need to create three methods - one which fires when a websocket is recieved, one which fires when a message is recieved on an open websocket your plugin owns, and one which will fire when an open websocket you own is closed (triggered either by remote closure, or by you calling handle.close()).</p>
                <p>When fired, your methods will be passed a single parameter which contains information about the websocket and methods you can hit to affect it.</p>
                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      websockets.register(
                      <i class="amber">url path (regex)</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">on-socket-opened method name</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">on-socket-messaged method name</i> <sup style="color: #444444;">str</sup>
                      ,
                      <i class="amber">on-socket-closed method name</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      Registers a method to handle websocket requests to /ws/p/ which match the specified regex. Returns True if the registration was successful.
                      <p>EG: </p>
                      <pre>
                        function onOpen(handle){
                            handle.close();
                        }

                        function onMessage(handle){
                        }

                        function onClose(handle){
                        }

                        log("Did bind WS: " + websockets.register("/ws/p/test", "onOpen", "onClose", "onMessage"));
                      </pre>
                    </v-pane-content>
                  </v-pane>


                  <v-pane>
                    <v-pane-header style="color: #000077;">
                      <md-icon md-font-library="material-icons">call_missed </md-icon>
                      function methodname (
                      <i class="amber">param1</i> <sup style="color: #444444;">obj</sup> ->
                      (
                        <i class="amber">url</i> <sup style="color: #444444;">str</sup> ,
                        <i class="amber">data</i> <sup style="color: #444444;">str</sup> ,
                        <i class="amber">addr</i> <sup style="color: #444444;">str</sup> ,
                        <i class="amber">id</i> <sup style="color: #444444;">int</sup> ,
                        <i class="green">write( <i class="amber">data</i> <sup style="color: #444444;">str</sup> )</i><sup style="color: #444444;">method</sup>
                        <i class="green">close( )</i><sup style="color: #444444;">method</sup>
                        <i class="green">isLoggedIn()</i> <sup style="color: #444444;">method</sup> ,
                        <i class="green">user()</i> <sup style="color: #444444;">method</sup> ,
                        <i class="green">session()</i> <sup style="color: #444444;">method</sup>

                      ) { }
                    </v-pane-header>
                    <v-pane-content>
                      Describes the callback of the methods whoes names are passed in websockets.register().
                      <p>Some notes:</p>
                      <ul>
                        <li>Use <i class="green">param1.write(data)</i> to write content to the socket. Note that it takes a string - you must serialize all objects.</li>
                        <li>Use <i class="green">param1.close()</i> Closes the websocket.</li>
                        <li><i class="green">param1.url</i> is the full request URL. EG: '/ws/p/blah'</li>
                        <li><i class="green">param1.data</i> contains the content of the message. Only present in the parameter for the onMessage callback.</li>
                        <li><i class="green">param1.addr</i> returns the address of the remote endpoint.</li>
                        <li><i class="green">param1.id</i> returns an integer that uniquely identifies the websocket from all other websockets since the time CNC was started.</li>
                        <li><i class="green">param1.isLoggedIn()</i> returns true if the browser that issued the request has a valid cookie, which corresponds with a web session for a user on CNC.</li>
                        <li><i class="green">param1.user()</i> returns the user object of the logged in user, if any. Format as described in data/user/model.go.</li>
                        <li><i class="green">param1.session()</i> returns the session object of the logged in user, if any. Format as described in data/session/model.go.</li>
                      </ul>

                      <p>EG: </p>
                      <pre>
                        function onOpen(handle){
                            log("WS opened: " + JSON.stringify(handle));
                            handle.write("LOLZ");
                        }

                        function onMessage(handle){
                            var d = handle.data;
                            log("Got WS message: " + d + "::" + handle.id + "::" + handle.addr);
                            handle.write("GONNA KILL TEH CONN LOLZ");
                            handle.close()
                        }

                        function onClose(handle){
                            log("WS closed: " + JSON.stringify(handle));
                        }

                        log("Did bind WS: " + websockets.register("/ws/p/chat", "onOpen", "onClose", "onMessage"));
                      </pre>
                    </v-pane-content>
                  </v-pane>

                </v-accordion>

              </v-pane-content>
            </v-pane>



            <!-- Web Request FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">http</md-icon>
                Web Requests
              </v-pane-header>

              <v-pane-content>
                <p>This feature allows you to make HTTP and HTTPS requests.

                <p>At the moment, there are a limited amount of parameters that can be configured.</p>
                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      http.get(
                      <i class="amber">fully qualified URL</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      <p>After making the request, the function returns an object that represents its response. This object has
                      the following properties:</p>

                      <ul>
                        <li><i class="green">Code</i> <sup style="color: #444444;">int</sup> - The HTTP response code (200, 404 etc)</li>
                        <li><i class="green">CodeStr</i> <sup style="color: #444444;">str</sup> - The HTTP response status expressed as a human readable string</li>
                        <li><i class="green">Data</i> <sup style="color: #444444;">str</sup> - The HTTP response body</li>
                        <li><i class="green">Address</i> <sup style="color: #444444;">str</sup> - The address in the request</li>
                      </ul>

                      <p>EG: </p>
                      <pre>
                        var data = http.get('http://ciphersink.net/');
                        console.log(JSON.stringify(data));
                      </pre>
                    </v-pane-content>
                  </v-pane>

                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      http.post(
                      <i class="amber">fully qualified URL</i> <sup style="color: #444444;">str</sup>,
                      <i class="amber">data</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      <p>Does a HTTP POST to the specified URL, with the second parameter used as the request body.</p>
                      <p>After making the request, the function returns an object that represents its response. This object has
                      the following properties:</p>

                      <ul>
                        <li><i class="green">Code</i> <sup style="color: #444444;">int</sup> - The HTTP response code (200, 404 etc)</li>
                        <li><i class="green">CodeStr</i> <sup style="color: #444444;">str</sup> - The HTTP response status expressed as a human readable string</li>
                        <li><i class="green">Data</i> <sup style="color: #444444;">str</sup> - The HTTP response body</li>
                        <li><i class="green">Address</i> <sup style="color: #444444;">str</sup> - The address in the request</li>
                      </ul>

                      <p>EG: </p>
                      <pre>
                        var res = http.post('http://posttestserver.com/post.php', 'text/html', 'LOLCAKES');
                        console.log(JSON.stringify(res));
                      </pre>
                    </v-pane-content>
                  </v-pane>


                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      http.postValues(
                      <i class="amber">fully qualified URL</i> <sup style="color: #444444;">str</sup>,
                      <i class="amber">values</i> <sup style="color: #444444;">obj</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      <p>Does a HTTP POST to the specified URL, Using a formencoded body specifying the given key-value pairs in <i>values</i>.</p>
                      <p>After making the request, the function returns an object that represents its response. This object has
                      the following properties:</p>

                      <ul>
                        <li><i class="green">Code</i> <sup style="color: #444444;">int</sup> - The HTTP response code (200, 404 etc)</li>
                        <li><i class="green">CodeStr</i> <sup style="color: #444444;">str</sup> - The HTTP response status expressed as a human readable string</li>
                        <li><i class="green">Data</i> <sup style="color: #444444;">str</sup> - The HTTP response body</li>
                        <li><i class="green">Address</i> <sup style="color: #444444;">str</sup> - The address in the request</li>
                      </ul>

                      <p>EG: </p>
                      <pre>
                        var resFinal = http.postValues('http://posttestserver.com/post.php', {yolo: 'swag', 'bruh': 'brah'});
                        console.log(JSON.stringify(resFinal));
                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>









            <!-- TELEGRAM FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">send</md-icon>
                Telegram
              </v-pane-header>

              <v-pane-content>
                <p>This feature allows you to integrate with telegram, a online chat service (like slack or whatsapp).</p>
                <p>Please note that if the dispatch queue of your plugin is full (due to high loads) a message may be silently blocked.</p>

                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      telegram.sendMsg(
                      <i class="amber">chat ID</i> <sup style="color: #444444;">int</sup>
                      ,
                      <i class="amber">message</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      Sends a message from the bot to the chat/group described by <i class="green">chat ID</i>.
                      <p>EG: </p>
                      <pre>
                        telegram.sendMsg(JSON.parse(data.get('chatid').Content), "Hi!");
                      </pre>
                    </v-pane-content>
                  </v-pane>


                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      telegram.onChatJoined(
                      <i class="amber">method name</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      Tells the system to call the method specified by <i class="green">method name</i>
                      whenever the bot is added to a chat group, or a new private chat is opened.
                      <p>The format of the parameter passed to the method can be found here:
                      <a href="https://godoc.org/github.com/Syfaro/telegram-bot-api#Message">https://godoc.org/github.com/Syfaro/telegram-bot-api#Message</a></p>
                      <p>EG: </p>
                      <pre>
                        function onChatJoined(msg) {
                            telegram.sendMsg(msg.Chat.ID, "Hi!");
                        }

                        telegram.onChatJoined("onChatJoined");
                      </pre>
                    </v-pane-content>
                  </v-pane>

                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      telegram.onChatMsg(
                      <i class="amber">method name</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      Tells the system to call the method specified by <i class="green">method name</i>
                      whenever a message is posted in any conversation the bot is in.
                      <p>The format of the parameter passed to the method can be found here:
                      <a href="https://godoc.org/github.com/Syfaro/telegram-bot-api#Message">https://godoc.org/github.com/Syfaro/telegram-bot-api#Message</a></p>

                      <p>EG: </p>
                      <pre>
                        function onChatMsg(msg) {
                            telegram.sendMsg(msg.Chat.ID, "Hi!");
                        }

                        telegram.onChatMsg("onChatMsg");
                      </pre>
                    </v-pane-content>
                  </v-pane>

                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      telegram.onChatLeft(
                      <i class="amber">method name</i> <sup style="color: #444444;">str</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      Tells the system to call the method specified by <i class="green">method name</i>
                      whenever the bot is kicked from a conversation.
                      <p>The format of the parameter passed to the method can be found here:
                      <a href="https://godoc.org/github.com/Syfaro/telegram-bot-api#Message">https://godoc.org/github.com/Syfaro/telegram-bot-api#Message</a></p>
                      <p>EG: </p>
                      <pre>
                        function onChatLeft(msg) {
                          log('chat killed');
                        }

                        telegram.onChatLeft("onChatLeft");
                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>




            <!-- TEMPLATE FUNCTIONS -->
            <v-pane>
              <v-pane-header>
                <md-icon md-font-library="material-icons">description</md-icon>
                Template
              </v-pane-header>

              <v-pane-content>
                <p>These functions allow you to use the golang text/template templating engine.</p>
                <p>The format of the templates is described in: <a href="https://golang.org/pkg/text/template/">https://golang.org/pkg/text/template/</a></p>

                <v-accordion>
                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      template.render(
                      <i class="amber">template</i> <sup style="color: #444444;">str</sup>,
                      <i class="amber">data</i> <sup style="color: #444444;">obj</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      Applies <i class="green">data</i> to the templates specified in <i class="green">template</i>, returning the final result of the render.
                      <p>EG: </p>
                      <pre ng-non-bindable>
                      log(template.render('Hi, {{.name}}', {'name': 'barry'}));
                      </pre>
                    </v-pane-content>
                  </v-pane>



                  <v-pane>
                    <v-pane-header class="green">
                      <md-icon md-font-library="material-icons">code</md-icon>
                      template.renderWeb(
                      <i class="amber">data</i> <sup style="color: #444444;">str</sup>,
                      <i class="amber">options</i> <sup style="color: #444444;">obj</sup>
                      )
                    </v-pane-header>
                    <v-pane-content>
                      <p>Wraps the HTML specified in <i class="green">data</i> in a consistent and pretty looking webpage. Options can be passed to change the appearance of the page.</p>
                      <p>Available options are:</p>
                      <ul>
                        <li>Title - Sets the banner and page title of the page.</li>
                        <li>Icon - Sets the banner icon.</li>
                        <li>PrimaryColour - Sets the primary colour of widgets.</li>
                        <li>AccentColour - Sets the accent colour of material design controls.</li>
                      </ul>
                      <p>EG: </p>
                      <pre ng-non-bindable>
                        function handleHi(web) {
                            log(web.url);
                            log(web.data);
                            web.write(template.renderWeb("LOLZ", {
                                                                    Title: "Test thingy",
                                                                    Icon: "close",
                                                                    AccentColour: "grey"
                                                                }));
                            web.done();
                        }

                        var didBind = web.handle("/p/hi", "handleHi");
                        log("Bound to /p/hi: " + didBind);
                      </pre>
                    </v-pane-content>
                  </v-pane>
                </v-accordion>

              </v-pane-content>
            </v-pane>

                </v-accordion>
        </md-content>
      </div>
    </div>

    {!{template "tailcontent"}!}

    <script>
    (function() {

        var app = angular.module('apiApp', [
          'md.data.table',
          'vAccordion',
          'ngMaterial']);

        //material colour scheme
        app.config(function($mdThemingProvider, $mdIconProvider){
          $mdThemingProvider.theme('default')
                              .primaryPalette('indigo')
                              .accentPalette('amber');
        });

    })();
    </script>

    <script src="/static/js/v-accordion.min.js"></script>
  </body>
</html>
