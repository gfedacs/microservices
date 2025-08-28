package db

import (
	"context"
	"fmt"

	"github.com/gfedacs/microservices/shipping/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)	

type Shipping struct {
	gorm.Model
	OrderID int64
	ShippingItems []ShippingItem
	DeliveryDays int32
}

type ShippingItem struct {
	gorm.Model
	ProductCode string
	Quantity int32
	ShippingID  uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, OpenErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if OpenErr != nil {
		return nil, fmt.Errorf("db connetction error: %v", OpenErr)
	}
	return &Adapter{db: db}, nil
}

func (a *Adapter) Save(ctx context.Context, shipping *domain.Shipping)  error{
	var shippingItems []ShippingItem
	for _, orderItem := range shipping.ShippingItems{
		shippingItems = append(shippingItems, ShippingItem{
			ProductCode: orderItem.ProductCode,
			Quantity: orderItem.Quantity,
			ShippingID: uint(shipping.ID),
		})
	}

	if shipping.ID == 0 {
		shippingModel := Shipping{
			OrderID: shipping.OrderID,
			DeliveryDays: shipping.DeliveryDays,
			ShippingItems: shippingItems,
		}
		res := a.db.Create(&shippingModel)
		if res.Error == nil{
			shipping.ID = int64(shippingModel.ID)
		}
		return res.Error
	}

	res := a.db.Model(&Shipping{}).Where("id = ?",shipping.ID).Updates(Shipping{
		OrderID: shipping.OrderID,
		ShippingItems: shippingItems,
		DeliveryDays: shipping.DeliveryDays,
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}