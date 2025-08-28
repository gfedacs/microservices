package ports

import (
	"github.com/gfedacs/microservices/shipping/internal/application/core/domain"
	"context"
)
type DBPort interface {
	Save(ctx context.Context ,shipping *domain.Shipping) error
}