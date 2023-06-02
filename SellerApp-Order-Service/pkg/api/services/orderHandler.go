package services

import (
	"context"
	"net/http"

	"github.com/SethukumarJ/sellerapp-order-svc/pkg/domain"
	pb "github.com/SethukumarJ/sellerapp-order-svc/pkg/pb"
	usecase "github.com/SethukumarJ/sellerapp-order-svc/pkg/usecase/interface"
	utils "github.com/SethukumarJ/sellerapp-order-svc/pkg/utils"
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

// UpdateOrder implements pb.OrderServiceServer
func (c *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	id, err := c.orderUseCase.UpdateOrder(ctx, req.OrderId, req.Status)
	if err != nil {
		return &pb.UpdateOrderResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

	return &pb.UpdateOrderResponse{
		Status: http.StatusOK,
		Id:     id,
	}, nil

}


// FetchOrder implements pb.OrderServiceServer
func (c *OrderService) FetchOrder(ctx context.Context, req *pb.FetchOrderRequest) (*pb.FetchOrderResponse, error) {
	filter := domain.Filter{
		Status:    req.Status,
		MinTotal:  float64(req.MinTotal),
		MaxTotal:  float64(req.MaxTotal),
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}

	pagnation := utils.Filter{
		Page:     int(req.Filter.Page),
		PageSize: int(req.Filter.PageSize),
	}

	orders, Metadata, err := c.orderUseCase.FetchOrder(ctx, int(req.UserId), filter, pagnation)
	if err != nil {
		return &pb.FetchOrderResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}
	var Od []*pb.Order
	for _, order := range orders {
		var reitems []*pb.Item
		for _, items := range order.Item {
			items := pb.Item{
				ID:          items.ID,
				Description: items.Description,
				Price:       float32(items.Price),
				Quantity:    int64(items.Quantity),
			}
			reitems = append(reitems, &items)
		}

		od := pb.Order{
			OrderId:      order.ID,
			Status:       order.Status,
			Item:         reitems,
			Total:        float32(order.Total),
			CurrencyUnit: order.CurrencyUnit,
		}
		Od = append(Od, &od)
	}

	return &pb.FetchOrderResponse{
		Status: http.StatusOK,
		Orders: Od,
		Metadata: &pb.Metadata{
			CurrentPage:  int64(Metadata.CurrentPage),
			PageSize:     int64(Metadata.PageSize),
			FirstPage:    int64(Metadata.FirstPage),
			LastPage:     int64(Metadata.LastPage),
			TotalRecords: int64(Metadata.TotalRecords),
		},
	}, nil

}

func NewOrderService(usecase usecase.OrderUseCase) *OrderService {
	return &OrderService{
		orderUseCase: usecase,
	}
}
