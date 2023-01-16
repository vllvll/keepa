package repositories

import (
	"database/sql"
	"log"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"

	"github.com/vllvll/keepa/internal/types"
)

type User struct {
	db *sql.DB
}

type UserInterface interface {
	IsExists(login string) bool
	CreateUser(login string, password string) (id int64, err error)
	GetUserHashByLogin(login string) (user types.User, err error)
	GetUserByID(userID int64) (user types.User, err error)
}

func NewUserRepository(db *sql.DB) UserInterface {
	return &User{db: db}
}

func (u *User) IsExists(login string) (isExist bool) {
	var count int
	err := u.db.QueryRow("SELECT 1 FROM users WHERE login = $1 LIMIT 1", login).Scan(&count)

	if err, ok := err.(*pq.Error); ok {
		if pgerrcode.IsConnectionException(string(err.Code)) {
			log.Fatalf("Error with database: %v", err)
		}

		return false
	}

	return true
}

func (u *User) CreateUser(login string, password string) (id int64, err error) {
	err = u.db.QueryRow(
		"INSERT INTO users (login, password_hash, created_at) VALUES ($1, $2, $3) RETURNING id;",
		login,
		password,
		time.Now(),
	).Scan(&id)

	if err, ok := err.(*pq.Error); ok {
		if pgerrcode.IsConnectionException(string(err.Code)) {
			log.Fatalf("Error with database: %v", err)
		}

		return 0, err
	}

	return id, nil
}

func (u *User) GetUserHashByLogin(login string) (user types.User, err error) {
	err = u.db.QueryRow("SELECT id, login, password_hash FROM users WHERE login = $1 LIMIT 1", login).Scan(&user.ID, &user.Login, &user.Hash)

	if err, ok := err.(*pq.Error); ok {
		if pgerrcode.IsConnectionException(string(err.Code)) {
			log.Fatalf("Error with database: %v", err)
		}

		return types.User{}, err
	}

	return user, nil
}

func (u *User) GetUserByID(userID int64) (user types.User, err error) {
	err = u.db.QueryRow("SELECT id, login, password_hash FROM users WHERE id = $1 LIMIT 1", userID).Scan(&user.ID, &user.Login, &user.Hash)

	if err, ok := err.(*pq.Error); ok {
		if pgerrcode.IsConnectionException(string(err.Code)) {
			log.Fatalf("Error with database: %v", err)
		}

		return types.User{}, err
	}

	return user, nil
}
