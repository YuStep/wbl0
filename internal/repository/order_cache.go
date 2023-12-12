package repository

import (
	"github.com/YuStep/wbl0/internal/models"
	"github.com/YuStep/wbl0/pkg/cache"
)

type OrderCache struct {
	oc *cache.OrderCache
}

func NewOrderCache(cache *cache.OrderCache) *OrderCache {
	return &OrderCache{oc: cache}
}

func (oc *OrderCache) CreateCache(or models.Order) {
	oc.oc.CreateCache(or)
}

func (oc *OrderCache) PreloadCache(orders []models.Order) {
	oc.oc.Preload(orders)
}

func (oc *OrderCache) GetOrderByIdCache(id string) *models.Order {
	order := oc.oc.GetOrderByUid(id)
	return order
}

func (oc *OrderCache) GetOrdersCache(limit int, offset int) []*models.Order {
	orders := oc.oc.GetOrders()
	sliceOrders := make([]*models.Order, 0)
	counter := 0
	if len(orders) > limit {
		for _, value := range orders {
			counter++
			if counter >= offset {
				if counter > offset {
					sliceOrders = append(sliceOrders, value)
				}
				if counter-offset >= limit {
					break
				}
			}
		}
	} else {
		for _, value := range orders {
			sliceOrders = append(sliceOrders, value)
		}
	}

	return sliceOrders
}

func (oc *OrderCache) GetOrdersCount() int {
	return oc.oc.GetLength()
}
