package migrations

import (
	"fmt"
	"notes-app-go/database"
	"notes-app-go/src/models"
	"reflect"
	"strings"
)

var tables = []interface{}{
	&models.Notes{},
	&models.Users{},
	&models.Authentications{},
}

func Migrate() {
	database.DBCon.AutoMigrate(tables...)
}

func Rollback() {
	for i := 0; i < len(tables); i++ {
		database.DBCon.Migrator().DropTable(tables[i])
	}
}

func Status() {
	var (
		conn        = database.DBCon
		colorReset  = "\033[0m"
		colorGreen  = "\033[32m"
		colorYellow = "\033[33m"
	)

	fmt.Printf("In database %s:\n", conn.Migrator().CurrentDatabase())
	for _, table := range tables {
		var name string

		t := reflect.TypeOf(table)
		if t.Kind() == reflect.Ptr {
			name = strings.ToLower(t.Elem().Name())
		} else {
			name = strings.ToLower(t.Name())
		}

		if conn.Migrator().HasTable(table) {
			fmt.Println("\t", name, "===>", string(colorGreen), "migrated", string(colorReset))
		} else {
			fmt.Println("\t", name, "===>", string(colorYellow), "not migrated", string(colorReset))
		}
	}
}
