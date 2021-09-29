package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)

func main() {
	dsn := "host=0.0.0.0 dbname=sample user=user password=user sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	d, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer d.Close()

	r := gin.Default()
	rg := r.Group("users")
	{
		rg.GET("", Index)
		rg.GET("/:id", Show)
		rg.POST("", Create)
		rg.PUT("/:id", Update)
		rg.DELETE("/:id", Delete)
	}
	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.Run()
}

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func Index(c *gin.Context) {
	var users []User
	err := db.Find(&users).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, users)
}

func Show(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Create(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	err := db.Create(&user).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(http.StatusOK, user)
}
func Delete(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, id)
}
