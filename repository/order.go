package repository

type Order struct {
	OrderID   int    `db:"id"`
	UserID    int    `db:"user_id"`
	ProductID int    `db:"product_id"`
	OrderDate string `db:"order_date"`
	Status    int    `db:"status"`
}

type OrderRepository interface {
	GetAllByUserID(int) ([]Order, error)
	GetByOrderId(int) (*Order, error)
	Create(Order) (*Order, error)
	Delete(Order) (int, error)
}
