package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"myAppNew/models"
)

type ReplyController struct {
	web.Controller
}

func (this *ReplyController) Get() {

}

func (this *ReplyController) Add() {
	tid := this.GetString("tid")
	err := models.AddReplay(tid, this.GetString("nickname"), this.GetString("content"))
	if err != nil {
		logs.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
}

func (this *ReplyController) Delete() {
	if !Login(this.Ctx) {
		return
	}
	tid := this.GetString("tid")
	err := models.DeleteReply(this.GetString("rid"))
	if err != nil {
		logs.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
}
