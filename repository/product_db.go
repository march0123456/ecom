package repository

import "github.com/jmoiron/sqlx"

type productRepositoryDB struct {
	db *sqlx.DB
}

func NewProductRepositoryDB(db *sqlx.DB) ProductRepository {
	return productRepositoryDB{db: db}
}

func (r productRepositoryDB) GetAll() ([]Product, error) {
	products := []Product{}
	query := "select * from products"
	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r productRepositoryDB) GetById(id int) (*Product, error) {
	product := Product{}
	query := "select * from products where product_id=?"
	err := r.db.Get(&product, query, id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r productRepositoryDB) Create(product Product) (*Product, error) {
	query := `INSERT INTO products (name, description, quantity) VALUES (?, ?, ?)`
	result, err := r.db.Exec(
		query,
		product.Name,
		product.Description,
		product.Quantity,
	)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	product.ProductID = int(id)

	return &product, nil
}
