package order

import "time"

type OrderRequest struct {
	Id            int                  `json:"id"`
	CustomerId    int                  `json:"customer_id"`
	OrderStatusId int                  `json:"order_status_id"`
	OrderDate     time.Time            `json:"order_date"`
	TotalPrice    float32              `json:"total_price"`
	OrderDetails  []OrderDetailRequest `json:"order_details"`
	Customer      CustomerRequest      `json:"customer"`
	OrderStatus   OrderDetailRequest   `json:"order_status"`
}

type OrderDetailRequest struct {
	Id         int     `json:"id"`
	OrderId    int     `json:"order_id"`
	ProductId  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float32 `json:"total_price"`
}

type CustomerRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProductRequest struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	InventoryId int     `json:"inventory_id"`
}

type OrderStatusRequest struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type InventoryRequest struct {
	Id    int `json:"id"`
	Stock int `json:"stock"`
}

func ConvertToOrder(order OrderRequest) Order {
	return Order{
		Id:            order.Id,
		CustomerId:    order.CustomerId,
		OrderStatusId: order.OrderStatusId,
		OrderDate:     order.OrderDate,
		TotalPrice:    order.TotalPrice,
	}
}
func ConvertToOrderDetail(orderDetail OrderDetailRequest) OrderDetail {
	return OrderDetail{
		Id:         orderDetail.Id,
		OrderId:    orderDetail.OrderId,
		ProductId:  orderDetail.ProductId,
		Quantity:   orderDetail.Quantity,
		TotalPrice: orderDetail.TotalPrice,
	}
}
