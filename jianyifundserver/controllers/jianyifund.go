package controllers

import (
	"fmt"
	"jianyifundserver/models"
	"time"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var o = orm.NewOrm()

type UserContractController struct {
	beego.Controller
}

func (this *UserContractController) Post() {
	var ob models.UserContractInfo
	var err error

	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err == nil {

		ob.Created = time.Now()
		ob.Updated = time.Now()
		name := ob.Name
		mailAddress := ob.MailAddress
		desc := ob.Desc
		title := ob.Title
		fmt.Println(ob)
		if created, id, err := o.ReadOrCreate(&ob, "MailAddress"); err == nil {
			if created {
				fmt.Println("New Insert an object. Id:", id)
			} else {
				ob.Name = name
				ob.MailAddress = mailAddress
				ob.Desc = desc
				ob.Title = title
				if num, err := o.Update(&ob); err == nil {
					fmt.Println(num)
				}
			}
			this.Data["json"] = "更新成功"
		}
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}
