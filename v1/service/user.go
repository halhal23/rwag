package service

import (
	"v1/db"
	"v1/entity"

	"github.com/gin-gonic/gin"
)

type UserService struct {}

type User entity.User

func (s *UserService) List() ([]User, error) {
	d := db.GetDB()
	var u []User
	if err := d.Find(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (s *UserService) GetById(id string) (User, error) {
	d := db.GetDB()
	var u User
	if err := d.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err 
	}
	return u, nil
}

func (s *UserService) Add(c *gin.Context) (User, error) {
	d := db.GetDB()
	var u User
	c.BindJSON(&u)
	if err := d.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (s *UserService) Update(c *gin.Context) (User, error) {
	d := db.GetDB()
	id := c.Params.ByName("id")
	var u User
	if err := d.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	c.BindJSON(&u)
	d.Save(&u)
	return u, nil
}

func (s *UserService) Delete(id string) (string, error) {
	d := db.GetDB()
	var u User
	if err := d.Where("id = ?", id).Delete(&u).Error; err != nil {
		return id, err
	}
	return id, nil
}