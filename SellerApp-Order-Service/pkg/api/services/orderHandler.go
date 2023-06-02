package services

import (
	"context"
	"net/http"

	"github.com/SethukumarJ/sellerapp-order-svc/pkg/domain"
	usecase "github.com/SethukumarJ/sellerapp-order-svc/pkg/usecase/interface"
	pb "github.com/SethukumarJ/sellerapp-order-svc/pkg/pb"
)

type OrderService struct {
	orderUseCase usecase.OrderUseCase
}

// CreateOrder implements pb.OrderServiceServer
func (c *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	items := make([]domain.Item, 0, len(req.Item))
	for _, item := range req.Item {
		items = append(items, domain.Item{
			ID:          item.ID,
			Description: item.Description,
			Price:       float64(item.Price),
			Quantity:    int(item.Quantity),
		})
	}

	order := domain.ReqOrder{
		ID:           req.OrderId,
		Status:       req.Status,
		Item:         items,
		Total:        float64(req.Total),
		CurrencyUnit: req.CurrencyUnit,
	}

	id, err := c.orderUseCase.CreateOrder(ctx, order)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusOK,
		Id:     id,
	}, nil

}

func NewOrderService(usecase usecase.OrderUseCase) *OrderService {
	return &OrderService{
		orderUseCase: usecase,
	}
}
