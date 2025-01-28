package database

import (
	"fmt"
	"log"

	"github.com/ohlmeier/todo-api/todo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Init(user, password, host, db string) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s", user, password, host, db)
	var err error
	Instance, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	Instance.AutoMigrate(&todo.List{})
	Instance.AutoMigrate(&todo.Item{})
}
