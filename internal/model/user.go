package model

type User struct {
	ID          string `json:"id"`
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	FathersName string `json:"fathers_name"`
	GroupNumber string `json:"group_number"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Debt        Debt   `json:"debt"`
}

type Token struct {
	Token string `json:"token"`
}
