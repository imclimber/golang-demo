package gorm

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // db driver
	//  _ "github.com/jinzhu/gorm/dialects/postgres"
	//  _ "github.com/jinzhu/gorm/dialects/sqlite"
	//  _ "github.com/jinzhu/gorm/dialects/mssql"
)

// 数据库字段修饰说明：
// PK：primary key 主键
// NN：not null 非空
// UQ：unique 唯一索引
// BIN：binary 二进制数据(比text更大)
// UN：unsigned 无符号（非负数）
// ZF：zero fill 填充0 例如字段内容是1 int(4), 则内容显示为0001
// AI：auto increment 自增

// Product ...
type Product struct {
	gorm.Model // gorm 框架自带基础模型
	Code       string
	Price      uint
}

func gormModelProduct() error {
	db, err := gorm.Open("mysql", "list-user:pacman@(47.112.200.141:3306)/entity-list?charset=utf8&parseTime=True&loc=Local&multiStatements=True")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("database connected")
	defer db.Close()

	// // db.LogMode(true)

	// // Migrate the schema
	// db.AutoMigrate(&Product{})

	// // Create
	// err = db.Create(&Product{Code: "L1212", Price: 1000}).Error
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// Read
	// var product = Product{}
	// db.First(&product, 1) // find product with id 1
	// fmt.Println(product)
	// db.First(&product, "code = ?", "L1212") // find product with code l1212
	// fmt.Println(product)

	// // Update - update product's price to 2000
	// err = db.Model(&product).Update("Price", 2000).Error
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// Delete - delete product（软删除）
	// var product = Product{
	// 	Model: gorm.Model{
	// 		ID: 3,
	// 	},
	// }
	// db.Delete(&product)

	return nil
}

// User ...
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string `gorm:"type:varchar(100);unique_index"`
	Role         string `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int    `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int    `gorm:"-"`               // 忽略本字段
}

func gormModelUser() error {
	db, err := gorm.Open("mysql", "list-user:pacman@(47.112.200.141:3306)/entity-list?charset=utf8&parseTime=True&loc=Local&multiStatements=True")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("database connected")
	defer db.Close()

	db.LogMode(true)

	var user = User{}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	// age 字段中必须 valid 字段为 true 才能插入成功
	// err = db.Create(&User{MemberNumber: "aaa", Age: sql.NullInt64{Int64: 30, Valid: true}}).Error
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// Read
	user = User{}
	db.First(&user, 1) // find product with id 1
	fmt.Println(user)
	if user.Age.Valid {
		fmt.Println(user.Age)
	} else {
		fmt.Println("age is nil")
	}

	// Update
	err = db.Model(&user).Update("Age", 40).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	// // Delete - delete user（软删除）
	// user = User{
	// 	Model: gorm.Model{
	// 		ID: 1,
	// 	},
	// }
	// db.Delete(&user)

	return nil
}
