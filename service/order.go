package service

type OrderRequest struct {
	ProductID int `json:"product_id"`
}

type OrderResponse struct {
	OrderID            int    `json:"order_id"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	OrderDate          string `json:"order_date"`
}

type OrderService interface {
	CreateOrder(int, OrderRequest) (*OrderResponse, error)
	GetOrders(int) ([]OrderResponse, error)
	CancelOrder(int) error
	GetOrder(int) (*OrderResponse, error)
}
