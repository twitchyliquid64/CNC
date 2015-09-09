package web

import (
  "github.com/hoisie/web"
)

func loginHandler(ctx *web.Context) {
  ctx.ResponseWriter.Write([]byte("LOL"))
}
