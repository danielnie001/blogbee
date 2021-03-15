package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

type LoginController struct {
	web.Controller
}

func (this *LoginController) Get() {
	isExit := this.GetString("exit") == "true"
	fmt.Println(isExit)
	if isExit {
		this.Ctx.SetCookie("uname","",-1,"/")
		this.Ctx.SetCookie("pwd","",-1,"/")
		this.Redirect("/",302)
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	uname := this.GetString("uname")
	pwd := this.GetString("passwd")
	autoLogin := this.GetString("autoLogin") == "on"
	appUname, _ := web.AppConfig.String("uname")
	appPwd, _ := web.AppConfig.String("pwd")
	if appUname == uname &&
		appPwd == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<32 - 1
		}
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	this.Redirect("/", 302)
	return
}

func Login (ctx *context.Context) bool {
	uname := ctx.GetCookie("uname")
	pwd := ctx.GetCookie("pwd")
	if uname == "" || pwd == "" {
		return false
	}
	appUname,_ := web.AppConfig.String("uname")
	appPwd,_ := web.AppConfig.String("pwd")
	return appUname == uname && appPwd == pwd
}
