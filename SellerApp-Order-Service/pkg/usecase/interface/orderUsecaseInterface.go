package interfaces

import (
	"context"

	"github.com/SethukumarJ/sellerapp-order-svc/pkg/domain"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, order domain.ReqOrder) (string, error)
}
