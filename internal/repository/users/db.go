package users

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

type UsersRepository interface {
	ValidApiKey(apiKey null.String) (null.Bool, error)
}

type usersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) UsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) ValidApiKey(apiKey null.String) (null.Bool, error) {

	valid := null.BoolFrom(false)

	err := r.db.Get(&valid, qryValidApiKey, apiKey)

	return valid, err
}
