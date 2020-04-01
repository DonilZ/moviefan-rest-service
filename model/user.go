package model

//User ...
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Login     string
	Password  string
}

//UserInfo ...
type UserInfo struct {
	FirstName string
	LastName  string
	Login     string
	Email     string
}
