package domain

import "time"

type Shipping struct {
	ID int64
	OrderID int64
	ShippingItems []ShippingItem
	DeliveryDays int32 
	CreatedAt int64
}

type ShippingItem struct {
	ProductCode string
	Quantity int32
}

func NewShipping(orderID int64, shippingItems []ShippingItem) Shipping{
	deliveryDays := calcDeliveryDays(shippingItems)
	return Shipping{
		OrderID: orderID,
		ShippingItems: shippingItems,
		DeliveryDays: deliveryDays,
		CreatedAt: time.Now().Unix(),
	}
}

func calcDeliveryDays(shippingItems []ShippingItem) int32 {
	totalQuant := int32(0)

	for _, item := range shippingItems {
		totalQuant += item.Quantity
	}

	deliveryDays := int32(1)
	if totalQuant > 0 {
		adicionalDays := (totalQuant -1 ) / 5
		deliveryDays += adicionalDays
	}
	return deliveryDays
}