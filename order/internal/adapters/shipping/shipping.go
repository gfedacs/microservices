package shipping_adapter

import (
	"context"
	"time"

	"github.com/gfedacs/microservices-proto/golang/shipping"
	"github.com/gfedacs/microservices/order/internal/application/core/domain"


	
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	client shipping.ShippingClient
}

// Inicializa o adapter do Shipping
func NewAdapter(shippingServiceUrl string) (*Adapter, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithCodes(codes.DeadlineExceeded, codes.ResourceExhausted, codes.Unavailable),
			grpc_retry.WithMax(5),
			grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
		)),
	}

	conn, err := grpc.NewClient(shippingServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := shipping.NewShippingClient(conn)
	return &Adapter{client: client}, nil
}

func (a *Adapter) Create(ctx context.Context, order *domain.Order) (int32, error) {
	var items []*shipping.ShippingItem
	for _, i := range order.OrderItems {
		items = append(items, &shipping.ShippingItem{
			ProductCode: i.ProductCode,
			Quantity:    i.Quantity,
		})
	}

	resp, err := a.client.Create(ctx, &shipping.CreateShippingRequest{
		OrderId: order.ID,
		Items:   items,
	})
	if err != nil {
		return 0, err
	}

	return resp.DeliveryDays, nil
}