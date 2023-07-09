package session

import (
	"financialManagement/internal/database/postgres"
	"math/rand"
	"time"
)

func generateSession() (session string) {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	maxChars := len(chars)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 32; i++ {
		session += string(chars[random.Intn(maxChars)])
	}
	return
}

func getNewSession(s *postgres.Storage) (session string) {
	session = generateSession()
	for s.ExistsSession(session) {
		session = generateSession()
	}
	return
}

func StartSession(s *postgres.Storage, login string) (session string, err error) {
	session, err = s.GetSession(login)
	if err == nil {
		return
	}
	session = getNewSession(s)
	err = s.CreateSession(login, session)
	return
}
