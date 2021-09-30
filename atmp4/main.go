// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var (
// 	db *gorm.DB
// 	err error
// )

// type User struct {
// 	gorm.Model
// 	Name string `gorm:"default:defaultName"`
// 	Email string `gorm:"default:defaultEmail"`
// }

// type Post struct {
// 	gorm.Model
// 	Content string
// 	UserID int
// 	User User
// }

// func main() {
// 	// set db
// 	dsn := "host=0.0.0.0 dbname=sample user=user password=password sslmode=disable"
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	db.AutoMigrate(&User{}, &Post{})
// 	defer func(){
// 		d, err := db.DB()
// 		if err != nil {
// 			panic(err)
// 		}
// 		d.Close()
// 	}()

// 	r := gin.Default()
// 	rg := r.Group("/users")
// 	{
// 		rg.GET("", Index)
// 		rg.GET("/:id", Show)
// 		rg.POST("", Create)
// 		rg.PUT("/:id", Update)
// 		rg.DELETE("/:id", Delete)
// 		rg.GET("/:id/posts", PostIndex)
// 	}
// 	r.GET("", func(c *gin.Context) {
// 		c.String(http.StatusOK, "helo")
// 	})
// 	r.Run()
// }

// func Index(c *gin.Context) {
// 	var users []User
// 	if err = db.Find(&users).Error; err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, users)
// }

// func Show(c *gin.Context) {
// 	var user User
// 	id := c.Params.ByName("id")
// 	if err = db.First(&user, "id = ?", id).Error; err != nil {
// 		c.AbortWithError(http.StatusNotFound, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// func Create(c *gin.Context) {
// 	var user User
// 	c.BindJSON(&user)
// 	if err = db.Model(&user).Create(&user).Error; err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// func Update(c *gin.Context) {
// 	var user User
// 	id := c.Params.ByName("id")
// 	if err = db.First(&user, "id = ?", id).Error; err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	c.BindJSON(&user)
// 	db.Save(&user)
// 	c.JSON(http.StatusOK, user)
// }

// func Delete(c *gin.Context) {
// 	var user User
// 	id := c.Params.ByName("id")
// 	if err = db.First(&user, "id = ?", id).Delete(&user).Error; err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// func PostIndex(c *gin.Context) {
// 	var posts []Post
// 	userId := c.Params.ByName("id")
// 	if err = db.Model(&Post{}).Where("user_id = ?", userId).Find(&posts).Error; err != nil {
// 	// if err = db.Find(&posts).Error; err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, posts)
// }

// 24.15

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main(){
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	var i int
	if x >= 0 {
		n := strconv.Itoa(x)
		ns := strings.Split(n, "")
		r := strings.Join(reverseSlise(ns), "")
		i, _ = strconv.Atoi(r)
	} else {
		f := math.Abs(float64(x))
		n := strconv.Itoa(int(f))
		ns := strings.Split(n, "")
		ns = append(ns, "-")
		r := strings.Join(reverseSlise(ns), "")
		i, _ = strconv.Atoi(r)
	}
	return i
}

func reverseSlise(s []string) []string {
	reversedSlice := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		reversedSlice[i] = s[len(s)-1-i]
	}
	return reversedSlice
}