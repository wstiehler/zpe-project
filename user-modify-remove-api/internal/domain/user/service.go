package user

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/wstiehler/zpeupdateuser-api/internal/environment"
	"github.com/wstiehler/zpeupdateuser-api/internal/infrastructure/logger"
	roleapi "github.com/wstiehler/zpeupdateuser-api/internal/infrastructure/role-api"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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

func (s *Service) GetUserInformation(email, password string) (*UserEntity, *RoleDTO, error) {

	userSaved, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userSaved.Password), []byte(password))
	if err != nil {
		return nil, nil, errors.New("invalid email or password")
	}

	roleInfo, err := roleapi.GetRoleByID(userSaved.RoleId)
	if err != nil {
		return nil, nil, fmt.Errorf("error retrieving role details: %w", err)
	}

	role := RoleDTO{
		Role: roleInfo.Role,
	}

	return &userSaved, &role, nil
}

func (s *Service) GetPermissionInformation(roleID uint) ([]string, error) {

	permissionInfo, err := roleapi.GetPermissionByRoleID(roleID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving permission details: %w", err)
	}

	var permissions []string
	for _, permission := range *permissionInfo {
		permissions = append(permissions, permission.Name)
	}

	return permissions, nil
}

func (s *Service) CreateJWToken(userID string, role string, permissions []string) (string, error) {

	env := environment.GetInstance()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":         userID,
		"role":        role,
		"permissions": permissions,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(env.SECRET))
	if err != nil {
		return "", fmt.Errorf("failed to create token: %w", err)
	}

	return tokenString, nil
}

func (s *Service) Login(db *gorm.DB, user UserLoginEntity) (*UserLoginDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	userSaved, roleInfo, err := s.GetUserInformation(user.Email, user.Password)
	if err != nil {
		logger.Error("Error on login", zap.Error(err))
		return nil, err
	}

	permissions, err := s.GetPermissionInformation(userSaved.RoleId)
	if err != nil {
		logger.Error("Error on login", zap.Error(err))
		return nil, err
	}

	tokenString, err := s.CreateJWToken(userSaved.Id, roleInfo.Role, permissions)
	if err != nil {
		logger.Error("Error on login", zap.Error(err))
		return nil, err
	}

	userLoginDTO := UserLoginDTO{
		Id:    userSaved.Id,
		Email: userSaved.Email,
		Role:  roleInfo.Role,
		Token: tokenString,
		Name:  userSaved.Name,
	}

	return &userLoginDTO, nil
}

func (s *Service) GetUserByID(userId string, db *gorm.DB) (*UserEntity, error) {
	logger, dispose := logger.New()
	defer dispose()

	userSaved, err := s.repo.GetUserByID(userId)

	if err != nil {
		logger.Error("User not found")
		return nil, err
	}

	return &userSaved, nil
}

func (s *Service) UpdateUser(userId string, user UserEntity, db *gorm.DB) (*UserDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	roleInfo, err := roleapi.GetRoleByID(user.RoleId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving role details: %w", err)
	}

	productDTO := UserDTO{
		Id:    userId,
		Name:  user.Name,
		Role:  RoleDTO{Role: roleInfo.Role},
		Email: user.Email,
	}

	err = s.repo.UpdateUser(&user, userId)

	if err != nil {
		logger.Error("Error on update user", zap.String("userId", userId), zap.String("error", err.Error()))
		return nil, err
	}

	logger.Debug("Successfull on update user", zap.String("userId", userId))

	return &productDTO, nil
}

func (s *Service) DeleteUser(userId string, db *gorm.DB) error {
	logger, dispose := logger.New()
	defer dispose()

	err := s.repo.DeleteUser(userId)

	if err != nil {
		logger.Error("Error on delete user", zap.String("userId", userId), zap.String("error", err.Error()))
		return err
	}
	logger.Info("Successfull on delete user", zap.String("userId", userId))

	return nil
}
