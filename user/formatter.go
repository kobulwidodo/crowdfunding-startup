package user

type UserFormat struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatterUser(user User, token string) UserFormat {
	newUser := UserFormat{
		Id:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return newUser
}
