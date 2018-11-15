package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	)

type User struct {
	Id int
	UserName string
	Pwd string
}
func init(){
	//生成表的三步骤
	//注册数据库
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/newsweb?charset=utf8")
	//注册表
	orm.RegisterModel(new(User))
	//让项目跑起来
	orm.RunSyncdb("default",false,true)

}