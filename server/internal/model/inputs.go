package model

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CategoryInput struct {
	Title string `json:"title"`
}
