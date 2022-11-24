package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"kingclover.com/api/pkg/config"
	"kingclover.com/api/pkg/database"
	"kingclover.com/api/pkg/model"
	"kingclover.com/api/pkg/model/token"
	"kingclover.com/api/pkg/res"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	r := &LoginRequest{}

	if err := c.ShouldBindJSON(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &res.ErrorJSON{Error: err.Error(), Success: 0})
		return
	}

	user := model.User{}
	result := database.DB.Find(&user, &model.User{Email: &r.Email})
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &res.ErrorJSON{Error: result.Error.Error(), Success: 0})
		return
	}

	if result.RowsAffected <= 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &res.ErrorJSON{Error: fmt.Sprintf("%s is not found", r.Email), Success: 0})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &res.ErrorJSON{Error: err.Error(), Success: 0})
		return
	}

	accessToken, refreshToken, err := token.CreateToken(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &res.ErrorJSON{Error: err.Error(), Success: 0})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(config.CookieNameAccessToken, accessToken, 0, "/", "localhost", true, true)
	c.SetCookie(config.CookieNameRefreshToken, refreshToken, 0, "/", "localhost", true, true)

	c.JSON(http.StatusOK, &res.LoginJSON{AccessToken: accessToken, Success: 1})
}
