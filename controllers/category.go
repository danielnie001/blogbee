package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"myAppNew/models"
)

type CategoryController struct {

	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.GetString("op")
	switch op {
	case "add":
		if !Login(c.Ctx){
			c.Redirect("/login",302)
			return
		}
		name := c.GetString("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			logs.Error(err)
		}
		c.Redirect("/category",302)
		return
	case "del":
		if !Login(c.Ctx){
			c.Redirect("/login",302)
			return
		}
		id :=c.GetString("id")
		if len(id) == 0 {
			break
		}
		err := models.DeleteCategory(id)
		if err != nil {
			logs.Error(err)
		}
	}
	c.Data["IsLogin"] = Login(c.Ctx)
	c.Data["IsCategory"] = true
	c.TplName = "category.html"
	var err error
	c.Data["Categories"],err = models.GetAllCategories()
	if err != nil {
		logs.Error(err)
	}
}
