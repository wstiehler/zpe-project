package detailuser

import (
	"errors"
	"strings"

	"github.com/wstiehler/zpedetailuser-api/internal/infrastructure/logger"
	roleapi "github.com/wstiehler/zpedetailuser-api/internal/infrastructure/role-api"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	repo Repository
}

var ErrUserNotFound = errors.New("user not found")

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func NormalizeString(s string) string {
	return strings.ToLower(s)
}

func (s *Service) GetUserByEmail(db *gorm.DB, email string) (UserEntity, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return UserEntity{}, err
	}
	return user, nil
}

func (s *Service) GetUserByID(db *gorm.DB, id string) (UserEntity, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return UserEntity{}, err
	}
	return user, nil
}

func (s *Service) GetUserByName(db *gorm.DB, name string) (UserEntity, error) {
	user, err := s.repo.GetUserByName(name)
	if err != nil {
		return UserEntity{}, err
	}
	return user, nil
}

func (s *Service) HandleUserNotFound(logger *zap.Logger, err error) ([]UserEntity, error) {
	if errors.Is(err, ErrUserNotFound) {
		logger.Info("User not found")
		users, err := s.repo.GetAllUsers()
		if err != nil {
			logger.Error("Error on get all users", zap.Error(err))
			return nil, err
		}
		return users, nil
	}
	return nil, err
}

func (s *Service) GetUserByCriteria(db *gorm.DB, criteria string, value string) ([]UserDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	validCriteria := map[string]bool{"email": true, "id": true, "name": true}
	if !validCriteria[criteria] {
		return nil, errors.New("invalid criteria")
	}

	var users []UserEntity

	switch criteria {
	case "email":
		user, err := s.GetUserByEmail(db, value)
		if err != nil {
			users, _ = s.HandleUserNotFound(logger, err)
		} else {
			users = append(users, user)
		}
	case "id":
		user, err := s.GetUserByID(db, value)
		if err != nil {
			users, _ = s.HandleUserNotFound(logger, err)
		} else {
			users = append(users, user)
		}
	case "name":
		user, err := s.GetUserByName(db, value)
		if err != nil {
			users, _ = s.HandleUserNotFound(logger, err)
		} else {
			users = append(users, user)
		}
	default:
		return nil, errors.New("invalid criteria")
	}

	var userDTOs []UserDTO
	for _, u := range users {
		role, err := roleapi.GetRoleByID(u.RoleId)
		if err != nil {
			logger.Error("Error retrieving role details", zap.Error(err))
			return nil, errors.New("error on communication api")
		}

		userDTOs = append(userDTOs, UserDTO{
			Id:    u.Id,
			Name:  u.Name,
			Email: u.Email,
			Role: RoleDTO{
				Role: NormalizeString(role.Role),
			},
		})
	}

	logger.Debug("Successfully retrieved users by criteria", zap.String("criteria", criteria), zap.String("value", value))
	return userDTOs, nil
}
