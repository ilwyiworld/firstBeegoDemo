package main

import (
	_ "firstBeegoDemo/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego"
)

type Userdetail struct {
	Id int
	Intro string
	Profile string
}
func init() {
	orm.RegisterDataBase("default","mysql","root:root@/godatabases?charset=utf8",30)
	orm.RegisterModel(new(Userdetail))
}

func main() {
	o:=orm.NewOrm()
	u:=Userdetail{Id:1}
	err:=o.Read(&u)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(u)
	}

	appname:=beego.AppConfig.String("appname")
	runmode:=beego.AppConfig.String("runmode")
	httpport,err:=beego.AppConfig.Int("httpport")
	testport,err:=beego.AppConfig.Int("test::httpport")	//获取指定section中的属性值
	fmt.Println(appname+";"+runmode)
	fmt.Println(testport)
	if err==nil{
		fmt.Println(httpport)
	}else{
		fmt.Println(err)
	}

	conf,err:=config.NewConfig("json","conf/testConf.json")
	if err==nil{
		fmt.Println(conf.String("dev::appname"))
	}else{
		fmt.Println(err)
	}

	beego.Run()
}

