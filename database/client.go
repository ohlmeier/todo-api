package database

import (
	"log"

	"github.com/ohlmeier/todo-api/todo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Instance *gorm.DB
	err      error
)

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&todo.List{})
	Instance.AutoMigrate(&todo.Item{})
	log.Println("Database Migration Completed...")
}
