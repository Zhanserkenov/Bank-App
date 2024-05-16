package database

import (
	"bank-app/pkg/helpers"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Структура данных для таблицы пользователей
type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Username  string     `gorm:"not null"`
	Email     string     `gorm:"not null;unique"`
	Password  string     `gorm:"not null"`
}

// Структура данных для таблицы счетов
type Account struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	Type       string     `gorm:"not null"`
	Name       string     `gorm:"not null"`
	UserID     uint       `gorm:"not null"`
	BalanceKZT int        `gorm:"not null;default:0"`
	BalanceUSD int        `gorm:"not null;default:0"`
	BalanceEUR int        `gorm:"not null;default:0"`
}

// Структура данных для таблицы транзакций
type Transaction struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	From      uint       `gorm:"not null"`
	To        uint       `gorm:"not null"`
	Amount    int        `gorm:"not null"`
	Currency  string
}

// Создание глобальной переменной для базы данных
var DB *gorm.DB

// Функция для инициализации базы данных
func InitDatabase() {
	// Получение переменных окружения для подключения к базе данных
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASS")
	if dbPassword == "" {
		dbPassword = "19660827"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "bank-app"
	}

	// Формирование строки подключения к базе данных
	connectionString := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", dbHost, dbUser, dbName, dbPassword)

	// Подключение к базе данных
	database, err := gorm.Open("postgres", connectionString)
	helpers.HandleErr(err)

	// Установка соединения с базой данных
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)

	// Создание таблиц на основе структур данных
	database.AutoMigrate(&User{}, &Account{}, &Transaction{})

	DB = database
}
