package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type orderRepositoryDB struct {
	db *sqlx.DB
}

func NewOrderRepositoryDB(db *sqlx.DB) OrderRepository {
	return orderRepositoryDB{db: db}
}

func (r orderRepositoryDB) Create(od Order) (*Order, error) {
	query := "insert into orders (order_id, user_id, product_id, order_date, status) values (?, ?, ?, ?, ?)"
	result, err := r.db.Exec(
		query,
		od.OrderID,
		od.UserID,
		od.ProductID,
		od.OrderDate,
		od.Status,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	od.OrderID = int(id)

	return &od, nil
}

func (r orderRepositoryDB) Delete(od Order) (int, error) {
	result, err := r.db.Exec("DELETE FROM orders WHERE order_id = ?", od.OrderID)

	// check for errors
	if err != nil {
		log.Fatal(err)
	}

	// check the number of rows affected by the delete operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return int(rowsAffected), nil
}

func (r orderRepositoryDB) GetAllByUserID(userID int) ([]Order, error) {
	query := "SELECT * FROM orders WHERE user_id =?"
	orders := []Order{}
	err := r.db.Select(&orders, query, userID)
	if err != nil {
		log.Fatal(err)
	}
	return orders, nil
}

func (r orderRepositoryDB) GetByOrderId(orderID int) (*Order, error) {
	order := Order{}
	query := "SELECT * FROM orders WHERE order_id = ?"
	err := r.db.Get(&order, query, orderID)
	if err != nil {
		log.Fatal(err)
	}
	return &order, nil
}
