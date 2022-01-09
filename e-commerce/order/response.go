package order

import "time"

type OrderResponse struct {
	Id            int                   `json:"id"`
	CustomerId    int                   `json:"customer_id"`
	OrderStatusId int                   `json:"order_status_id"`
	OrderDate     time.Time             `json:"order_date"`
	TotalPrice    float32               `json:"total_price"`
	OrderStatus   OrderStatusResponse   `json:"order_status"`
	Customer      CustomerResponse      `json:"customer"`
	OrderDetails  []OrderDetailResponse `json:"order_details"`
}

type OrderDetailResponse struct {
	Id         int     `json:"id"`
	OrderId    int     `json:"order_id"`
	ProductId  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float32 `json:"total_price"`
}

type CustomerResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProductResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	InventoryId int     `json:"inventory_id"`
}

type OrderStatusResponse struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type InventoryResponse struct {
	Id    int `json:"id"`
	Stock int `json:"stock"`
}

func ConvertToOrderResponse(order Order) OrderResponse {
	orderDetail := order.OrderDetails
	var orderDetailResponse []OrderDetailResponse

	for _, o := range orderDetail {
		orderDetailResponse = append(orderDetailResponse, convertToOrderDetailResponse(o))
	}

	return OrderResponse{
		Id:            order.Id,
		CustomerId:    order.CustomerId,
		OrderStatusId: order.OrderStatusId,
		OrderDate:     order.OrderDate,
		TotalPrice:    order.TotalPrice,
		OrderStatus:   OrderStatusResponse(order.OrderStatus),
		Customer:      CustomerResponse(order.Customer),
		OrderDetails:  orderDetailResponse,
	}
}

func convertToOrderDetailResponse(orderDetail OrderDetail) OrderDetailResponse {
	return OrderDetailResponse{
		Id:         orderDetail.Id,
		OrderId:    orderDetail.OrderId,
		ProductId:  orderDetail.ProductId,
		Quantity:   orderDetail.Quantity,
		TotalPrice: orderDetail.TotalPrice,
	}
}
