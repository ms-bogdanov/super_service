package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RentBook int    `json:"rent_book"`
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Author struct {
	AuthorID int    `json:"author_id"`
	Name     string `json:"author_name"`
}

type TakeBookRequest struct {
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}
