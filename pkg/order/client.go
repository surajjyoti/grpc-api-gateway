package order

import (
	"fmt"

	"github.com/surajjyoti-personal/api-gateway/pkg/config"
	"github.com/surajjyoti-personal/api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OrderServiceUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}
