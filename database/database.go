package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
)

func InitDB() *gorm.DB {
	var err error
	dsn := "notes-app-api:12345678$Api@tcp(127.0.0.1:3306)/db-notes-app"
	DBCon, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		//panic("failed to connect database")
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}
	return DBCon
}
