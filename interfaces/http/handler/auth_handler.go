package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authapp "github.com/vishalyadav0987/todo-list-api/internal/application/auth"
	"github.com/vishalyadav0987/todo-list-api/pkg/response"
)

type AuthHandler struct {
	registerUc *authapp.RegisterUsecase
}

func NewAuthHnadler(
	registerUc *authapp.RegisterUsecase,
) *AuthHandler {
	return &AuthHandler{
		registerUc: registerUc,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req authapp.RegisterRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid body")
	}

	user, err := h.registerUc.Execute(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"id":      user.ID(),
	})
}
