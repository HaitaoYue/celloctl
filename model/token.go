package model

type UserInfo struct {
	ID string `json:"pk"`
	Name string `json:"username"`
	Email string `json:"email"`
}

type TokenResponse struct {
	Token string `json:"token"`
	User UserInfo `json:"user"`
}
