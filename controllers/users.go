package controllers

import (
	"Stachowsky/Teacher_App/models"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var u models.Auth
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var user models.User
	if err := models.CheckUserName(&user, u.UserName); err != nil {
		c.JSON(400, "Incorrect username!")
		return
	}
	if err := models.CheckUserEmail(&user, u.Email); err != nil {
		c.JSON(400, "Incorrect email!")
		return
	}
	if err := CheckHashPassword(u.Password, user.Password); !err {
		c.JSON(400, "Incorrect password!")
		return
	}

	tokens := models.Tokens{}
	tokens.AccessToken = CreateAccessToken(user.ID, user.RoleID, &user)
	tokens.RefreshToken = CreateRefreshToken(user.ID, user.RoleID, &user)
	c.JSON(200, tokens)
}

func RegisterUser(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, err.Error())
		return
	}
	if err := models.GetUserName(&u, u.UserName); err != nil {
		c.JSON(400, "This account exists in database!")
		return
	}
	u.Password, _ = GenerateHashPassword(u.Password)
	if err := models.CreateUser(&u); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, u)
}
