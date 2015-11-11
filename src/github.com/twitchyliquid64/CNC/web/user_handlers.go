package web

import (
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/data"
  "github.com/hoisie/web"
  "encoding/json"
)

func addPermissionUserHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-user", "addPermission() called unauthorized, aborting")
    return
  }

  username := ctx.Params["username"]
  success, usr := user.GetByUsername(username, data.DB)

  if !success {
    ctx.Abort(500, "ERROR NOT FOUND")
    return
  }

  usr.Permissions = append(usr.Permissions, user.Permission{Name: ctx.Params["perm"]})
  data.DB.Save(&usr)
}

func deletePermissionUserHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-user", "addPermission() called unauthorized, aborting")
    return
  }

  username := ctx.Params["username"]
  success, usr := user.GetByUsername(username, data.DB)

  if !success {
    ctx.Abort(500, "ERROR NOT FOUND")
    return
  }

  //TODO: Refactor so DB code is in model/user rather than in the handler
  data.DB.Where("user_id = ? and name = ?", usr.ID, ctx.Params["perm"]).Delete(&user.Permission{})
  ctx.ResponseWriter.Write([]byte("GOOD"))
}


func deleteUserHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-user", "deleteUser() called unauthorized, aborting")
    return
  }

  username := ctx.Params["username"]

  //TODO: Refactor so DB code is in model/user rather than in the handler
  data.DB.Where("username = ?", username).Delete(&user.User{})
  ctx.ResponseWriter.Write([]byte("GOOD"))
}

func resetPasswordHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-user", "resetPassword() called unauthorized, aborting")
    return
  }

  username := ctx.Params["username"]
  success, usr := user.GetByUsername(username, data.DB)

  if !success {
    ctx.Abort(500, "ERROR NOT FOUND")
    return
  }

  var authMethod user.AuthenticationMethod
  data.DB.Where("user_id = ? and method_type = ?", usr.ID, user.AUTH_PASSWD).First(&authMethod)
  authMethod.Value = ctx.Params["pass"]
  data.DB.Save(&authMethod)
}

func newUserHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-user", "newUser() called unauthorized, aborting")
    return
  }

  decoder := json.NewDecoder(ctx.Request.Body)
  var usr user.User
  err := decoder.Decode(&usr)
  if err != nil {
      logging.Error("web-user", "newUserHandlerAPI() failed to decode JSON:", err)
      ctx.Abort(500, "JSON error")
      return
  }

  data.DB.Create(&usr)
  ctx.ResponseWriter.Write([]byte("GOOD"))
}

func updateUserHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-user", "updateUser() called unauthorized, aborting")
    return
  }

  decoder := json.NewDecoder(ctx.Request.Body)
  var usr user.User
  err := decoder.Decode(&usr)
  if err != nil {
      logging.Error("web-user", "updateUserHandlerAPI() failed to decode JSON:", err)
      ctx.Abort(500, "JSON error")
      return
  }

  logging.Info("web-user", usr)
  data.DB.Save(&usr)
  ctx.ResponseWriter.Write([]byte("GOOD"))
}

func getUserHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-user", "getUsers() called unauthorized, aborting")
    return
  }

  username := ctx.Params["username"]
  success, usr := user.GetByUsername(username, data.DB)

  if !success {
    ctx.Abort(500, "ERROR NOT FOUND")
    return
  }

  user.LoadEphemeral(&usr, data.DB) //populates all the addresses/emails

  d, err := json.Marshal(usr)
  if err != nil {
    logging.Error("web-user", err)
  }
  ctx.ResponseWriter.Write(d)
}

func getUsersHandlerAPI(ctx *web.Context) {
  isLoggedIn, usr, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!usr.IsAdmin()){
    logging.Warning("web-user", "getUsers() called unauthorized, aborting")
    return
  }

  users := user.GetAll(data.DB)
  d, err := json.Marshal(users)
  if err != nil {
    logging.Error("web-user", err)
  }
  ctx.ResponseWriter.Write(d)
}

func loginHandler(ctx *web.Context) {

  usrname := ctx.Params["user"]
  passwd := ctx.Params["pass"]

  isValidLogin, usr := user.CheckAuthByPassword(usrname, passwd, data.DB)

  if isValidLogin {
    logging.Info("web", "User '", usrname, "' has authenticated.")
    skey := session.CreateSession(int(usr.ID), "web", data.DB)
    ctx.SetCookie(web.NewCookie(COOKIE_KEY_NAME, skey, 60*60*24*20))
    ctx.ResponseWriter.Write([]byte("GOOD"))
  }else{
    ctx.Abort(500, "ERROR NOT FOUND")
  }
}


func logoutHandler(ctx *web.Context) {
  isLoggedIn, user, s := getSessionByCookie(ctx)
  if isLoggedIn {
    logging.Info("web", "Now logging out:", user.Username)
    session.Delete(s, data.DB)
    deleteSessionKey(ctx)
  } else {
    logging.Warning("web", "/logout called with an invalid session!")
  }
  ctx.Redirect(302, "/")
}
