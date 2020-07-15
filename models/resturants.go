package models

import "github.com/jinzhu/gorm"

type Restaurant struct {
	gorm.Model
	Name            string `json:"name"`
	NoOfTwoSeaters  int    `json:"noOfTwoSeaters"`
	NoOfFourSeaters int    `json:"noOfFourSeaters"`
	NoOFSixSeaters  int    `json:"noOfSixSeaters"`
	Latitude        string `json:"latitude"`
	Longitude       string `json:"longitude"`
	UserID          uint   `json:"ownerId"`
	User            User   `gorm:"foreignkey:user_id;association_foreignkey:id"`
	OpeningTime     string `json:"openingTime"`
	ClosingTime     string `json:"closingTime"`
}
type Menu struct {
	gorm.Model
	Restaurant   Restaurant `gorm:"foreignkey:restaurant_id;association_foreignkey:id"`
	RestaurantID uint       `json:"restaurantsId"`
	FoodName     string     `json:"foodName"`
	Price        int        `json:"price"`
}
