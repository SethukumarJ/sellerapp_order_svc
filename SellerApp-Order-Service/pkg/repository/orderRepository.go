package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/SethukumarJ/sellerapp-order-svc/pkg/domain"
	interfaces "github.com/SethukumarJ/sellerapp-order-svc/pkg/repository/interface"
	utils "github.com/SethukumarJ/sellerapp-order-svc/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderDatabase struct {
	DB *gorm.DB
}

// FetchOrder implements interfaces.OrderRepository
func (o *orderDatabase) FetchOrder(ctx context.Context, userid int, filter domain.Filter, pagenation utils.Filter) ([]domain.Order, utils.Metadata, error) {

	sql := "SELECT * FROM orders WHERE 1=1"
	if filter.Status != "" {
		sql += fmt.Sprintf(" AND status='%s'", filter.Status)
	}
	if filter.MinTotal != 0 {
		sql += fmt.Sprintf(" AND total>=%v", filter.MinTotal)
	}
	if filter.MaxTotal != 0 {
		sql += fmt.Sprintf(" AND total<=%v", filter.MaxTotal)
	}
	if filter.SortBy != "" {
		sql += fmt.Sprintf(" ORDER BY %s", filter.SortBy)
		if filter.SortOrder == "desc" {
			sql += " DESC"
		}
	}

	sql += fmt.Sprintf(" LIMIT %v OFFSET %v", pagenation.Limit(), pagenation.Offset())

	order := []domain.Order{}
	var TotalRecords int64
	err := o.DB.Raw(sql).Scan(&order).Count(&TotalRecords).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("no order in the list")
	}

	return order, utils.ComputeMetaData(int(TotalRecords), pagenation.Page, pagenation.PageSize), err
}

// UpdateOrder implements interfaces.OrderRepository
func (o *orderDatabase) UpdateOrder(ctx context.Context, orderid string, status string) (string, error) {
	order := domain.Order{}
	err := o.DB.Model(&order).Where("id = ?", orderid).Update("status", status).Error
	return orderid, err
}

// FindItem implements interfaces.OrderRepository
func (c *orderDatabase) FindItem(ctx context.Context, id string) (domain.Item, error) {

	var item domain.Item
	err := c.DB.Where("id = ?", id).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("no item in the list")
	}


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
