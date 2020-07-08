package servers

import (
	"GoResturantAPI/middlewares"
	"GoResturantAPI/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)



func (ser *Server) UserCreate(c *gin.Context) {
	users := models.User{}
	err := c.BindJSON(&users)
	if err != nil {
		panic(err.Error())
	}
	hash, err := HashPassword(users.Password)
	if err != nil {
		fmt.Println("There was an error creating your account")
		fmt.Println(err.Error())
	}
	users.HashPassword = hash
	err = ser.DB.CreateUser(&users)
	if err != nil {
		panic(err.Error())
	}
	c.JSON(200, gin.H{
		"User": users,
	})
}

func (ser *Server) GetAllUsers(c *gin.Context) {
	//is Valid Token
	var users []models.User
	users = ser.DB.AllUsers()
	c.JSON(200, gin.H{
		"Users": users,
	})
}
func (ser *Server) Login(c *gin.Context) {
	users := models.User{}
	err := c.BindJSON(&users)
	if err != nil {
		panic(err.Error())
	}
	bl, str := ser.IsAuthenticated(&users)
	if bl == false && str == "Unregistered phone number..." {
		c.JSON(500, gin.H{
			"Message": "No such User..",
		})
	} else if bl == true && str == "" {
		token, err := GetAuthToken()
		if err != nil {
			c.JSON(400,gin.H{
				"message":"something is wrong..",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "You are successfully logged in...",
				"token":   token,
			})
		}

	} else if bl == false && str == "" {
		c.JSON(400, gin.H{
			"message": "Incorrect Password",
		})
	}
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (ser *Server) IsAuthenticated(users *models.User) (bool, string) {
	_, err := ser.DB.LoginUser(users)
	if err != nil {
		return false, "Unregistered phone number..."
	}
	bl := CheckPasswordHash(users.Password, users.HashPassword)
	return bl, ""
}

func GetAuthToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = models.User{}.ID
	claims["exp"] = time.Now().AddDate(0,6,0).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	authToken, err := token.SignedString(middlewares.TokenSecret)
	return authToken, err
}
