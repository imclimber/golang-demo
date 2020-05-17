package main

import (
	"fmt"

	"github.com/golang-demo/utils"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// ORMDB ...
var ORMDB orm.Ormer

// ConnectToDatabase ...
func ConnectToDatabase() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	dbConfig := utils.Config.DB
	dbConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host,
		dbConfig.Port, dbConfig.DBName)
	orm.RegisterDataBase("default", "mysql", dbConnStr, maxIdle, maxConn)

	// register model
	orm.RegisterModel(new(User))

	// create table
	// 对注册的模型建表
	orm.RunSyncdb("default", false, true)

	// 简单的设置 Debug 为 true 打印查询的语句
	orm.Debug = true

	ORMDB = orm.NewOrm()
	ORMDB.Using("default") // 默认使用 default，你可以指定为其他数据库
}
