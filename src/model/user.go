package model

//User ...
type User struct {
	ID                  int
	FirstName, LastName string
	Email, Login        string
	Password            string
}

//UserInfo ...
type UserInfo struct {
	FirstName, LastName, Login, Email string
}
