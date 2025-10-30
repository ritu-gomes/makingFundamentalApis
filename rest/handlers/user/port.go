package user

import "ecommerce/domain"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	Get(email, pass string) (*domain.User, error)
}