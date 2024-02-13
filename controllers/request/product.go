package request

import "github.com/mustafakemalgordesli/go-commerce/models"

type ProductCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Price       uint64 `json:"price" validate:"required"`
	Description string `json:"description"`
	// ImageUrl    string `json:"imageurl"`
	CategoryId int `json:"categoryid"`
	Quantity   int `json:"quantity"`
}

func (productCreateRequest ProductCreateRequest) ToModel() (models.Product, error) {
	return models.Product{
		Name:        productCreateRequest.Name,
		Price:       productCreateRequest.Price,
		Description: productCreateRequest.Description,
		CategoryId:  productCreateRequest.CategoryId,
		Quantity:    productCreateRequest.Quantity,
	}, nil
}
