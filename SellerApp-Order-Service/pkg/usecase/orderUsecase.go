package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/SethukumarJ/sellerapp-order-svc/pkg/domain"
	repository "github.com/SethukumarJ/sellerapp-order-svc/pkg/repository/interface"
	interfaces "github.com/SethukumarJ/sellerapp-order-svc/pkg/usecase/interface"
)

type orderUseCase struct {
	orderRepo repository.OrderRepository
}

// CreateOrder implements interfaces.OrderUseCase
func (o *orderUseCase) CreateOrder(ctx context.Context, order domain.ReqOrder) (string, error) {
	items := order.Item
	fmt.Println(len(items))
	if len(items) == 0 {
		return "", errors.New("there no item in order")
	}

	var it []string
	for _, item := range items {
		id, err := o.orderRepo.CreateItem(ctx, item)
		if err != nil {
			return "", err
		}
		it = append(it, id)
	}
	itemIDs := strings.Join(it, ",")
	od := domain.Order{
		ID:           order.ID,
		Status:       order.Status,
		Item_id:      itemIDs,
		Total:        order.Total,
		CurrencyUnit: order.CurrencyUnit,
	}
	id, err := o.orderRepo.CreateOrder(ctx, od)
	return id, err
}

func NewOrderUseCase(repo repository.OrderRepository) interfaces.OrderUseCase {
	return &orderUseCase{
		orderRepo: repo,
	}
}
