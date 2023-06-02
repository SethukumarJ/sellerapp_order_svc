package routes

import (
	"net/http"

	"github.com/SethukumarJ/sellerapp/pkg/domain"
	"github.com/SethukumarJ/sellerapp/pkg/order_svc/pb"
	"github.com/SethukumarJ/sellerapp/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Update Order
// @ID Updateorder
// @Tags Order-service
// @Produce json
// @Security BearerAuth
// @Param updateorderdetials body domain.UpdateOrder{} true "Update Order Detials"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /order [put]
func UpdateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := domain.UpdateOrder{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := ctx.Writer.Header().Get("userId")

	res, err := c.UpdateOrder(ctx, &pb.UpdateOrderRequest{
		UserId:  id,
		OrderId: body.ID,
		Status:  body.Status,
	})

	if err != nil {
		responses := response.ErrorResponse("Failed to Update Order", err.Error(), nil)
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
