package users

import (
	"crypto/sha256"
	"fmt"

	"github.com/jmoiron/sqlx"
)

//Store is visible outside of package for easy access from othe rparts of code
/*
	The idea here is to use the package in easy way.
	1) When app starts run the migrations etc
	2) Initialize models (by running Setup() )
	3) Use package like this
	```
		user, err := users.Store.FindByEmail("test@test.com")
		if ...
	```
*/
var Store UserStoreInterface
var db *sqlx.DB

//UserStoreInterface interface that enables us to mock in tests database layer.
type UserStoreInterface interface {
	FindByEmail(email string) (*User, error)
	CreateUser(email string, password string) (*User, error)
}

//UserStore is wrapping connection and providing interface for the page level access
type UserStore struct {
	db *sqlx.DB
}

//Setup
func Setup(connection *sqlx.DB) {
	Store = &UserStore{db: connection}
}

//CreateUser creates user into db or returns error
func (store *UserStore) CreateUser(email string, password string) (*User, error) {
	encryptedPassword := HashPassword(password)

	user := &User{Email: email, Password: encryptedPassword}
	tx := store.db.MustBegin()
	_, err := tx.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return store.FindByEmail(email)
}

//FindByEmail fetches user by email
func (store *UserStore) FindByEmail(email string) (*User, error) {
	user := &User{}

	err := store.db.Get(user, "SELECT * FROM users WHERE email=$1", email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

//HashPassword Needs to go to their own package/utils or auth related shit
func HashPassword(password string) string {
	bytes := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", bytes)
}
