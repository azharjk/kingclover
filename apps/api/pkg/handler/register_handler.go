package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"kingclover.com/api/pkg/config"
	"kingclover.com/api/pkg/database"
	"kingclover.com/api/pkg/model"
	"kingclover.com/api/pkg/model/token"
	"kingclover.com/api/pkg/res"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=4"`
}

func RegisterHandler(c *gin.Context) {
	r := &RegisterRequest{}

	if err := c.ShouldBindJSON(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &res.ErrorJSON{Error: err.Error(), Success: 0})
		return
	}

	p, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &res.ErrorJSON{Error: err.Error(), Success: 0})
		return
	}

	result := database.DB.Create(&model.User{Email: &r.Email, Password: string(p)})
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &res.ErrorJSON{Error: result.Error.Error(), Success: 0})
		return
	}

	user := &model.User{}
	result.Scan(&user)

	accessToken, refreshToken, err := token.CreateToken(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &res.ErrorJSON{Error: err.Error(), Success: 0})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(config.CookieNameAccessToken, accessToken, 0, "/", "localhost", true, true)
	c.SetCookie(config.CookieNameRefreshToken, refreshToken, 0, "/", "localhost", true, true)

	c.JSON(http.StatusOK, &res.RegisterJSON{AccessToken: accessToken, Success: 1})
}
