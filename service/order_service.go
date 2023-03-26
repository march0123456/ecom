package service

import (
	"ecommerce/repository"
	"log"
	"time"
)

type orderService struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
}

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository) OrderService {
	return orderService{orderRepo: orderRepo,
		productRepo: productRepo,
	}
}

func (s orderService) CreateOrder(userID int, request OrderRequest) (*OrderResponse, error) {
	order := repository.Order{
		ProductID: request.ProductID,
		UserID:    userID,
		OrderDate: time.Now().Format("2006-1-2 15:04:05"),
		Status:    1,
	}
	newOrder, err := s.orderRepo.Create(order)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	product, err := s.productRepo.GetById(order.ProductID)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	response := OrderResponse{
		OrderID:            newOrder.OrderID,
		ProductName:        product.Name,
		ProductDescription: product.Description,
		OrderDate:          newOrder.OrderDate,
	}
	return &response, nil

}
func (s orderService) GetOrders(userID int) ([]OrderResponse, error) {
	orders, err := s.orderRepo.GetAllByUserID(userID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responses := []OrderResponse{}
	for _, order := range orders {
		product, err := s.productRepo.GetById(order.ProductID)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		responses = append(responses, OrderResponse{
			OrderID:            order.OrderID,
			ProductName:        product.Name,
			ProductDescription: product.Description,
			OrderDate:          order.OrderDate,
		})
	}
	return responses, nil
}
func (s orderService) CancelOrder(orderID int) error {
	order := repository.Order{
		OrderID: orderID,
	}
	row, err := s.orderRepo.Delete(order)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if row > 0 {
		log.Fatal("Delete order success with row affected")
		return nil
	}
	log.Fatal("No order to delete")
	return nil
}
func (s orderService) GetOrder(orderID int) (*OrderResponse, error) {
	order, err := s.orderRepo.GetByOrderId(orderID)
	if err != nil {
		log.Fatal(err)
	}
	product, err := s.productRepo.GetById(order.ProductID)
	if err != nil {
		log.Fatal(err)
	}
	orderResponse := OrderResponse{
		OrderID:            order.OrderID,
		ProductName:        product.Name,
		ProductDescription: product.Description,
		OrderDate:          order.OrderDate,
	}

	return &orderResponse, nil
}
