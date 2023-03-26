package repository

import "github.com/jmoiron/sqlx"

type userRepositoryDB struct {
	db *sqlx.DB
}

func NewUserRepositoryDB(db *sqlx.DB) UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) Insert(user User) (*User, error) {
	query := "insert into users (username, password) values (?, ?)"
	result, err := r.db.Exec(
		query,
		user.Username,
		user.Password,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.UserID = int(id)

	return &user, nil
}

func (r userRepositoryDB) Get(username string) (*User, error) {

	user := User{}
	query := "select * from users where username=?"
	err := r.db.Get(&user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
