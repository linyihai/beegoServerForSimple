package controllers

import (
	"jianyifundserver/models"
	"jianyifundserver/service"
	"time"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

var o = orm.NewOrm()

const DONECODE = 0
const FAILCODE = -1
type RetJSON struct {
	ResCode int
	Data interface{}
	FailMsg string
}


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
			this.Data["json"] = RetJSON{FailMsg:err.Error(), ResCode:FAILCODE}
			this.ServeJSON()
			return
		}
		if !b {
			validFailMsg := "" 
			for _, err := range valid.Errors {
				beego.Warning(err.Key, err.Message)
				validFailMsg = validFailMsg + err.Key + " " + err.Message + ";"
			}
			this.Data["json"] = validFailMsg
		} else if created, id, err := o.ReadOrCreate(&ob, "MailAddress"); err == nil {
			if created {
				beego.Informational("New Insert an object. Id:", id)
			} else {
				ob.Name = name
				ob.MailAddress = mailAddress
				ob.Desc = desc
				ob.Title = title
				if _, err := o.Update(&ob); err == nil {
					beego.Debug("update new user contract info")
				}
			}
			this.Data["json"] = RetJSON{Data:ob, ResCode:FAILCODE}
		}
	} else {
		this.Data["json"] = RetJSON{FailMsg:err.Error(), ResCode:FAILCODE}
	}
	this.ServeJSON()
}

type UserRiskAssessmentController struct {
	beego.Controller
}

func (this *UserRiskAssessmentController) Post() {
	var ob models.UserRiskAssessmentInfo
	var err error

	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err == nil {
		AnswerDetail := string(this.Ctx.Input.RequestBody)
		score, scoreGrade := service.CalRiskAssessment(AnswerDetail, ob.InvestorType)

		ob.Score = score
		ob.ScoreGrade = scoreGrade
		ob.Created = time.Now()
		ob.Updated = time.Now()
		ob.AnswerDetail = AnswerDetail

		Name := ob.Name
		MailAddress := ob.MailAddress
		UserIDNum := ob.UserIDNum
		IDType := ob.IDType
		InvestorType := ob.InvestorType

		valid := validation.Validation{}
		b, err := valid.Valid(&ob)
		beego.Debug(b)
		if err != nil {
			this.Data["json"] = RetJSON{FailMsg:err.Error(), ResCode:FAILCODE}
			this.ServeJSON()
			return
		}
		
		if !b {
			validFailMsg := "" 
			for _, err := range valid.Errors {
				beego.Warning(err.Key, err.Message)
				validFailMsg = validFailMsg + err.Key + " " + err.Message + ";"
			}
			this.Data["json"] = validFailMsg
		} else if created, id, err := o.ReadOrCreate(&ob, "MailAddress"); err == nil {
			if created {
				beego.Informational("New Insert an object. Id:", id)
			} else {
				ob.Name = Name
				ob.MailAddress = MailAddress
				ob.UserIDNum = UserIDNum 
				ob.IDType = IDType 
				ob.InvestorType = InvestorType 
				ob.AnswerDetail = AnswerDetail 
				ob.Score = score
				ob.ScoreGrade = scoreGrade
				if _, err := o.Update(&ob); err == nil {
					beego.Debug("update new user risk assessment info")
				}
			}
			this.Data["json"] = RetJSON{ResCode: DONECODE, Data: ob}
		}
	} else {
		this.Data["json"] = RetJSON{FailMsg:err.Error(), ResCode:FAILCODE}
	}
	this.ServeJSON()
}

func (this *UserRiskAssessmentController) Get() {
	
	mailAddress := this.GetString("mailAddress")
	if mailAddress != "" {
		ob := models.UserRiskAssessmentInfo{MailAddress: mailAddress} 
		err :=  o.Read(&ob, "MailAddress")
		if err != nil {
			this.Data["json"] = RetJSON{FailMsg:"未找到邮箱对应的记录", ResCode:FAILCODE}
		} else {
			this.Data["json"] = RetJSON{ResCode: DONECODE, Data: ob}
		}
	} else {
		this.Data["json"] = RetJSON{FailMsg:"邮箱为空", ResCode:FAILCODE}
	}
	this.ServeJSON()
}
