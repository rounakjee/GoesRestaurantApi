package DatabseProvider

import (
	"GoResturantAPI/models"
)

type Database interface {
	CreateUser(user *models.User) error
	AllUsers() []models.User
	LoginUser(user *models.User) (bool,error)
	CreateRestaurant(restaurant *models.Restaurant) error
}
