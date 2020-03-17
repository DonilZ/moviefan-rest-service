package db

import (
	"fmt"
	"webApp/src/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//Config ...
type Config struct {
	ConnectString string
}

//PgDb ...
type PgDb struct {
	dbConn *sqlx.DB

	sqlSelectUsers *sqlx.Stmt
	sqlSelectFilms *sqlx.Stmt
}

//InitDb ...
func InitDb(cfg Config) (*PgDb, error) {
	dbConn, err := sqlx.Connect("postgres", cfg.ConnectString)

	if err != nil {
		return nil, err
	}

	p := &PgDb{dbConn: dbConn}

	if err := p.dbConn.Ping(); err != nil {
		return nil, err
	}
	if err := p.createTablesIfNotExist(); err != nil {
		return nil, err
	}
	if err := p.prepareSQLStatements(); err != nil {
		return nil, err
	}

	return p, nil
}

//SelectAllUsers ...
func (p *PgDb) SelectAllUsers() ([]*model.User, error) {
	users := make([]*model.User, 0)

	if err := p.sqlSelectUsers.Select(&users); err != nil {
		return nil, err
	}

	return users, nil
}

//SelectAllFilms ...
func (p *PgDb) SelectAllFilms() ([]*model.Film, error) {
	films := make([]*model.Film, 0)

	if err := p.sqlSelectFilms.Select(&films); err != nil {
		return nil, err
	}

	return films, nil
}

//SelectAllUserFilms ...
func (p *PgDb) SelectAllUserFilms(userID int) ([]*model.Film, error) {
	films := make([]*model.Film, 0)

	queryStr := fmt.Sprintf("SELECT * FROM films WHERE films.id IN "+
		"(SELECT film_id FROM user_films WHERE user_id=%d)", userID)

	rows, _ := p.dbConn.Query(queryStr)

	for rows.Next() {
		var (
			ID   int
			year int
		)
		var name string

		rows.Scan(&ID, &name, &year)
		films = append(films, &model.Film{ID: ID, Name: name, Year: year})
	}

	return films, nil
}

//AddUser ...
func (p *PgDb) AddUser(newUser *model.User) error {
	tx := p.dbConn.MustBegin()
	_, err := tx.Exec(
		"INSERT INTO users (firstname, lastname, email, login, password) VALUES ($1, $2, $3, $4, $5)",
		newUser.FirstName, newUser.LastName, newUser.Email, newUser.Login, newUser.Password)
	tx.Commit()
	return err
}

//AddFilm ...
func (p *PgDb) AddFilm(newFilm *model.Film) error {
	tx := p.dbConn.MustBegin()

	_, err := tx.Exec(
		"INSERT INTO films (name, year) VALUES ($1, $2)",
		newFilm.Name, newFilm.Year)
	tx.Commit()

	return err
}

//AddUserFilm ...
func (p *PgDb) AddUserFilm(userID int, filmID int) error {
	tx := p.dbConn.MustBegin()
	_, err := tx.Exec(
		"INSERT INTO user_films (user_id, film_id) VALUES ($1, $2)",
		userID, filmID)
	tx.Commit()
	return err
}

//DeleteUserFilm ...
func (p *PgDb) DeleteUserFilm(userID, filmID int) error {
	tx := p.dbConn.MustBegin()
	queryStr := fmt.Sprintf("DELETE FROM user_films WHERE user_id=%d AND film_id=%d", userID, filmID)
	_, err := tx.Exec(queryStr)
	tx.Commit()
	return err
}

//GetUserByLogin ...
func (p *PgDb) GetUserByLogin(userLogin string) (*model.User, error) {

	queryStr := fmt.Sprintf("SELECT * FROM users WHERE login='%s'", userLogin)
	rows, err := p.dbConn.Query(queryStr)

	if err != nil {
		return nil, err
	}

	var desiredUser model.User

	for rows.Next() {
		var ID int
		var firstName, lastName string
		var email, login, password string

		err = rows.Scan(&ID, &firstName, &lastName, &email, &login, &password)

		if err != nil {
			return nil, err
		}

		desiredUser.ID = ID
		desiredUser.FirstName = firstName
		desiredUser.LastName = lastName
		desiredUser.Email = email
		desiredUser.Login = login
		desiredUser.Password = password
	}

	return &desiredUser, nil
}

//GetFilmByID ...
func (p *PgDb) GetFilmByID(filmID int) (*model.Film, error) {

	queryStr := fmt.Sprintf("SELECT * FROM films WHERE id=%d", filmID)
	rows, err := p.dbConn.Query(queryStr)

	if err != nil {
		return nil, err
	}

	var desiredFilm model.Film

	for rows.Next() {
		var ID, year int
		var name string

		err = rows.Scan(&ID, &name, &year)

		if err != nil {
			return nil, err
		}

		desiredFilm.ID = ID
		desiredFilm.Name = name
		desiredFilm.Year = year
	}

	return &desiredFilm, nil
}

//GetFilmIDByNameAndYear ...
func (p *PgDb) GetFilmIDByNameAndYear(name *string, year int) (int, error) {
	queryStr := fmt.Sprintf("SELECT id FROM films WHERE name='%s' AND year=%d", *name, year)
	rows, err := p.dbConn.Query(queryStr)

	if err != nil {
		return -1, err
	}

	desiredFilmID := -1

	for rows.Next() {
		err = rows.Scan(&desiredFilmID)
		if err != nil {
			return -1, err
		}
	}

	return desiredFilmID, nil
}

func (p *PgDb) createTablesIfNotExist() error {
	createUsersSQL := `

       CREATE TABLE IF NOT EXISTS users (
		id SERIAL NOT NULL PRIMARY KEY,
		firstname TEXT NOT NULL,
		lastname TEXT,
		email TEXT NOT NULL UNIQUE,
		login TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	   );

	`
	createFilmsSQL := `

       CREATE TABLE IF NOT EXISTS films (
       id SERIAL NOT NULL PRIMARY KEY,
       name TEXT NOT NULL,
       year INTEGER NOT NULL);

	`
	rowsUsersTable, err := p.dbConn.Query(createUsersSQL)
	if err != nil {
		return err
	}
	rowsUsersTable.Close()

	rowsFilmsTable, err := p.dbConn.Query(createFilmsSQL)

	if err != nil {
		return err
	}
	rowsFilmsTable.Close()

	return nil
}

func (p *PgDb) prepareSQLStatements() (err error) {

	if p.sqlSelectUsers, err = p.dbConn.Preparex(
		"SELECT Id, FirstName, LastName, Login, Email, Password FROM users",
	); err != nil {
		return err
	}

	if p.sqlSelectFilms, err = p.dbConn.Preparex(
		"SELECT Id, Name, Year FROM films",
	); err != nil {
		return err
	}

	return nil
}
