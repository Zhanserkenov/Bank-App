package migrations

import (
	"bank-app/pkg/database"
	"bank-app/pkg/helpers"
	"bank-app/pkg/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createAccounts() {
	users := &[2]interfaces.User{
		{Username: "Martin", Email: "martin@martin.com"},
		{Username: "Michael", Email: "michael@michael.com"},
	}
	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		database.DB.Create(&user)

		account := &interfaces.Account{
			UserID:     user.ID,
			BalanceKZT: uint(10000 * (i + 1)),
			BalanceUSD: 0,
			BalanceEUR: 0,
		}
		database.DB.Create(&account)
	}
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	Transactions := &interfaces.Transaction{}
	database.DB.AutoMigrate(&User, &Account, &Transactions)

	createAccounts()
}
