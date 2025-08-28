package domain

import "time";

type OrderItem struct {
	ProductCode string
	UnitPrice float32
	Quantity int32
}

type Order struct {
	ID int64
	CustomerID int64
	Status string
	OrderItems []OrderItem
	CreatedAt int64
	DeliveryDays int32 
}

func NewOrder(customerId int64, orderItems []OrderItem,) Order{
	return Order{
		CreatedAt: time.Now().Unix(),
		Status: "Pending",
		CustomerID: customerId,
		OrderItems: orderItems,
	}
}

func (o *Order) TotalPrice() float32{
	var totalPrice float32
	for _, orderItem := range o.OrderItems {
		totalPrice += orderItem.UnitPrice * float32(orderItem.Quantity)
	}
	return totalPrice
}