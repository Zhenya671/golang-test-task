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
	PayOff(userId string, input model.Debt) (model.Debt, error)
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
	query := `insert into users(lastname, firstname, fathersname, group_id, login, password) values ($1, $2, $3, (SELECT id FROM groups WHERE group_number = $4), $5, $6) returning id, login`

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

func (r *Repository) PayOff(userId string, input model.Debt) (model.Debt, error) {
	var debt model.Debt

	tx, err := r.DB.Begin()
	if err != nil {
		return debt, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var userDebt float64
	err = tx.QueryRow(`
	SELECT amount 
	FROM debt
	WHERE user_id = $1
	FOR UPDATE`, userId).Scan(&userDebt)
	if err != nil {
		return debt, err
	}

	amount := userDebt - input.Amount
	if amount <= 0.00 {
		debt.CashBack = amount * -1
		amount = 0
	}

	row := tx.QueryRow(`
	UPDATE debt 
	SET amount = $1
	WHERE user_id = $2
	RETURNING id, amount`, amount, userId)
	err = row.Scan(&debt.ID, &debt.Amount)
	if err != nil {
		return debt, err
	}
	return debt, nil
}
