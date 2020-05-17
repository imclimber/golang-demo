package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/golang-demo/utils"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(100)"`
}

var configFilename string
var configDirs string

func init() {
	const (
		defaultConfigFilename = "config"
		configUsage           = "Name of the config file, without extension"
		defaultConfigDirs     = "./configs/"
		configDirUsage        = "Directories to search for config file, separated by ','"
	)
	flag.StringVar(&configFilename, "f", defaultConfigFilename, configUsage)
	flag.StringVar(&configDirs, "p", defaultConfigDirs, configDirUsage)

	err := utils.NewConfiguration(configFilename, strings.Split(configDirs, ","))
	if err != nil {
		panic(fmt.Errorf("Error parsing config, %s", err))
	}

	ConnectToDatabase()
}

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

func main() {
	var err error

	// // insert
	// user := User{Name: "slene"}
	// id, err := ORMDB.Insert(&user)
	// if err != nil {
	// 	fmt.Errorf("insert error: %v", err)
	// 	return
	// }
	// fmt.Println("insert return id:", id)

	// insertMulti
	userOne := User{Name: "oneName"}
	userTwo := User{Name: "twoName"}
	var users = []User{userOne, userTwo}
	num, err := ORMDB.InsertMulti(len(users), &users)
	if err != nil {
		fmt.Errorf("InsertMulti error: %v", err)
		return
	}
	fmt.Println("InsertMulti return num:", num)

	// // update
	// user := User{Id: 1, Name: "astaxie"}
	// num, err := ORMDB.Update(&user)
	// if err != nil {
	// 	fmt.Errorf("update error: %v", err)
	// 	return
	// }
	// fmt.Println("update return num:", num)

	// read one
	// u := User{Id: 1}
	// err = ORMDB.Read(&u)
	// if err != nil {
	// 	fmt.Errorf("read error: %v", err)
	// 	return
	// }
	// fmt.Println("read result:", u)

	// delete
	// u := User{Id: 1}
	// num, err := ORMDB.Delete(&u)
	// if err != nil {
	// 	fmt.Errorf("delete error: %v", err)
	// 	return
	// }
	// fmt.Println("delete return num:", num)
}
