package model

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CategoryInput struct {
	Title string `json:"title"`
}

type TransactionInput struct {
	CategoryId string  `json:"categoryId"`
	Title      string  `json:"title"`
	Amount     float32 `json:"amount"`
	Currency   string  `json:"currency"`
	Type       string  `json:"type"`
}
