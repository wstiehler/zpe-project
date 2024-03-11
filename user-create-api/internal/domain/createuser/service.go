package createuser

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/wstiehler/zpecreateuser-api/internal/infrastructure/logger"
	roleapi "github.com/wstiehler/zpecreateuser-api/internal/infrastructure/role-api"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	repo Repository
}

func NewService(repo *Repository) *Service {
	return &Service{*repo}
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

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		logger.Error("Error on encript password")
		return nil, errors.New("error on encript password")
	}

	user.Id = uuid.New().String()
	user.Email = NormalizeString(user.Email)
	user.Password = string(hash)

	role, err := roleapi.GetRoleByID(user.RoleId)
	if err != nil {
		logger.Error("Error retrieving role details", zap.Error(err))
		return nil, errors.New("error on communication api")
	}

	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		logger.Error("Error creating new user", zap.Error(err))
		return nil, errors.New("error on creating new user")
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
