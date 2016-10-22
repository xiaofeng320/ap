package main

import (
	"ap/models"
	_ "ap/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql.username")+
		":"+beego.AppConfig.String("mysql.password")+"@tcp(10.0.11.111:3306)/yc_account1?charset=utf8&parseTime=true&charset=utf8", 30)
	orm.RegisterModel(new(models.Account))
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
