package useraccounts

import (
	"fmt"

	"bank-app/internal/database"
	"bank-app/internal/transactions"
	"bank-app/pkg/helpers"
	"bank-app/pkg/interfaces"
)

func getAccount(id uint) *interfaces.Account {
	account := &interfaces.Account{}
	if database.DB.Where("id = ? ", id).First(&account).RecordNotFound() {
		return nil
	}
	return account
}

func Transaction(userId uint, from uint, to uint, amount int, currency string, jwt string) map[string]interface{} {

	isValid := helpers.ValidateToken(fmt.Sprint(userId), jwt)
	if isValid {

		fromAccount := getAccount(from)
		toAccount := getAccount(to)

		if fromAccount == nil || toAccount == nil {
			return map[string]interface{}{"message": "Account not found"}
		} else if fromAccount.UserID != userId {
			return map[string]interface{}{"message": "You are not the owner of the account"}
		}

		var fromBalance uint
		switch currency {
		case "KZT":
			fromBalance = fromAccount.BalanceKZT
		case "USD":
			fromBalance = fromAccount.BalanceUSD
		case "EUR":
			fromBalance = fromAccount.BalanceEUR
		default:
			return map[string]interface{}{"message": "Invalid currency"}
		}

		if fromBalance < uint(amount) {
			return map[string]interface{}{"message": "Not enough balance"}
		}

		switch currency {
		case "KZT":
			fromAccount.BalanceKZT -= uint(amount)
			toAccount.BalanceKZT += uint(amount)
		case "USD":
			fromAccount.BalanceUSD -= uint(amount)
			toAccount.BalanceUSD += uint(amount)
		case "EUR":
			fromAccount.BalanceEUR -= uint(amount)
			toAccount.BalanceEUR += uint(amount)
		}

		database.DB.Save(&fromAccount)
		database.DB.Save(&toAccount)

		transactions.CreateTransaction(from, to, amount, currency)

		return map[string]interface{}{"message": "Transaction successful"}
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}

func ConvertCurrency(userID uint, fromCurrency string, toCurrency string, amount uint) error {

	account := getAccount(userID)

	switch fromCurrency {
	case "KZT":
		if account.BalanceKZT < amount {
			return nil
		}
		account.BalanceKZT -= amount
	case "USD":
		if account.BalanceUSD < amount {
			return nil
		}
		account.BalanceUSD -= amount
	case "EUR":
		if account.BalanceEUR < amount {
			return nil
		}
		account.BalanceEUR -= amount
	}

	amountFloat := float64(amount)

	exchangeRates := map[string]float64{
		"KZT": 1,
		"USD": 0.0022,
		"EUR": 0.0020,
	}

	convertedAmount := uint(amountFloat / exchangeRates[fromCurrency] * exchangeRates[toCurrency])

	switch toCurrency {
	case "KZT":
		account.BalanceKZT += convertedAmount
	case "USD":
		account.BalanceUSD += convertedAmount
	case "EUR":
		account.BalanceEUR += convertedAmount
	}

	if err := database.DB.Save(&account).Error; err != nil {
		return err
	}

	return nil
}
