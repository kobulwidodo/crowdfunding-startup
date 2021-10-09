package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {

		response := helper.ApiResponse("Harap isi semua input", http.StatusBadRequest, "gagal", helper.FormatBindError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Terjadi kesalahan dari server", http.StatusInternalServerError, "gagal", nil)
		c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.ApiResponse("Akun berhasil dibuat", http.StatusOK, "sukses", user.FormatterUser(newUser, "tokentokentoken"))

	c.JSON(http.StatusOK, response)
	return
}
