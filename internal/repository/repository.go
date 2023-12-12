package repository

import (
	"github.com/YuStep/wbl0/internal/models"
	"github.com/YuStep/wbl0/pkg/cache"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Order interface {
	CreateTable() error
	SaveOrder(order models.Order) error
	GetOrders() ([]models.Order, error)
}

type Cache interface {
	CreateCache(or models.Order)
	PreloadCache(orders []models.Order)
	GetOrderByIdCache(id string) *models.Order
	GetOrdersCache(limit int, offset int) []*models.Order
	GetOrdersCount() int
}

type Repository struct {
	Order
	Cache
}

func NewRepository(db *pgxpool.Pool, cache *cache.OrderCache) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db),
		Cache: NewOrderCache(cache),
	}
}
