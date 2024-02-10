package response

import "github.com/mustafakemalgordesli/go-commerce/models"

type AuthResponse struct {
	Id        int    `json:"id"`
	Email     string `json:"email" validate:"email,required"`
	FirstName string `json:"firstname" validate:"required,min=3"`
	LastName  string `json:"lastname"`
	Password  string `json:"password" validate:"required,min=6"`
}

func ToModel(user models.User) AuthResponse {
	return AuthResponse{
		Id:        user.Id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
	}
}
