package repository

import (
	"context"
	"database/sql"
	"fmt"
	"super_service/config"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(cfg config.PgConfig) *UserStorage {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil
	}

	return &UserStorage{
		db: db,
	}
}

const (
	queryTakeBook   = `update books set user_id = $1 where book_id = $2 and user_id is null`
	querySearchUser = `select true from users where id = $1`
)

type UserRepository interface {
	TakeBook(uid, bid int) error
}

func (us UserStorage) TakeBook(uid, bid int) error {
	_, err := us.db.ExecContext(context.Background(), queryTakeBook, uid, bid)
	if err != nil {
		return err
	}

	return nil
}

func (us UserStorage) SearchUser(uid int) error {
	err := us.db.QueryRowContext(context.Background(), querySearchUser, uid).Scan(&uid, &uid)
	if err != nil {
		return err
	}

	return nil
}

func (us UserStorage) SearchBook(bid int) error {
	err := us.db.QueryRowContext(context.Background(), queryTakeBook, bid).Scan(&bid)
	if err != nil {
		return err
	}

	return nil
}
