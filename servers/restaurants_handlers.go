package servers

import (
	"GoResturantAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ser *Server) RestaurantCreate(c *gin.Context) {
	restaurant := models.Restaurant{}
	err := c.BindJSON(&restaurant)
	if err != nil {
		panic(err.Error())
	}
	err = ser.DB.CreateRestaurant(&restaurant)
	if err != nil {
		panic(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"Restaurant": &restaurant,
	})
}
func (ser *Server) GetAllRestaurants(c *gin.Context) {
	restaurants := ser.DB.AllRestaurants()
	c.JSON(http.StatusOK, gin.H{
		"AllRestaurants": restaurants,
	})
}
