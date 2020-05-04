package models

type User struct {
	Name		string	`json:"name"`
	Username	string	`json:"username"`
	Email		string	`json:"email"`
	Id			int		`json:"id"`
	Pass		string	`json:"pass"`
}

type JSONErr struct {
	Error string `json:"error"`
}

type JSONSuccess struct {
	Success string `json:"success"`
}
