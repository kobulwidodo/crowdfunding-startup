package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterInput) (User, error)
}

type service struct {
	repository Repositpory
}

func NewService(repository Repositpory) *service {
	return &service{repository: repository}
}

func (s *service) RegisterUser(input RegisterInput) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}

	user := User{
		Name:         input.Name,
		Occupation:   input.Occupation,
		Email:        input.Email,
		PasswordHash: string(passwordHash),
		Role:         "USER",
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
