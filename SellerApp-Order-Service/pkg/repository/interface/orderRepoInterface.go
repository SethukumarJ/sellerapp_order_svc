package interfaces

import (
	"context"

	"github.com/SethukumarJ/sellerapp-order-svc/pkg/domain"
	utils "github.com/SethukumarJ/sellerapp-order-svc/pkg/utils"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order domain.Order) (string, error)
	CreateItem(ctx context.Context, item domain.Item) (string, error)
	FindItem(ctx context.Context, id string ) (domain.Item, error)
	UpdateOrder(ctx context.Context, orderid, status string) (string, error)
	FetchOrder(ctx context.Context, userid int, filter domain.Filter, pagenation utils.Filter) ([]domain.Order, utils.Metadata, error)
}
