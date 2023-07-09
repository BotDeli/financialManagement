package postgres

import (
	"errors"
	"financialManagement/pkg/errFormat"
	"financialManagement/pkg/hashing"
)

type User struct {
	Login    string
	Password string
}

const (
	location = "internal.database.postgres.users"
)

var (
	ErrInvalidUser           = errors.New("invalid login or password")
	ErrUserAlreadyRegistered = errors.New("user already registered")
	ErrUserDontRegistered    = errors.New("user not registered")
	ErrInvalidPassword       = errors.New("invalid password")
)

// REGISTRATION

func (s *Storage) RegisterUser(u User) error {
	if dontCorrectUser(u) {
		return ErrInvalidUser
	}
	if s.ExistsUser(u.Login) {
		return ErrUserAlreadyRegistered
	}
	u.Password = hashing.GetHash(u.Password)
	err := insertUser(s, u)
	return err
}

func dontCorrectUser(u User) bool {
	return len(u.Login) < 6 || len(u.Password) < 6
}

func insertUser(s *Storage, u User) error {
	query := "INSERT INTO users(login, password) VALUES ($1, $2)"
	_, err := s.db.Exec(query, u.Login, u.Password)
	return err
}

// EXISTS

func (s *Storage) ExistsUser(login string) bool {
	u, err := findUserByLogin(s, login)
	return err == nil && u.Login != "" && u.Password != ""
}

func findUserByLogin(s *Storage, login string) (u User, err error) {
	query := "SELECT * FROM users WHERE login = $1"
	err = s.db.QueryRow(query, login).Scan(&u.Login, &u.Password)
	return
}

// REMOVE

func (s *Storage) RemoveUser(login string) error {
	query := "DELETE FROM users WHERE login = $1"
	_, err := s.db.Exec(query, login)
	return err
}

// LOGIN

func (s *Storage) LoginUser(u User) (session string, err error) {
	correct, err := correctUserAuthorization(s, u)
	if correct {
		session = "successful"
	}
	return
}

func correctUserAuthorization(s *Storage, u User) (correct bool, err error) {
	if dontCorrectUser(u) {
		err = ErrInvalidUser
	}
	if !s.ExistsUser(u.Login) {
		err = ErrUserDontRegistered
		return
	}
	alreadyHashPassword := hashing.GetHash(u.Password)
	successfulHashPassword, err := s.getHashPassword(u.Login)
	if err != nil {
		return
	}
	if alreadyHashPassword != successfulHashPassword {
		err = ErrInvalidPassword
		return
	}
	return true, nil
}

func (s *Storage) getHashPassword(login string) (string, error) {
	u, err := findUserByLogin(s, login)
	if err != nil {
		err = errFormat.FormatError(location, "getHashPassword", err)
	}
	return u.Password, err
}

func (u *User) getSession() string {
	return hashing.GetHash(u.Login + u.Password)
}
