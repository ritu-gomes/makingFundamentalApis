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

func (svc *service) List(page, limit int64) ([]*domain.Product, error) {
	return svc.prdRepo.List(page, limit)
}

func (svc *service) Delete(productId int) error {
	return svc.prdRepo.Delete(productId)
}

func (svc *service) Count() (int64, error) {
	return svc.prdRepo.Count()
}

func (svc *service) Update(p domain.Product) (*domain.Product, error) {
	return svc.prdRepo.Update(p)
}