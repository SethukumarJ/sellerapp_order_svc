package interfaces

import (
	"context"

	"github.com/SethukumarJ/sellerapp-order-svc/pkg/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order domain.Order) (string, error)
	CreateItem(ctx context.Context, item domain.Item) (string, error)
}
