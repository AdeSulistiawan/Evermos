package order

import (
	"time"
)

type Order struct {
	Id            int `gorm:"primaryKey"`
	CustomerId    int
	OrderStatusId int
	OrderDate     time.Time
	TotalPrice    float32
	OrderStatus   OrderStatus
	Customer      Customer
	OrderDetails  []OrderDetail
}

type OrderDetail struct {
	Id         int `gorm:"primaryKey"`
	OrderId    int
	ProductId  int
	Quantity   int
	TotalPrice float32
}

type Customer struct {
	Id   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Price       float32
	InventoryId int
}

type OrderStatus struct {
	Id          int `gorm:"primaryKey"`
	Description string
}

type Inventory struct {
	Id    int `gorm:"primaryKey"`
	Stock int
}
