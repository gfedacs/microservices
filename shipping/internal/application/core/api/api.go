package api

import(
	"context"
	"log"

	"github.com/gfedacs/microservices/shipping/internal/application/core/domain"
	"github.com/gfedacs/microservices/shipping/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application{
	return &Application{
		db: db,
	}
}

func (a Application) Process(ctx context.Context,shipping domain.Shipping) (domain.Shipping, error) {
log.Println("Processing Delivery...")

err := a.db.Save(ctx, &shipping)
if err != nil {
	return domain.Shipping{}, err
}
return shipping, nil
}