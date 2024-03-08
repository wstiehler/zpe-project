package createuser

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wstiehler/zpecreateuser-api/internal/infrastructure/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
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

	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		logger.Error("Error creating new user", zap.Error(err))
		return nil, err
	}

	logger.Info("User created successfully", zap.String("Name", createdUser.Email))

	userResponse := UserDTO{
		Id:    createdUser.Id,
		Name:  createdUser.Name,
		Email: createdUser.Email,
		Role:  createdUser.Role,
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
		Role:  user.Role,
	}

	logger.Debug("Successfull on get user by email", zap.String("email", email))

	return &userResponse, nil
}
