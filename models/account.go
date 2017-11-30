package models

type Account struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Profile  Detail `json:"profile"`
}

type Detail struct {
	Nickname string `json:"nickname"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Email   string `json:"email"`
}
