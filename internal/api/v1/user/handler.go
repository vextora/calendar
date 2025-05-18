package user

import (
	"net/http"
	"oncomapi/internal/api/v1/user/dto"
	"oncomapi/pkg/response"
	"oncomapi/pkg/validation"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService Service
}

func NewUserHandler(userService Service) *Handler {
	return &Handler{userService: userService}
}

func (h *Handler) Register(c *gin.Context) {
	input := dto.RegisterRequest{}

	if !validation.Validate(c, &input) {
		return
	}

	user, err := h.userService.Register(&input)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SendSuccess(c, DetailResponse(user), "User created")
}

func (h *Handler) Login(c *gin.Context) {
	input := dto.LoginRequest{}

	if !validation.Validate(c, &input) {
		return
	}

	token, err := h.userService.Login(&input)
	if err != nil {
		//logs.Error(err.Error())
		response.SendError(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.SendSuccess(c, gin.H{"token": token}, "Success")
}
