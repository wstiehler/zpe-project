package createuser

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/wstiehler/zpecreateuser-api/internal/infrastructure/logger"
	roleapi "github.com/wstiehler/zpecreateuser-api/internal/infrastructure/role-api"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func NormalizeString(s string) string {
	return strings.ToLower(s)
}

func (s *Service) CreateUser(db *gorm.DB, user *UserEntity) (*UserDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	if user == nil {
		return nil, errors.New("user is nil")
	}

	existingUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		logger.Error("Error retrieving user by email", zap.Error(err))
		return nil, err
	}

	if existingUser.Email != "" {
		logger.Error("Cannot create user. Email already exists", zap.String("Email", user.Email))
		return nil, errors.New("email already exists")
	}

	user.Id = uuid.New().String()
	user.Email = NormalizeString(user.Email)

	role, err := roleapi.GetRoleByID(user.RoleId)
	if err != nil {
		logger.Error("Error retrieving role details", zap.Error(err))
		return nil, err
	}

	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		logger.Error("Error creating new user", zap.Error(err))
		return nil, err
	}

	logger.Info("User created successfully", zap.String("Name", createdUser.Email))

	userResponse := UserDTO{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role: RoleDTO{
			Role: NormalizeString(role.Role),
		},
	}

	return &userResponse, nil
}

func (s *Service) GetUserByEmail(email string, db *gorm.DB) (*UserDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		logger.Error("Error on get user by email", zap.String("error", err.Error()))
		return nil, err
	}

	userResponse := UserDTO{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	logger.Debug("Successfull on get user by email", zap.String("email", email))

	return &userResponse, nil
}
