package repository

type Product struct {
	ProductID   int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Quantity    int    `db:"quantity"`
}

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetById(int) (*Product, error)
	Create(Product) (*Product, error)
}
