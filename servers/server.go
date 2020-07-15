package servers

import (
	"GoResturantAPI/DatabseProvider"
	"GoResturantAPI/middlewares"
	"GoResturantAPI/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Server struct {
	DB DatabseProvider.Database
}

func NewServer(db DatabseProvider.Database) Server {
	return Server{
		db,
	}
}
func (ser *Server) Start(port string) {
	router := ser.SetupRouter()
	err := router.Run(":" + port)
	if err != nil {
		panic(err.Error())
	}
}
func (ser *Server) SetupRouter() *gin.Engine {
	router := gin.Default()
	usersGroup := router.Group("users")
	{
		usersGroup.POST("createUser", ser.UserCreate)
		usersGroup.GET("getAllUsers", middlewares.TokenAuthMiddleware(), ser.GetAllUsers)
		usersGroup.GET("getUser/:id", ser.GetUser)
		usersGroup.POST("login", ser.Login)
	}
	restaurantsGroup := router.Group("restaurants")
	{
		restaurantsGroup.POST("createRestaurant", ser.RestaurantCreate)
		restaurantsGroup.GET("getAllRestaurants", ser.GetAllRestaurants)
	}
	menuGroup := router.Group("menus")
	{
		menuGroup.POST("createMenu", ser.MenuCreate)
	}
	return router
}
func InitialMigration(db *gorm.DB) {

	db.CreateTable(&models.User{})       //,&models.Restaurant{},&models.Menu{})
	db.CreateTable(&models.Restaurant{}) //.AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.CreateTable(&models.Menu{})
	db.Model(&models.Restaurant{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Menu{}).AddForeignKey("restaurant_id", "restaurants(id)", "RESTRICT", "RESTRICT")
}
