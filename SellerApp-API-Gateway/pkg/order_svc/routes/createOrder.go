package routes

import (
	"net/http"

	"github.com/SethukumarJ/sellerapp/pkg/domain"
	"github.com/SethukumarJ/sellerapp/pkg/order_svc/pb"
	"github.com/SethukumarJ/sellerapp/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Create Order
// @ID createorder
// @Tags Order-service
// @Produce json
// @Security BearerAuth
// @Param orderdetials body domain.Order{} true "Order Detials"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /order [post]
func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := domain.Order{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := ctx.Writer.Header().Get("userId")

	items := make([]*pb.Item, 0, len(body.Item))
	for _, pbItem := range body.Item {
		item := &pb.Item{
			ID:          pbItem.ID,
			Description: pbItem.Description,
			Price:       float32(pbItem.Price),
			Quantity:    int64(pbItem.Quantity),
		}
		items = append(items, item)
	}

	res, err := c.CreateOrder(ctx, &pb.CreateOrderRequest{
		UserId:       id,
		OrderId:      body.ID,
		Status:       body.Status,
		Item:         items,
		Total:        float32(body.Total),
		CurrencyUnit: body.CurrencyUnit,
	})

	if err != nil {
		responses := response.ErrorResponse("Failed to Create Order", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		response.ResponseJSON(*ctx, responses)
		return
	}

	responses := response.SuccessResponse(true, "SUCCESS", res)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	response.ResponseJSON(*ctx, responses)
	return

}
