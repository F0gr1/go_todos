package main

import (
	"todo/db"
	"todo/model"
)

func main() {
	db := db.Connection()
	defer db.Close()

	db.AutoMigrate(&model.Login{})
	db.AutoMigrate(&model.Todo{})
}
