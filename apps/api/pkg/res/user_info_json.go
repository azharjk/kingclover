package res

type UserJSON struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserInfoJSON struct {
	Data    UserJSON `json:"data"`
	Success uint8    `json:"success"`
}
