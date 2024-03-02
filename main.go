package main

import (
	"bank-app/api"
	"bank-app/database"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	database.DB = db

	api.StartApi()
}
