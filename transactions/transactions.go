package transactions

import (
	"bank-app/database"
	"bank-app/helpers"
	"bank-app/interfaces"
	"sort"
)

func CreateTransaction(From uint, To uint, Amount int, Currency string) {
	transaction := &interfaces.Transaction{From: From, To: To, Amount: Amount, Currency: Currency}
	database.DB.Create(&transaction)
}

func GetTransactionsByAccount(id uint) []interfaces.ResponseTransaction {
	transactions := []interfaces.ResponseTransaction{}
	database.DB.Table("transactions").Select("id, transactions.from, transactions.to, amount, currency").Where(interfaces.Transaction{From: id}).Or(interfaces.Transaction{To: id}).Scan(&transactions)
	return transactions
}

func GetMyTransactions(id string, jwt string, sortBy string, currency string, offset int, pageSize int) map[string]interface{} {
	// Validate JWT
	isValid := helpers.ValidateToken(id, jwt)
	if isValid {
		// Find and return transactions
		accounts := []interfaces.ResponseAccount{}
		database.DB.Table("accounts").Select("id, name, balance_kzt, balance_eur, balance_usd").Where("user_id = ? ", id).Scan(&accounts)

		transactions := []interfaces.ResponseTransaction{}
		for i := 0; i < len(accounts); i++ {
			accTransactions := GetTransactionsByAccount(accounts[i].ID)
			transactions = append(transactions, accTransactions...)
		}

		if currency != "" {
			filteredTransactions := []interfaces.ResponseTransaction{}
			for _, transaction := range transactions {
				if transaction.Currency == currency {
					filteredTransactions = append(filteredTransactions, transaction)
				}
			}
			transactions = filteredTransactions
		}

		if sortBy == "amount" {
			sort.Slice(transactions, func(i, j int) bool {
				return transactions[i].Amount < transactions[j].Amount
			})
		}

		// Perform pagination
		start := offset
		end := offset + pageSize
		if start > len(transactions) {
			start = len(transactions)
		}
		if end > len(transactions) {
			end = len(transactions)
		}
		transactions = transactions[start:end]

		var response = map[string]interface{}{"message": "all is fine"}
		response["data"] = transactions
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
