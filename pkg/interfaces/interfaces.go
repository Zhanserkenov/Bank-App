package interfaces

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type       string
	Name       string
	BalanceKZT uint
	BalanceUSD uint
	BalanceEUR uint
	UserID     uint
}

type Transaction struct {
	gorm.Model
	From     uint
	To       uint
	Amount   int
	Currency string
}

type ResponseTransaction struct {
	ID       uint
	From     uint
	To       uint
	Amount   int
	Currency string
}

type ResponseAccount struct {
	ID      uint
	Name    string
	Balance int
}

type ResponseUser struct {
	ID       uint
	Username string
	Email    string
	Accounts []ResponseAccount
}

type Validation struct {
	Value string
	Valid string
}

type ErrResponse struct {
	Message string
}
