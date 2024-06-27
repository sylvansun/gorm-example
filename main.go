package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm_example/model"
)

var DB *gorm.DB
var mysqlLogger logger.Interface

func init() {
	username := "root"
	password := "root"
	host := "localhost"
	port := 3306
	dbName := "gorm"
	timeout := "10s"

	mysqlLogger = logger.Default.LogMode(logger.Info)

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbName, timeout)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false, // default setting is also false
	})
	if err != nil {
		panic("connect to database failed, error=" + err.Error())
	}
	DB = db
}

func main() {
	DB = DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
	err := DB.AutoMigrate(&model.Student{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = DB.Create(model.CreateStudent()).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	err = DB.Create(model.CreateStudentInBatch()).Error
	if err != nil {
		fmt.Println(err)
		return
	}
}
