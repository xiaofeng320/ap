package controllers

import (
	"ap/service"
	"strconv"

	"github.com/astaxie/beego"
)

type AccountController struct {
	beego.Controller
}

func (c *AccountController) Find() {
	//id := c.Ctx.Input.Param("account_id")
	id := c.GetString("account_id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	t, ok := service.Account.Find(intid)
	//t, ok := models.DefaultTaskList.Find(intid)
	beego.Info("Found", ok)

	if !ok {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte("task not found"))
		return
	}
	c.Data["json"] = &t
	c.ServeJSON()
}
