package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kingclover.com/api/pkg/config"
	"kingclover.com/api/pkg/database"
	"kingclover.com/api/pkg/model"
	"kingclover.com/api/pkg/req"
	"kingclover.com/api/pkg/res"
)

func LogoutHandler(c *gin.Context) {
	accessToken, err := req.IsAuthenticated(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &res.ErrorJSON{Error: err.Error(), Success: 0})
		return
	}

	result := database.DB.Delete(&model.AccessToken{}, &model.AccessToken{UserID: accessToken.UserID})
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &res.ErrorJSON{Error: result.Error.Error(), Success: 0})
		return
	}

	if result.RowsAffected <= 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &res.ErrorJSON{Error: "some error occured", Success: 0})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(config.CookieNameAccessToken, "loggedout", -1, "/", "localhost", true, true)

	c.JSON(http.StatusOK, &res.LogoutJSON{Success: 1})
}
