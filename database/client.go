package database

import (
	"fmt"
	"log"

	"github.com/ohlmeier/todo-api/todo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Instance *gorm.DB
	err      error
)

func Connect(user, password, host, db string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, db)
	Instance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
