package request

import (
	"github.com/mustafakemalgordesli/go-commerce/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email     string `json:"email" validate:"email,required"`
	FirstName string `json:"firstname" validate:"required,min=3"`
	LastName  string `json:"lastname"`
	Password  string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=6"`
}

func (registerRequest RegisterRequest) ToModel() (models.User, error) {

	HashedByte, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 10)

	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Email:     registerRequest.Email,
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Password:  string(HashedByte),
	}, nil
}
