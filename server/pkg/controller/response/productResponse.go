package response

import "ims-intro/pkg/domain"

type ProductResponse struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int64   `json:"quantity"`
	Category string  `json:"category"`
}

func toProductResponse(product *domain.Product) *ProductResponse {
	return &ProductResponse{
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
		Category: product.Category,
	}
}

func ToProductResponseList(products []*domain.Product) []*ProductResponse {
	var responses []*ProductResponse
	for _, product := range products {
		responses = append(responses, toProductResponse(product))
	}
	return responses
}
