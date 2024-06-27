package dal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var mysqlLogger logger.Interface

func Initialize() {
	initDB()
}

func initDB() {
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
	DB = DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
	fmt.Printf("init db successfully: %v\n", DB)
}
