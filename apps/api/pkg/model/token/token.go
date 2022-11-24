package token

import (
	"math/rand"

	"github.com/google/uuid"
	"kingclover.com/api/pkg/database"
	"kingclover.com/api/pkg/model"
)

const accessTokenLength = 16
const refreshTokenLength = 24

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func CreateAccessToken(uid uuid.UUID) (string, error) {
	b := make([]rune, accessTokenLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	h := Hash(string(b))

	result := database.DB.Create(&model.AccessToken{Content: &h, UserID: uid})
	if result.Error != nil {
		return "", result.Error
	}

	return string(b), nil
}

func CreateRefreshToken(uid uuid.UUID) (string, error) {
	b := make([]rune, refreshTokenLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	h := Hash(string(b))

	result := database.DB.Create(&model.RefreshToken{Content: &h, UserID: uid})
	if result.Error != nil {
		return "", result.Error
	}

	return string(b), nil
}

func CreateToken(uid uuid.UUID) (string, string, error) {
	accessToken, err := CreateAccessToken(uid)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := CreateRefreshToken(uid)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}
