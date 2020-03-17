package model

type db interface {
	SelectAllUsers() ([]*User, error)
	AddUser(newUser *User) error
	GetUserByLogin(userLogin string) (*User, error)

	SelectAllFilms() ([]*Film, error)
	AddFilm(newFilm *Film) error
	GetFilmByID(filmID int) (*Film, error)
	GetFilmIDByNameAndYear(name *string, year int) (int, error)

	SelectAllUserFilms(userID int) ([]*Film, error)
	AddUserFilm(userID int, filmID int) error
	DeleteUserFilm(userID, filmID int) error
}

//Model ...
type Model struct {
	db
}

//New ...
func New(db db) *Model {
	return &Model{
		db: db,
	}
}

//Users ...
func (m *Model) Users() ([]*User, error) {
	return m.SelectAllUsers()
}

//Films ...
func (m *Model) Films() ([]*Film, error) {
	return m.SelectAllFilms()
}

//UserFilms ...
func (m *Model) UserFilms(userID int) ([]*Film, error) {
	return m.SelectAllUserFilms(userID)
}
