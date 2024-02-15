package request

import "github.com/mustafakemalgordesli/go-commerce/models"

type CategoryCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Priorty  int    `json:"priorty"`
	ParentId int    `json:"parentid"`
	IsActive bool   `json:"isactive"`
}

func (categoryCreateRequest CategoryCreateRequest) ToModel() models.Category {
	return models.Category{
		Name:     categoryCreateRequest.Name,
		Priorty:  categoryCreateRequest.Priorty,
		ParentId: categoryCreateRequest.ParentId,
	}
}
