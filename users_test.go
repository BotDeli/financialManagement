package tests

import (
	"financialManagement/internal/database/postgres"
	"testing"
)

const (
	ErrDontCorrectHandlingError = "dont correct handling error"
	ErrSessionEmpty             = "Session cant be empty"

	TestLogin    = "<LOGIN>"
	TestPassword = "<PASSWORD>"
)

var (
	TestUser = postgres.User{
		Login:    TestLogin,
		Password: TestPassword,
	}
)

// EXISTS FALSE USER

func TestFalseExistsUser(t *testing.T) {
	ok := STORAGE.ExistsUser(TestLogin)
	if ok {
		t.Error("Expected false, got true")
	}
}

// REGISTER USER

func TestRegisterEmptyUser(t *testing.T) {
	u := postgres.User{}
	err := STORAGE.RegisterUser(u)
	errIsNilDontCorrectHandling(t, err)
}

func errIsNilDontCorrectHandling(t *testing.T, err error) {
	if err == nil {
		t.Error(ErrDontCorrectHandlingError)
	}
}

func TestRegisterEmptyLogin(t *testing.T) {
	u := postgres.User{Password: TestPassword}
	err := STORAGE.RegisterUser(u)
	errIsNilDontCorrectHandling(t, err)
}

func TestRegisterEmptyPassword(t *testing.T) {
	u := postgres.User{Login: TestLogin}
	err := STORAGE.RegisterUser(u)
	errIsNilDontCorrectHandling(t, err)
}

func TestRegisterFiveCharsLogin(t *testing.T) {
	u := postgres.User{Login: "<abc>", Password: TestPassword}
	err := STORAGE.RegisterUser(u)
	errIsNilDontCorrectHandling(t, err)
}

func TestRegisterFiveCharsPassword(t *testing.T) {
	u := postgres.User{Login: TestLogin, Password: "<def>"}
	err := STORAGE.RegisterUser(u)
	errIsNilDontCorrectHandling(t, err)
}

func TestTrueRegisterUser(t *testing.T) {
	err := STORAGE.RegisterUser(TestUser)
	if err != nil {
		t.Error(err)
	}
}

func TestFalseRegisterUser(t *testing.T) {
	err := STORAGE.RegisterUser(TestUser)
	if err == nil {
		t.Error(ErrDontCorrectHandlingError, "| user already exists")
	}
}

// TRUE EXISTS USER

func TestTrueExistsUser(t *testing.T) {
	ok := STORAGE.ExistsUser(TestLogin)
	if !ok {
		t.Error("Expected true, got false")
	}
}

func TestRemoveUser(t *testing.T) {
	err := STORAGE.RemoveUser(TestLogin)
	if err != nil {
		t.Error(err)
	}
}

// LOGIN

func TestLoginEmptyUser(t *testing.T) {
	session, err := STORAGE.LoginUser(postgres.User{
		Login:    "",
		Password: "",
	})
	errIsNilDontCorrectHandling(t, err)
	sessionIsDontEmpty(t, session)
}

func sessionIsDontEmpty(t *testing.T, session string) {
	if session != "" {
		t.Errorf("Session dont be empty, got: %s", session)
	}
}

func TestLoginDontExistsUser(t *testing.T) {
	session, err := STORAGE.LoginUser(TestUser)
	errIsNilDontCorrectHandling(t, err)
	sessionIsDontEmpty(t, session)
}

func TestLoginExistsUser(t *testing.T) {
	registerUser(t)
	session, err := STORAGE.LoginUser(TestUser)
	if err != nil {
		t.Error(err)
	}
	sessionIsEmpty(t, session)
	removeUser(t)
}

func sessionIsEmpty(t *testing.T, session string) {
	if session == "" {
		t.Error(ErrSessionEmpty)
	}
}

func registerUser(t *testing.T) {
	err := STORAGE.RegisterUser(TestUser)
	if err != nil {
		t.Error(err)
	}
}

func removeUser(t *testing.T) {
	err := STORAGE.RemoveUser(TestLogin)
	if err != nil {
		t.Error(err)
	}
}

func TestLoginDontCorrectPassword(t *testing.T) {
	registerUser(t)
	session, err := STORAGE.LoginUser(postgres.User{
		Login:    TestLogin,
		Password: "PASSWORD",
	})
	errIsNilDontCorrectHandling(t, err)
	sessionIsDontEmpty(t, session)
	removeUser(t)
}
