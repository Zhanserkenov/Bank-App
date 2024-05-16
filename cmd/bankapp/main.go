package main

import (
	"bank-app/api"
	"bank-app/internal/database"
)

func main() {
	//migrations.Migrate()

	database.InitDatabase()
	api.StartApi()
}
