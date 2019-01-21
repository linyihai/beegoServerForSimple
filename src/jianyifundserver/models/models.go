package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
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

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:mysql@tcp(127.0.0.1:3306)/jianyifund?charset=utf8")
	orm.RegisterModel(new(UserContractInfo))
	//自动创建表 参数二为是否开启创建表   参数三是否更新表
	// orm.RunSyncdb("default", true, true)
	orm.Debug = true
}
