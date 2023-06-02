package repository

import (
	"context"

	"github.com/SethukumarJ/sellerapp-order-svc/pkg/domain"
	interfaces "github.com/SethukumarJ/sellerapp-order-svc/pkg/repository/interface"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderDatabase struct {
	DB *gorm.DB
}

// FindItem implements interfaces.OrderRepository
func (c *orderDatabase) FindItem(ctx context.Context, id string) (domain.Item, error) {

	var item domain.Item
	 err := c.DB.Where("id = ?", id).First(&item).Error

	 return item, err
}

// CreateItem implements interfaces.OrderRepository
func (o *orderDatabase) CreateItem(ctx context.Context, item domain.Item) (string, error) {
	err := o.DB.Create(&item).Error
	return item.ID, err
}

// CreateOrder implements interfaces.OrderRepository
func (o *orderDatabase) CreateOrder(ctx context.Context, order domain.Order) (string, error) {
	order.ID = uuid.New().String()
	err := o.DB.Create(&order).Error
	return order.ID, err
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &orderDatabase{DB}
}
