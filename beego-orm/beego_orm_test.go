package beego_orm

import (
	"flag"
	"fmt"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-demo/utils"
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
		defaultConfigDirs     = "../configs/"
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

func TestConnectToDatabase(t *testing.T) {
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
