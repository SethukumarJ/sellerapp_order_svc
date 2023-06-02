package interfaces

import (
	"context"

	"github.com/SethukumarJ/sellerapp-order-svc/pkg/domain"
	utils "github.com/SethukumarJ/sellerapp-order-svc/pkg/utils"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, order domain.ReqOrder) (string, error)
	UpdateOrder(ctx context.Context, orderid, status string) (string, error)
	FetchOrder(ctx context.Context, userid int, filter domain.Filter, pagenation utils.Filter) ([]domain.ReqOrder, utils.Metadata, error)
}
