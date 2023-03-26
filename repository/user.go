package repository

type User struct {
	UserID   int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type UserRepository interface {
	Insert(User) (*User, error)
	Get(string) (*User, error)
}
