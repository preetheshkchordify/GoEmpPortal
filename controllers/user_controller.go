package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preetheshkchordify/go-mongo-crud/dto"
	"github.com/preetheshkchordify/go-mongo-crud/errors"
	"github.com/preetheshkchordify/go-mongo-crud/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// RegisterUser handles user registration
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.RegisterUserRequest true "User registration details"
// @Success 201 {object} dto.RegisterUserResponse
// @Failure 400 {object} errors.AppError
// @Failure 409 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /api/users/register [post]
func (c *UserController) RegisterUser(ctx *gin.Context) {
	var req dto.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errors.WrapError(err, "Invalid input data"))
		return
	}

	resp, err := c.userService.RegisterUser(ctx.Request.Context(), req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			ctx.Error(appErr)
		} else {
			ctx.Error(errors.WrapError(err, "Failed to register user"))
		}
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// GetUserByID handles retrieving a user by ID
// @Summary Get user by ID
// @Description Get user details by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /api/users/{id} [get]
func (c *UserController) GetUserByID(ctx *gin.Context) {
	// TODO: Implement get user by ID
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

// GetUserByEmail handles retrieving a user by email
// @Summary Get user by email
// @Description Get user details by email address
// @Tags users
// @Accept json
// @Produce json
// @Param email query string true "User email"
// @Success 200 {object} models.User
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /api/users [get]
func (c *UserController) GetUserByEmail(ctx *gin.Context) {
	// TODO: Implement get user by email
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

// UpdateUser handles updating a user
// @Summary Update user
// @Description Update user details
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.User true "Updated user details"
// @Success 200 {object} models.User
// @Failure 400 {object} errors.AppError
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /api/users/{id} [put]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	// TODO: Implement update user
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

// DeleteUser handles deleting a user
// @Summary Delete user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /api/users/{id} [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	// TODO: Implement delete user
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}
