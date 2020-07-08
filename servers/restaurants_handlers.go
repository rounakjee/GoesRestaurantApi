package servers

import (
	"GoResturantAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ser *Server)restaurantCreate(c *gin.Context){
	restaurant := models.Restaurant{}
	err := c.BindJSON(&restaurant)
	if err != nil{
		panic(err.Error())
	}
	err = ser.DB.CreateRestaurant(&restaurant)
	if err!= nil{
		panic(err.Error())
	}
	c.JSON(http.StatusOK,gin.H{
		"Restaurant":&restaurant,
	})
}
