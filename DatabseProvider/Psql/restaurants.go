package Psql

import "GoResturantAPI/models"

func (psql *PSqlDB)CreateRestaurant(restaurant *models.Restaurant)error{
	return psql.Create(restaurant).Error
}
