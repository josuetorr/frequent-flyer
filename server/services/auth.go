package services

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Signup(email string, password string) error {
	return nil
}
