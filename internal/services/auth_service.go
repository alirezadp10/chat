package services

type AuthService struct{}

func NewAuthService() *AuthService {
    return &AuthService{}
}

func (s *AuthService) Login() {
}

func (s *AuthService) Register() {
}
