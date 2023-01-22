package product

import (
	"github.com/gin-gonic/gin"
	"github.com/surajjyoti-personal/api-gateway/pkg/auth"
	"github.com/surajjyoti-personal/api-gateway/pkg/config"
	"github.com/surajjyoti-personal/api-gateway/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.GetProduct)
}

func (svc *ServiceClient) GetProduct(ctx *gin.Context) {
	routes.GetProduct(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
