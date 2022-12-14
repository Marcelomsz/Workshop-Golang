package domain

type Dentist struct {
	Id          	int     `json:"id"`
	FirstName       string  `json:"firstname" binding:"required"`
	LastName		string	`json:"lastname" binding:"required"`
	Registration    string  `json:"registration" binding:"required"`
}
