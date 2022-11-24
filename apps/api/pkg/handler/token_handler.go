package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"kingclover.com/api/pkg/config"
	"kingclover.com/api/pkg/database"
	"kingclover.com/api/pkg/model"
	"kingclover.com/api/pkg/model/token"
	"kingclover.com/api/pkg/res"
)

func TokenHandler(c *gin.Context) {
	plainToken, err := c.Cookie(config.CookieNameRefreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &res.ErrorJSON{Error: err.Error(), Success: 0})
		return
	}

	h := token.Hash(plainToken)

	refreshTokenModel := model.RefreshToken{}
	result := database.DB.Find(&refreshTokenModel, &model.RefreshToken{Content: &h})
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &res.ErrorJSON{Error: result.Error.Error(), Success: 0})
		return
	}

	if result.RowsAffected <= 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &res.ErrorJSON{Error: fmt.Sprintf("%s is not found", plainToken), Success: 0})
		return
	}

	user := model.User{}
	result = database.DB.Find(&user, &model.User{ID: refreshTokenModel.UserID})
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &res.ErrorJSON{Error: result.Error.Error(), Success: 0})
		return
	}

	if result.RowsAffected <= 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &res.ErrorJSON{Error: fmt.Sprintf("%s user id is not found", refreshTokenModel.UserID), Success: 0})
		return
	}

	accessToken, refreshToken, err := token.CreateToken(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &res.ErrorJSON{Error: err.Error(), Success: 0})
		return
	}

	result = database.DB.Delete(&model.RefreshToken{}, &model.RefreshToken{ID: refreshTokenModel.ID})
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &res.ErrorJSON{Error: result.Error.Error(), Success: 0})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(config.CookieNameAccessToken, accessToken, 0, "/", "localhost", true, true)
	c.SetCookie(config.CookieNameRefreshToken, refreshToken, 0, "/", "localhost", true, true)

	c.JSON(http.StatusOK, &res.TokenJSON{AccessToken: accessToken, RefreshToken: refreshToken, Success: 1})
}
