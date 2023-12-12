package cache

import (
	"fmt"
	"github.com/YuStep/wbl0/internal/models"
	"sync"
)

type OrderCache struct {
	cache map[string]*models.Order
	mu    sync.Mutex
}

func NewCache() *OrderCache {
	return &OrderCache{
		cache: map[string]*models.Order{},
		mu:    sync.Mutex{},
	}
}

func (oc *OrderCache) CreateCache(or models.Order) {

	oc.mu.Lock()
	oc.cache[or.OrderUID] = &or
	oc.mu.Unlock()
	fmt.Printf("Cache written: %s\n", or.OrderUID)
}

func (oc *OrderCache) Preload(orders []models.Order) {
	fmt.Printf("DB returns len: %d\n", len(orders))
	oc.mu.Lock()
	for _, or := range orders {
		oc.cache[or.OrderUID] = &or
	}
	oc.mu.Unlock()
}

func (oc *OrderCache) GetOrderByUid(uid string) *models.Order {
	return oc.cache[uid]
}

func (oc *OrderCache) GetOrders() map[string]*models.Order {
	return oc.cache
}

func (oc *OrderCache) GetLength() int {
	return len(oc.cache)
}
