package order

import (
	"github.com/SethukumarJ/sellerapp/pkg/order_svc/routes"
	"github.com/SethukumarJ/sellerapp/pkg/auth"
	"github.com/SethukumarJ/sellerapp/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	auth := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	order := r.Group("order")

	order.Use(auth.AuthRequired)
	order.POST("/", svc.CreateOrder)

	return svc
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
