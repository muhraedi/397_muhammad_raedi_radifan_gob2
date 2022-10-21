package repository

import "hello-testing/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
