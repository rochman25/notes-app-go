package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"notes-app-go/config"
)

var (
	DBCon *gorm.DB
)

func InitDB() *gorm.DB {
	var err error
	dsn := config.GetEnv("DB_USERNAME") + ":" + config.GetEnv("DB_PASSWORD") + "@tcp(" + config.GetEnv("DB_HOSTNAME") + ":3306)/" + config.GetEnv("DB_DATABASE")
	DBCon, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//	//panic("failed to connect database")
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}
	return DBCon
}
