package api

import (
	"context"
	"log"

	"github.com/huseyinbabal/microservices/payment/internal/application/core/domain"
	"github.com/huseyinbabal/microservices/payment/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	log.Println("Iniciando o Charge...")

	if payment.TotalPrice > 1000.00 {
		return domain.Payment{}, status.Error(codes.InvalidArgument, "Payment cannot be more than a thousand.")
	}
	err := a.db.Save(ctx, &payment)
	if err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}