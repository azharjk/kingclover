package res

type TokenJSON struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Success      uint8  `json:"success"`
}
