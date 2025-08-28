package api

import (
	"context"
	"log"
	"time"

	"github.com/gfedacs/microservices/order/internal/application/core/domain"
	"github.com/gfedacs/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db ports.DBPort
	payment ports.PaymentPort
	shipping ports.ShippingPort
	stock ports.StockPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, shipping ports.ShippingPort, stock ports.StockPort) *Application {
	return &Application{
		db: db,
		payment: payment,
		shipping: shipping,
		stock: stock,
	}
}

func (a Application) PlaceOrder(ctx context.Context,order domain.Order) (domain.Order, error) {
	ctxTimeout , cancel := context.WithTimeout (context.Background(), 2*time.Second )
	defer cancel()
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}

	productCodes := []string{}
	for _, item := range order.OrderItems {
		productCodes = append(productCodes, item.ProductCode)
	}

	exists, missing, err := a.stock.ExistsStockItems(productCodes)
	if err != nil {
		return domain.Order{}, status.Error(codes.Internal, "Error checking stock")
	}
	if !exists {
		return domain.Order{}, status.Errorf(codes.InvalidArgument, "The following items do not exist: %v", missing)
	}


	Quantity := int32(0)
	for _,item := range order.OrderItems {
		Quantity += item.Quantity
	}
	if Quantity > 50 {
		order.Status = "Canceled"
		quantityErr := a.db.Save(&order)
		if quantityErr != nil {
		return domain.Order{}, quantityErr
		}
		return order, status.Error(codes.InvalidArgument, "Quantity cannot be more than 50.")
	}

	paymentErr := a.payment.Charge(ctxTimeout, &order)
	if paymentErr != nil {
		if status.Code(paymentErr) == codes.DeadlineExceeded {
			log.Fatalf("Erro: %v",paymentErr)
		}
		order.Status = "Canceled"
		if saveErr := a.db.Save(&order); saveErr != nil {
			return domain.Order{}, saveErr
		}
		return order, paymentErr
	}

	deliveryDays, err := a.shipping.Create(ctxTimeout, &order)
	if err != nil {
		order.Status = "Paid"
		_ = a.db.Save(&order)
		return order, err
	}

	order.DeliveryDays = deliveryDays
	order.Status = "Paid"
	updateErr := a.db.Save(&order)
	if updateErr != nil {
		return domain.Order{}, err
	}
	return order, nil
}
