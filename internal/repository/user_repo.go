package repository

import (
	"GameApi/internal/model"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetByAll(ctx context.Context) ([]model.User, error) {
	var users []model.User

	err := r.db.SelectContext(ctx, &users, "select * from users")
	return users, err
}

func (r *UserRepo) GetById(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.db.GetContext(ctx, &user, "select * from users where id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, err
}
