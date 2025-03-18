package services

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/preetheshkchordify/go-mongo-crud/dto"
	"github.com/preetheshkchordify/go-mongo-crud/models"
	"github.com/preetheshkchordify/go-mongo-crud/repositories"
	"github.com/preetheshkchordify/go-mongo-crud/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repositories.NewUserRepository(),
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req dto.RegisterUserRequest) (*dto.RegisterUserResponse, error) {
	// Validate role ID
	roleID, err := primitive.ObjectIDFromHex(req.Role)
	if err != nil {
		return nil, fmt.Errorf("invalid role ID format: %v", err)
	}

	// Check if user already exists
	existingUser, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	// Generate and hash password
	defaultPassword := utils.GeneratePassword()
	hashedPassword, err := utils.HashPassword(defaultPassword)
	if err != nil {
		return nil, fmt.Errorf("error generating password hash: %v", err)
	}

	// Create new user
	newUser := &models.User{
		ID:           primitive.NewObjectID(),
		Name:         req.Name,
		Email:        req.Email,
		EmployeeID:   req.EmployeeID,
		Designation:  req.Designation,
		RoleID:       roleID,
		Password:     hashedPassword,
		IsFirstLogin: true,
	}

	// Save user to database
	createdUser, err := s.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	// Send welcome email
	go func() {
		templatePath := filepath.Join("templates", "welcome_email.html")
		emailData := map[string]string{
			"email":           createdUser.Email,
			"defaultPassword": defaultPassword,
			"name":            createdUser.Name,
		}

		err := utils.SendEmail(utils.Email{
			To:           []string{createdUser.Email},
			Subject:      "Your account details",
			TemplatePath: templatePath,
			Placeholders: emailData,
			CC:           []string{},
			Attachments:  []string{},
		})
		if err != nil {
			fmt.Printf("Error sending welcome email: %v\n", err)
		}
	}()

	return &dto.RegisterUserResponse{
		Message:  "User registered successfully",
		Username: req.Email,
		Password: defaultPassword,
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	return s.userRepo.FindByID(ctx, id)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.userRepo.FindByEmail(ctx, email)
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	return s.userRepo.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	return s.userRepo.Delete(ctx, id)
}
