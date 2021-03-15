package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"myAppNew/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "index.html"
	c.Data["IsLogin"] = Login(c.Ctx)
	topics, err := models.GetAllTopics(c.GetString("label"),c.GetString("cate"),true)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["Topics"] = topics
	}
	categories,err:=models.GetAllCategories()
	if err != nil {
		logs.Error(err)
	}
	c.Data["Categories"] = categories
}
