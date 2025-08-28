package db

import(
	"fmt"
	"github.com/gfedacs/microservices/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)	

type Order struct {
	gorm.Model
	CustomerID int64
	Status string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice float32
	Quantity int32
	OrderId uint
}

type StockItem struct {
    ID          uint   `gorm:"primaryKey"`
    ProductCode string `gorm:"column:product_code"`
    Name        string
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

func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order
	res := a.db.First(&orderEntity, id)
	var orderItems []domain.OrderItem

	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice: orderItem.UnitPrice,
			Quantity: orderItem.Quantity,
		})
	}
	order := domain.Order{
		ID: int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		OrderItems: orderItems,
		CreatedAt: orderEntity.CreatedAt.UnixNano(),
	}
	return order, res.Error
}

func (a Adapter) Save(order *domain.Order) error{
	var orderItems []OrderItem
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice: orderItem.UnitPrice,
			Quantity: orderItem.Quantity,
			OrderId: uint(order.ID),
		})
	}
	if order.ID == 0{
		orderModel := Order{
		CustomerID: order.CustomerID,
		Status: order.Status,
		OrderItems: orderItems,
		}
		res := a.db.Create(&orderModel)
		if res.Error == nil {
			order.ID = int64(orderModel.ID)
		}
		return res.Error
	}

	res := a.db.Model(&Order{}).Where("id = ?",order.ID).Updates(Order{
		CustomerID: order.CustomerID,
		Status: order.Status,
		OrderItems: orderItems,
	})
	if res.Error != nil {
		return res.Error
	}
	return  nil
}

func (a *Adapter) ExistsStockItems(productCodes []string) (bool, []string, error) {
	var existing []StockItem
	if err := a.db.Where("product_code IN ?", productCodes).Find(&existing).Error; err != nil {
		return false, nil, err
	}

	missing := []string{}
	existingMap := make(map[string]bool)
	for _, item := range existing {
		existingMap[item.ProductCode] = true
	}

	for _, code := range productCodes {
		
		if !existingMap[code] {
			missing = append(missing, code)

		}
	}

	return len(missing) == 0, missing, nil
}