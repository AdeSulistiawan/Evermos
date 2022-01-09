package order

import (
	"e-commerce/logger"
	"errors"
)

type Service interface {
	GetAll() ([]Order, error)
	Save(order OrderRequest) (Order, error)
	GetById(id int) (Order, error)
	Update(id int, order OrderRequest) (Order, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Order, error) {
	return s.repository.GetAll()
}

func (s *service) Save(order OrderRequest) (Order, error) {
	var err error
	orderDetailRequest := order.OrderDetails

	var orderDetails []OrderDetail
	for _, o := range orderDetailRequest {
		orderDetails = append(orderDetails, ConvertToOrderDetail(o))
		product, _ := s.repository.GetProductById(o.ProductId)
		stock, _ := s.repository.GetStock(product.Id)

		if stock.Stock-o.Quantity < 0 {
			logger.WarningLogger.Println("Stock limited")
			return ConvertToOrder(order), errors.New("Stock limited")
		}
	}

	data := Order{
		CustomerId:    order.CustomerId,
		TotalPrice:    order.TotalPrice,
		OrderStatusId: order.OrderStatusId,
		OrderDate:     order.OrderDate,
		OrderDetails:  orderDetails,
	}

	newOrder, err := s.repository.Save(data)

	return newOrder, err
}

func (s *service) GetById(id int) (Order, error) {
	return s.repository.GetById(id)
}

func (s *service) Update(ID int, orderReq OrderRequest) (Order, error) {
	orderData, err := s.repository.GetById(ID)

	if orderReq.OrderStatusId == Paid {
		orderDetailRequest := orderReq.OrderDetails
		var orderDetails []OrderDetail
		for _, o := range orderDetailRequest {
			orderDetails = append(orderDetails, ConvertToOrderDetail(o))
			product, _ := s.repository.GetProductById(o.ProductId)
			stock, _ := s.repository.GetStock(product.Id)
			stock.Stock = stock.Stock - o.Quantity

			if stock.Stock-o.Quantity < 0 {
				logger.WarningLogger.Println("Stock limited")
				return ConvertToOrder(orderReq), errors.New("Stock limited")
			} else {
				s.repository.UpdateStock(stock)
			}
		}
	}

	orderData.CustomerId = orderReq.CustomerId
	orderData.OrderStatusId = orderReq.OrderStatusId
	orderData.OrderDate = orderReq.OrderDate
	orderData.TotalPrice = orderReq.TotalPrice

	orderUpdate, err := s.repository.Update(orderData)

	return orderUpdate, err
}
