package controller

import (
	"net/http"
	"v1/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {}

func (UserController) Index(c *gin.Context) {
	var s service.UserService
	users, err := s.List()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func (ctrl *UserController) Create(c *gin.Context) {
	var s service.UserService
	user, err := s.Add(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (ctrl *UserController) Update(c *gin.Context) {
	var s service.UserService
	user, err := s.Update(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (ctrl *UserController) Delete(c *gin.Context) {
	var s service.UserService 
	inputId := c.Params.ByName("id")
	id, err := s.Delete(inputId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, gin.H{"Deleted id: ": id})
	}
}