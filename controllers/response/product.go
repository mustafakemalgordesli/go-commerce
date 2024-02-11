package response

import "github.com/mustafakemalgordesli/go-commerce/models"

type ProductResponse struct {
	Id int `json:"id"`
}

func ToModel(product models.Product) ProductResponse {
	return ProductResponse{
		Id: product.Id,
	}
}
