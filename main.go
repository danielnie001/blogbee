package main

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"myAppNew/models"
	_ "myAppNew/routers"
	"os"
)

func init() {
	models.RegisterDB()
}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default",false,true)

	//创建目录
	os.MkdirAll("attachment",os.ModePerm)
	//beego.SetStaticPath("/attachment","attachment")
	beego.Run()
}

