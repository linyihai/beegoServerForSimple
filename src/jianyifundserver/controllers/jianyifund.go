package controllers

import (
	"jianyifundserver/models"
	"time"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
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
		valid := validation.Validation{}
		b, err := valid.Valid(&ob)
		if err != nil {
			this.Data["json"] = err.Error()
			this.ServeJSON()
			return
		}
		if !b {
			validFailMsg := "" 
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
				validFailMsg = validFailMsg + err.Key + " " + err.Message + ";"
			}
			this.Data["json"] = validFailMsg
		} else if created, id, err := o.ReadOrCreate(&ob, "MailAddress"); err == nil {
			if created {
				log.Println("New Insert an object. Id:", id)
			} else {
				ob.Name = name
				ob.MailAddress = mailAddress
				ob.Desc = desc
				ob.Title = title
				if num, err := o.Update(&ob); err == nil {
					log.Println(num)
				}
			}
			if data, err := json.Marshal(&ob); err == nil {
				this.Data["json"] = data 
			} else {
				this.Data["json"] = err.Error()
			}
		}
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}
