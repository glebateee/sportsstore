package models

type Repository interface {
	GetProduct(id int) Product
	GetProducts() []Product
	SaveProduct(*Product)
	GetCategories() []Category
	SaveCategory(*Category)
	Seed()

	GetProductPage(page, pageSize int) (products []Product, totalAvailable int)
	GetProductPageCategory(categoryId int, page, pageSize int) (products []Product, totalAvailable int)

	GetOrder(id int) Order
	GetOrders() []Order
	SaveOrder(*Order)
}
