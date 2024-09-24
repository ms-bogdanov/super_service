package repository

import (
	"context"
	"database/sql"
	"fmt"
	"super_service/config"
	"super_service/internal/model"
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
	queryTakeBook   = `insert into register(user_id, book_id) values ($1, $2)`
	queryReturnBook = `update register set return_at = now() where user_id = $1 and book_id = $2
		and return_at is null`
	querySearchUser       = `select true from users where id = $1`
	querySearchBook       = `select true from books where id = $1`
	queryCheckAuthor      = `select count(*) from authors`
	queryAddAuthor        = `insert into authors ("name") values ($1)`
	queryAddBook          = `insert into books ("title", author_id) values ($1, $2)`
	queryAddUsers         = `insert into users ("name") values ($1)`
	queryGetUsersAndBooks = `select u.name, b.title, a.name from users u
    	left join register r on r.user_id = u.id
    	left join books b on r.book_id = b.id
    	left join authors a on b.author_id = a.id`
	queryCheckBook  = `select count(*) from books`
	queryCheckUsers = `select count(*) from users`
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

func (us UserStorage) ReturnBook(uid, bid int) error {
	_, err := us.db.ExecContext(context.Background(), queryReturnBook, uid, bid)
	if err != nil {
		return err
	}

	return nil
}

func (us UserStorage) SearchUser(uid int) error {
	var ok bool

	err := us.db.QueryRowContext(context.Background(), querySearchUser, uid).Scan(&ok)
	if err != nil {
		return err
	}

	return nil
}

func (us UserStorage) SearchBook(bid int) error {
	var ok bool

	err := us.db.QueryRowContext(context.Background(), querySearchBook, bid).Scan(&ok)
	if err != nil {
		return err
	}

	return nil
}

func (us UserStorage) GetAuthorCount() int {
	var count int
	err := us.db.QueryRowContext(context.Background(), queryCheckAuthor).Scan(&count)
	if err != nil {
		return 0
	}

	return count
}

func (us UserStorage) GetBookCount() int {
	var count int
	err := us.db.QueryRowContext(context.Background(), queryCheckBook).Scan(&count)
	if err != nil {
		return 0
	}

	return count
}

func (us UserStorage) GetUsersCount() int {
	var count int
	err := us.db.QueryRowContext(context.Background(), queryCheckUsers).Scan(&count)
	if err != nil {
		return 0
	}

	return count
}

func (us UserStorage) AddAuthor(v string) error {
	_, err := us.db.ExecContext(context.Background(), queryAddAuthor, v)
	if err != nil {
		return err
	}

	return nil
}

func (us UserStorage) AddBook(title string, author int) error {
	_, err := us.db.ExecContext(context.Background(), queryAddBook, title, author)
	if err != nil {
		return err
	}

	return nil
}

func (us UserStorage) AddUser(v string) error {
	_, err := us.db.ExecContext(context.Background(), queryAddUsers, v)
	if err != nil {
		return err
	}

	return nil
}

func (us UserStorage) GetUsersAndBooks() map[string][]model.Book {
	rows, err := us.db.QueryContext(context.Background(), queryGetUsersAndBooks)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var userAndBooks = make(map[string][]model.Book)
	var name string
	var book model.Book

	for rows.Next() {
		err = rows.Scan(&name, &book.Title, &book.Author)
		if err != nil {
			fmt.Println(err)
		}
		userAndBooks[name] = append(userAndBooks[name], book)
	}
	return userAndBooks
}
