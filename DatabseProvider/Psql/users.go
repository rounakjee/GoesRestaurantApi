package Psql

import (
	"GoResturantAPI/models"
	"github.com/jinzhu/gorm"
)

func (psql *PSqlDB) CreateUser(user *models.User) error {
	return psql.Create(user).Error
}
func (psql *PSqlDB) AllUsers() []models.User {
	var users []models.User
	psql.Find(&users)
	return users
}
func (psql *PSqlDB) LoginUser(user *models.User) (bool, error) {
	if err := psql.Where(&models.User{PhoneNumber: user.PhoneNumber}).Find(&user).Error; gorm.IsRecordNotFoundError(err) {
		return false, err
	} else {
		return true, nil
	}
}
