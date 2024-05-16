package models

type Token struct {
	token string
}

func (s *Token) Set(token string) {
	s.token = token
}

func (s *Token) Get() string {
	return s.token
}
