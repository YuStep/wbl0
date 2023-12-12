package service

import (
	"github.com/YuStep/wbl0/internal/models"
	"github.com/YuStep/wbl0/internal/repository"
)

type Order interface {
	GetAll(limit int, offset int) []*models.Order
	GetOrdersCount() int
	GetById(id string) (*models.Order, error)
	Preload()
	SaveOrder(order models.Order)
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order, repos.Cache),
	}
}
