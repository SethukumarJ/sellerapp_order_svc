package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/SethukumarJ/sellerapp-order-svc/pkg/domain"
	repository "github.com/SethukumarJ/sellerapp-order-svc/pkg/repository/interface"
	interfaces "github.com/SethukumarJ/sellerapp-order-svc/pkg/usecase/interface"
	utils "github.com/SethukumarJ/sellerapp-order-svc/pkg/utils"
)

type orderUseCase struct {
	orderRepo repository.OrderRepository
}

func (o *orderUseCase) FetchOrder(ctx context.Context, userid int, filter domain.Filter, pagenation utils.Filter) ([]domain.ReqOrder, utils.Metadata, error) {
	order, metadata, err := o.orderRepo.FetchOrder(ctx, userid, filter, pagenation)
	if err != nil {
		return []domain.ReqOrder{}, metadata, err
	}

	

	Rorder := []domain.ReqOrder{}
	for _, od := range order {
		fmt.Println("itemid",od.Item_id)
		itemId := strings.Split(od.Item_id, ",")
		items := []domain.Item{}
		for _, id := range itemId {
			item, err := o.orderRepo.FindItem(ctx, id)
			if err == nil {
				items = append(items, item)
			}
			
		}
		recorder := domain.ReqOrder{
			ID:           od.ID,
			Status:       od.Status,
			Item:         items,
			Total:        od.Total,
			CurrencyUnit: od.CurrencyUnit,
		}
		Rorder = append(Rorder, recorder)

	}

	return Rorder, metadata, err
}

// UpdateOrder implements interfaces.OrderUseCase
func (o *orderUseCase) UpdateOrder(ctx context.Context, orderid string, status string) (string, error) {
	id, err := o.orderRepo.UpdateOrder(ctx, orderid, status)
	return id, err
}

// CreateOrder implements interfaces.OrderUseCase
func (o *orderUseCase) CreateOrder(ctx context.Context, order domain.ReqOrder) (string, error) {
	items := order.Item
	fmt.Println(len(items))
	if len(items) == 0 {
		return "", errors.New("there is no item in order")
	}

	var it []string
	var id string

	for _, item := range items {
		_, err := o.orderRepo.FindItem(ctx, item.ID)
		if err != nil {
			id, err = o.orderRepo.CreateItem(ctx, item)
			if err != nil {
				return "", err
			}
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
