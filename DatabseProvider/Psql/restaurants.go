package Psql

import "GoResturantAPI/models"

func (psql *PSqlDB) CreateRestaurant(restaurant *models.Restaurant) error {
	return psql.Create(restaurant).Error
}
func (psql *PSqlDB) AllRestaurants() []models.Restaurant {
	var restaurants []models.Restaurant
	psql.Preload("User").Find(&restaurants)
	return restaurants
}
