package service

import (
	"fmt"
	"github.com/YuStep/wbl0/internal/models"
	"github.com/YuStep/wbl0/internal/repository"
	"log"
)

type OrderService struct {
	repoDB    repository.Order
	repoCache repository.Cache
}

func NewOrderService(repoDB repository.Order, repoCache repository.Cache) *OrderService {
	return &OrderService{repoDB, repoCache}
}

func (o *OrderService) SaveOrder(order models.Order) {
	err := o.repoDB.SaveOrder(order)
	if err != nil {
		log.Print(err.Error())
	}

	o.repoCache.CreateCache(order)
}

func (o *OrderService) Preload() {
	orders, err := o.repoDB.GetOrders()
	if err != nil {
		fmt.Printf("Cannot insert order: %v\n", err)
	}

	o.repoCache.PreloadCache(orders)

}

func (o *OrderService) GetAll(limit int, offset int) []*models.Order {
	orders := o.repoCache.GetOrdersCache(limit, offset)

	return orders
}

func (o *OrderService) GetById(id string) (*models.Order, error) {
	order := o.repoCache.GetOrderByIdCache(id)

	return order, nil
}

func (o *OrderService) GetOrdersCount() int {
	return o.repoCache.GetOrdersCount()
}
