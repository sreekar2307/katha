package response

import "github.com/sreekar2307/khata/model/table"

type User struct {
	Email string `json:"email"`
}

func NewUser(user table.User) User {
	return User{
		Email: user.Email,
	}
}
