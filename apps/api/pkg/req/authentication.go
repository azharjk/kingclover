package req

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"kingclover.com/api/pkg/config"
	"kingclover.com/api/pkg/database"
	"kingclover.com/api/pkg/model"
	"kingclover.com/api/pkg/model/token"
)

func IsExpires(accessToken *model.AccessToken) error {
	if time.Until(accessToken.Expires) > 0 {
		return nil
	}

	err := fmt.Errorf("%s token expires", *accessToken.Content)

	result := database.DB.Delete(&model.AccessToken{}, &model.AccessToken{ID: accessToken.ID})
	if result.Error != nil {
		return err
	}

	if result.RowsAffected <= 0 {
		return err
	}

	return err
}

func IsAuthenticated(c *gin.Context) (*model.AccessToken, error) {
	plainToken, err := c.Cookie(config.CookieNameAccessToken)
	if err != nil {
		return nil, err
	}

	h := token.Hash(plainToken)

	accessToken := model.AccessToken{}
	result := database.DB.Find(&accessToken, &model.AccessToken{Content: &h})
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected <= 0 {
		return nil, fmt.Errorf("(unknown) token expires")
	}

	if err = IsExpires(&accessToken); err != nil {
		return nil, err
	}

	return &accessToken, nil
}
