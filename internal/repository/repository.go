package repository

import (
	"database/sql"
	"errors"
	"github.com/Zhenya671/golang-test-task/internal/model"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type IUserRepository interface {
	SignIn(logIn model.User) (model.User, error)
	SignUp(user model.User) (model.User, error)
}

type Repository struct {
	DB *sql.DB
}

func NewRepository(pgConf string) (*Repository, error) {
	db, err := sql.Open("pgx", pgConf)
	if err != nil {
		return nil, err
	}

	return &Repository{
		DB: db,
	}, nil
}

func (r *Repository) SignIn(logIn model.User) (model.User, error) {
	query := `select id, login from users where login=$1 and password=$2`
	var savedUser model.User

	row := r.DB.QueryRow(query, &logIn.Login, &logIn.Password)

	if err := row.Scan(&savedUser.ID, &savedUser.Login); err != nil {
		return savedUser, err
	}

	if savedUser.ID == "" {
		return savedUser, errors.New("")
	}

	return savedUser, nil
}

func (r *Repository) SignUp(user model.User) (model.User, error) {
	query := `insert into users(lastname, firstname, fathersname, group_number, login, password) values ($1,$2) returning id, login`

	var savedUser model.User
	row := r.DB.QueryRow(query, user.LastName, user.FirstName, user.FathersName, user.GroupNumber, user.Login, user.Password)

	if err := row.Scan(&savedUser.ID, &savedUser.Password); err != nil {
		if err == sql.ErrNoRows {
			return savedUser, err
		}
		return savedUser, err
	}
	return savedUser, nil
}
