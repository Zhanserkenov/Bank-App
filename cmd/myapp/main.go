package main

import (
	"bank-app/api"
	"bank-app/database"
)

func main() {
	//migrations.Migrate()

	database.InitDatabase()
	api.StartApi()
}
