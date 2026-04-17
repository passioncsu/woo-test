package service

import (
	"errors"

	"student-admin/backend/internal/config"
	"student-admin/backend/internal/model"
	"student-admin/backend/internal/repository"
	"student-admin/backend/pkg/jwt"
)

type AuthService struct {
	repo *repository.AdminRepository
	cfg  *config.Config
}

func NewAuthService(repo *repository.AdminRepository, cfg *config.Config) *AuthService {
	return &AuthService{repo: repo, cfg: cfg}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	Username    string `json:"username"`
}

func (s *AuthService) Login(req LoginRequest) (*LoginResponse, error) {
	admin, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if !repository.CheckPassword(req.Password, admin.Password) {
		return nil, errors.New("invalid username or password")
	}

	token, err := jwt.GenerateToken(s.cfg.JWT.Secret, admin.ID, admin.Username, s.cfg.JWT.Expiration)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &LoginResponse{
		AccessToken: token,
		Username:    admin.Username,
	}, nil
}

func (s *AuthService) Register(req LoginRequest) error {
	_, err := s.repo.FindByUsername(req.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := repository.HashPassword(req.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}

	admin := &model.Admin{
		Username: req.Username,
		Password: hashedPassword,
	}

	return s.repo.Create(admin)
}
