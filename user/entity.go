package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
}

type RegisterInput struct {
	Name       string `binding:"required"`
	Occupation string `binding:"required"`
	Email      string `binding:"required,email"`
	Password   string `binding:"required"`
}

type LoginInput struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type EmailCheckInput struct {
	Email string `binding:"required,email"`
}
