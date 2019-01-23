package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego"
)

type UserContractInfo struct {
	Id          int
	Name        string    `json:"name" valid:"Required"`
	MailAddress string    `orm:"unique" json:"mailaddress" valid:"Email; MaxSize(100); Required"`
	Title       string    `json:"title" valid:"Required"`
	Desc        string    `json:"desc" valid:"Required"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)" `
}

type UserRiskAssessmentInfo struct {
	Id          int
	Name        string    `json:"name" valid:"Required"`
	MailAddress string    `json:"mailAddress" orm:"unique" valid:"Email; MaxSize(100); Required"`
	UserIDNum   string    `json:"userIDNumber" valid:"Required"`
	IDType      int    	  `json:"idType" valid:"Required" description:"1: 个人;2:机构"`
	InvestorType int      `json:"investorType" valid:"Required" description:"1: 身份证;2:护照"` // 注意数值为0是，验证时会不通过
	AnswerDetail string   `orm:type(json) valid:"Required"`
	Score       int    	  `valid:"Required"`
	ScoreGrade  string    `valid:"Required"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)" `
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("postgresqlgo"))
	orm.RegisterModel(new(UserContractInfo), new(UserRiskAssessmentInfo))
	//自动创建表 参数二为是否开启创建表   参数三是否更新表
	// orm.RunSyncdb("default", true, true)
	orm.Debug = true
}
