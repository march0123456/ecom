package service

import (
	"ecommerce/repository"
	"log"
)

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return productService{productRepo: productRepo}
}

func (s productService) GetAllProduct() ([]ProductResponse, error) {
	products, err := s.productRepo.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	productResponses := []ProductResponse{}
	for _, product := range products {
		productResponse := ProductResponse{
			ProductID:   product.ProductID,
			Name:        product.Name,
			Description: product.Description,
			Quantity:    product.Quantity,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}

func (s productService) GetProductById(id int) (*ProductResponse, error) {
	product, err := s.productRepo.GetById(id)
	if err != nil {
		log.Fatal(err)
	}

	productResponse := ProductResponse{
		ProductID:   product.ProductID,
		Name:        product.Name,
		Description: product.Description,
		Quantity:    product.Quantity,
	}

	return &productResponse, nil
}

func (s productService) CreateProduct(request ProductRequest) (*ProductResponse, error) {

	prod := repository.Product{
		Name:        request.Name,
		Description: request.Description,
		Quantity:    request.Quantity,
	}

	newProduct, err := s.productRepo.Create(prod)
	if err != nil {
		log.Fatal(err)
	}

	response := ProductResponse{
		ProductID:   newProduct.ProductID,
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Quantity:    newProduct.Quantity,
	}

	return &response, nil
}
