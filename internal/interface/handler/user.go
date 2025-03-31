package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tryuuu/go-task-api/internal/usecase"
)

type SignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// UserUsecaseはインターフェースを満たすためCreateUserメソッドを持つ
type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: uc,
	}
}

// ポインタレシーバー使用
func (h *UserHandler) SignupHandler(c *gin.Context) {
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.UserUsecase.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}
