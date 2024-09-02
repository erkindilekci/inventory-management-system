package service

import (
	"errors"
	"ims-intro/pkg/domain"
	"ims-intro/pkg/repository"
	"ims-intro/pkg/service/dto"
)

type IProductService interface {
	Add(productCreate *dto.ProductCreate) error
	GetAllProducts() []*domain.Product
	GetAllProductsByCategory(category string) []*domain.Product
	UpdateProductById(updatedProduct *dto.ProductCreate, productId int64) error
	DeleteById(productId int64) error
}

type ProductService struct {
	productRepository repository.IProductRepository
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &ProductService{productRepository}
}

func (service *ProductService) Add(productCreate *dto.ProductCreate) error {
	err := validateProductCreate(productCreate)
	if err != nil {
		return err
	}

	product := productCreateToProduct(productCreate)
	return service.productRepository.AddProduct(product)
}

func (service *ProductService) GetAllProducts() []*domain.Product {
	return service.productRepository.GetAllProducts()
}

func (service *ProductService) GetAllProductsByCategory(category string) []*domain.Product {
	return service.productRepository.GetProductsByCategory(category)
}

func (service *ProductService) UpdateProductById(updatedProduct *dto.ProductCreate, productId int64) error {
	err := service.productRepository.CheckProductExistence(productId)
	if err != nil {
		return nil
	}

	err = validateProductCreate(updatedProduct)
	if err != nil {
		return err
	}

	product := productCreateToProduct(updatedProduct)
	return service.productRepository.UpdateProductById(product, productId)
}

func (service *ProductService) DeleteById(productId int64) error {
	err := service.productRepository.CheckProductExistence(productId)
	if err != nil {
		return nil
	}

	return service.productRepository.DeleteProductById(productId)
}

func validateProductCreate(productCreate *dto.ProductCreate) error {
	if productCreate.Name == "" {
		return errors.New("name can't be empty")
	}
	if productCreate.Price < 0 {
		return errors.New("price can't be less than zero")
	}
	if productCreate.Quantity < 0 {
		return errors.New("quantity can't be less than zero")
	}
	if productCreate.Category == "" {
		return errors.New("category can't be empty")
	}
	return nil
}

func productCreateToProduct(productCreate *dto.ProductCreate) *domain.Product {
	return &domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Quantity: productCreate.Quantity,
		Category: productCreate.Category,
	}
}
