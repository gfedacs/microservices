package grpc

import (
	"context"
	"fmt"

	"github.com/gfedacs/microservices-proto/golang/shipping"
	"github.com/gfedacs/microservices/shipping/internal/application/core/domain"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)



func (a Adapter) Create(ctx context.Context, request *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error){
	log.WithContext(ctx).Info("Creating shipping...")

	orderItems := make([]domain.ShippingItem, len(request.GetItems()))
    for i, item := range request.GetItems() {
        orderItems[i] = domain.ShippingItem{
            ProductCode: item.ProductCode,
            Quantity:    item.Quantity,
        }
    }

	newShipping := domain.NewShipping(request.OrderId, orderItems)
	result, err := a.api.Process(ctx, newShipping)
	code := status.Code(err)
	if code == codes.InvalidArgument{
		return nil, err
	} else if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("Failed to process shipping. %v", err)).Err()
	}
	return &shipping.CreateShippingResponse{DeliveryDays: result.DeliveryDays}, nil
}