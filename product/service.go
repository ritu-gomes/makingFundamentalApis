package product

import (
	"ecommerce/domain"
)

type service struct {
	prdRepo ProductRepo
}

func NewService(prdRepo ProductRepo) Service {
	return &service{
		prdRepo: prdRepo,
	}
}

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	return svc.prdRepo.Create(p)
}

func (svc *service) Get(productId int) (*domain.Product, error) {
	return svc.prdRepo.Get(productId)
}

func (svc *service) List() ([]*domain.Product, error) {
	return svc.prdRepo.List()
}

func (svc *service) Delete(productId int) error {
	return svc.Delete(productId)
}

func (svc *service) Update(p domain.Product) (*domain.Product, error) {
	return svc.prdRepo.Update(p)
}