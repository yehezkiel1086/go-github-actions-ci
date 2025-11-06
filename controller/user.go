package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-github-actions-ci/model"
	"github.com/yehezkiel1086/go-github-actions-ci/storage/postgres"
	"github.com/yehezkiel1086/go-github-actions-ci/util"
)

type UserController struct {
	db *postgres.DB
}

func InitUserController(db *postgres.DB) *UserController {
	return &UserController{
		db: db,
	}
}

type UserReq struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (uc *UserController) Register(c *gin.Context) {
	// bind input	
	var req UserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash password
	hashed, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Password = hashed

	// create new user
	db := uc.db.GetDB()
	if err := db.Create(&model.User{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
