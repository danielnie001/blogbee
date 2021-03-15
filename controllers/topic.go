package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"myAppNew/models"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (t *TopicController) Get() {
	t.Data["IsLogin"] = Login(t.Ctx)
	t.Data["IsTopic"] = true
	t.TplName = "topic.html"
	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		logs.Error(err)
	} else {
		t.Data["Topics"] = topics
	}
}

func (this *TopicController) Add() {
	var err error
	this.TplName = "topic_add.html"
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		logs.Error(err)
	}
}

func (this *TopicController) Post() {
	if !Login(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	//解析表单
	title := this.GetString("title")
	content := this.GetString("content")
	label := this.GetString("label")
	category := this.GetString("category")
	tid := this.GetString("tid")

	//上传附件
	_, fh, err := this.GetFile("attachment")
	if err != nil {
		logs.Error(err)
	}
	var attachment string
	if fh != nil {
		//保存附件
		attachment = fh.Filename
		logs.Info(attachment)
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			logs.Error(err)
		}
	}
	if len(tid) == 0 {
		err = models.AddTopic(title, category, label, content, attachment)
	} else {
		err = models.ModifyTopic(tid, title, label, category, content, attachment)
	}

	if err != nil {
		logs.Error(err)
	}
	this.Redirect("/topic", 302)
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"
	tid := this.Ctx.Input.Param("0")
	topic, err := models.GetTopic(false, tid)
	if err != nil {
		logs.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["IsLogin"] = Login(this.Ctx)
	this.Data["Topic"] = topic
	this.Data["Labels"] = strings.Split(topic.Labels, " ")
	replies, err := models.GetAllReplies(tid)
	if err != nil {
		logs.Error(err)
		return
	}
	this.Data["Replies"] = replies
}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"
	tid := this.GetString("tid")
	topic, err := models.GetTopic(true, tid)
	if err != nil {
		logs.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["IsLogin"] = Login(this.Ctx)
	this.Data["Tid"] = tid
	this.Data["Topic"] = topic
}

func (this *TopicController) Delete() {
	if !Login(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	err := models.DeleteTopic(this.GetString("tid"))
	if err != nil {
		logs.Error(err)
	}
	this.Redirect("/topic", 302)
}
