package main

import (
	/**
	1. 引入mysql驱动,和 gorm所需的包
	*/
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id   uint `gorm:"primary_key";AUTO_INCREMENT`
	Name string
	Age  int
}

type Product struct {
	Id    uint `gorm:"primary_key";AUTO_INCREMENT`
	Name  string
	Price int
}

type Order struct {
	Id   uint `gorm:"primary_key";AUTO_INCREMENT`
	Name string
	No   int
}

func main() {
	db := ConnectMysql()
	defer db.Close()

	fmt.Println(db.HasTable(&User{}))
	fmt.Println(db.HasTable(&Product{}))
	fmt.Println(db.HasTable(&User{}))

	AutoMigration(db)

	fmt.Println(db.HasTable(&User{}))
	fmt.Println(db.HasTable(&Product{}))
	fmt.Println(db.HasTable(&User{}))
}

/**
2.连接对应的数据库地址，格式如下
user:password@/dbname?charset=utf8&parseTime=True&loc=Local
*/
func ConnectMysql() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:13435011052-mysqlMYSQL@(caojunhui.com:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return
}

/**
3. 设置GORM中对应的模型（视图）自动新增到数据库（自动创建对应的表、自动创建索引，出于数据保护的目的不会修改列属性）
*/
func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(&Product{}, &Order{})

	// 也可以在创建表时指定一些表的配置信息
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
}

/**
4. 检查是否存在表
*/
func checkTable(db *gorm.DB, i interface{}, tableName string) bool {
	// Check model `User`'s table exists or not
	return db.HasTable(i)
}

/**
5. 插入表
*/
func createTable(db *gorm.DB, i interface{}) *gorm.DB {
	// Create table for model `User`
	if db.HasTable(i) == true {
		fmt.Println("table already exist !")
		return db
	}
	db.Create(&User{})
	return db.CreateTable(&User{})
	// will append "ENGINE=InnoDB" to the SQL statement when creating table `users`
	// db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
}
