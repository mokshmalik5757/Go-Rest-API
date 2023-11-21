package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mokshmalik5757/Go-Rest-API/config"
	"github.com/mokshmalik5757/Go-Rest-API/models"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, &users)
}

func CreateUser (c *gin.Context) {
	user := models.User{}
	config.DB.Create(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(http.StatusCreated, &user)
}
func DeleteUser (c *gin.Context) {
	var id string = c.Param("id")
	user := models.User{}
	intNum, err := strconv.ParseInt(id, 8,8)
	config.DB.Where("id = ?", intNum).Delete(&user)
	if err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted succesfully"})
}
func UpdateUser (c *gin.Context) {
	var id string = c.Param("id")
	idNum, err := strconv.ParseInt(id, 8,8)
	user := models.User{}

	if err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
	}

	config.DB.Where("id = ?", idNum).First(&user)
	c.Bind(&user)
	config.DB.Save(&user)
	c.JSON(http.StatusOK, &user)

}
