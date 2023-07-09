package tests

import (
	"financialManagement/internal/session"
	"testing"
)

func TestStartSession(t *testing.T) {
	sessia, err := session.StartSession(STORAGE, TestLogin)
	if err != nil {
		t.Error(err)
	}
	if sessia == "" {
		sessionIsEmpty(t, sessia)
	}
}
