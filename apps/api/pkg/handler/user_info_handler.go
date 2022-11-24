package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kingclover.com/api/pkg/database"
	"kingclover.com/api/pkg/model"
	"kingclover.com/api/pkg/req"
	"kingclover.com/api/pkg/res"
)

func UserInfoHandler(c *gin.Context) {
	accessToken, err := req.IsAuthenticated(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &res.ErrorJSON{Error: err.Error(), Success: 0})
		return
	}

	user := model.User{}
	result := database.DB.Find(&user, &model.User{ID: accessToken.UserID})
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &res.ErrorJSON{Error: result.Error.Error(), Success: 0})
		return
	}

	if result.RowsAffected <= 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &res.ErrorJSON{Error: result.Error.Error(), Success: 0})
		return
	}

	c.JSON(http.StatusOK, &res.UserInfoJSON{Data: res.UserJSON{
		ID:    user.ID.String(),
		Email: *user.Email,
	}, Success: 1})
}
