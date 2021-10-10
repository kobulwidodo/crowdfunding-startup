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
		return
	}

	response := helper.ApiResponse("Akun berhasil dibuat", http.StatusOK, "sukses", user.FormatterUser(newUser, "tokentokentoken"))

	c.JSON(http.StatusOK, response)
	return
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ApiResponse("Harap isi semua input", http.StatusUnprocessableEntity, "gagal", helper.FormatBindError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var userLoggedin user.User
	userLoggedin, err := h.userService.Login(input)
	if err != nil {
		errorRes := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login gagal", http.StatusUnprocessableEntity, "gagal", errorRes)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Berhasil login", http.StatusOK, "sukses", user.FormatterUser(userLoggedin, "tokentokentoken"))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.EmailCheckInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ApiResponse("Harap isi semua Input", http.StatusUnprocessableEntity, "gagal", helper.FormatBindError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isAvail, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorRes := gin.H{"errors": "Server Failed"}
		response := helper.ApiResponse("Gagal memeriksa email", http.StatusInternalServerError, "gagal", errorRes)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	metaMsg := "Email sudah terdaftar"
	if isAvail {
		metaMsg = "Email Tersedia"
	}

	response := helper.ApiResponse(metaMsg, http.StatusOK, "sukses", gin.H{"is_available": isAvail})
	c.JSON(http.StatusOK, response)
}
