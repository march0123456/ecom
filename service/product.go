package service

type ProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
type ProductResponse struct {
	ProductID   int    `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type ProductService interface {
	GetAllProduct() ([]ProductResponse, error)
	GetProductById(int) (*ProductResponse, error)
	CreateProduct(ProductRequest) (*ProductResponse, error)
}
