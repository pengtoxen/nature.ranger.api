package main

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "nature.ranger.api/initial"
	_ "nature.ranger.api/routers"
	"net/http"
)

type res struct {
	code interface{}
	msg  string
}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("userInfo").(string)
	if !ok && ctx.Request.RequestURI != "/login/login" {
		//ctx.Output.JSON(map[string]interface{}{"code": 1, "msg": "please login"}, false, false)
	}
}

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.ErrorHandler("404", apiNotFound)
	beego.ErrorHandler("401", apiNotPermission)
	beego.Run()
}

func apiNotFound(rw http.ResponseWriter, r *http.Request) {
	apiFoundFail(rw, r, "api not api not found")
}

func apiNotPermission(rw http.ResponseWriter, r *http.Request) {
	apiFoundFail(rw, r, "api not permission")
}

func apiFoundFail(rw http.ResponseWriter, r *http.Request, msg string) {
	var ret res = res{code: 1, msg: msg}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(ret)
}
