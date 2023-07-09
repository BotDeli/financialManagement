package postgres

func (s *Storage) ExistsSession(session string) bool {
	query := "SELECT login FROM sessions WHERE key = $1"
	var login string
	err := s.db.QueryRow(query, session).Scan(&session)
	return err == nil && login != ""
}

func (s *Storage) GetSession(login string) (session string, err error) {
	query := "SELECT key FROM sessions WHERE login = $1"
	err = s.db.QueryRow(query, login).Scan(&session)
	return
}

func (s *Storage) CreateSession(login, session string) error {
	query := "INSERT INTO sessions (login, key) VALUES ($1, $2)"
	_, err := s.db.Exec(query, login, session)
	return err
}
