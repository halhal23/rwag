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

func init() {
	dsn := "host=0.0.0.0 dbname=sample user=user password=password sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&User{})
}

func main() {
	r := gin.Default()
	u := r.Group("users")
	{
		u.GET("", Index)
		u.GET("/:id", Show)
		u.POST("", Create)
		u.PUT("/:id", Update)
		u.DELETE("/:id", Delete)
	}
	r.Run()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func Index(c *gin.Context) {
	var u []User
	if err = db.Find(&u).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, u)
	}
}

func Show(c *gin.Context) {
	var u User
	id := c.Params.ByName("id")
	if err = db.Where("id = ?", id).First(&u).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, u)
	}
}

func Create(c *gin.Context) {
	var u User
	c.BindJSON(&u)
	if err = db.Create(&u).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, u)
	}
}

func Update(c *gin.Context) {
	var u User
	id := c.Params.ByName("id")
	if err = db.Where("id = ?", id).First(&u).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&u)
	db.Save(&u)
	c.JSON(http.StatusOK, u)
}

func Delete(c *gin.Context) {
	var u User
	id := c.Params.ByName("id")
	if err = db.Where("id = ?", id).Delete(&u).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, gin.H{"Deleted id: ": id})
	}
}