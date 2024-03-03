package main

import (
	"bank-app/pkg/api"
	"bank-app/pkg/database"
)

func main() {
	//migrations.Migrate()

	database.InitDatabase()
	api.StartApi()
}
