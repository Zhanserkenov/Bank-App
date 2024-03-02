package database

import (
	"bank-app/interfaces"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=bank-app password=19660827 sslmode=disable")
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(200)
	return db, nil
}

func CreateUser(user *interfaces.User) error {
	return DB.Create(user).Error
}

func CreateAccount(account *interfaces.Account) error {
	return DB.Create(account).Error
}

func GetUserByID(id uint) (*interfaces.User, error) {
	user := &interfaces.User{}
	err := DB.First(user, id).Error
	return user, err
}

func UpdateUser(user *interfaces.User) error {
	return DB.Save(user).Error
}

func DeleteUser(user *interfaces.User) error {
	return DB.Delete(user).Error
}
