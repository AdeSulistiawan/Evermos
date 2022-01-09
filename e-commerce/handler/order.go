package handler

import (
	"e-commerce/logger"
	"e-commerce/order"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService order.Service
}

func NewOrderHandler(orderService order.Service) *orderHandler {
	return &orderHandler{orderService}
}

func (h *orderHandler) PostOrder(c *gin.Context) {
	var orderRequest order.OrderRequest

	err := c.ShouldBindJSON(&orderRequest)

	o, err := h.orderService.Save(orderRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":   http.StatusBadRequest,
			"errors": err,
		})
		logger.ErrorLogger.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    order.ConvertToOrderResponse(o),
			"message": "Successfully create order",
		})
		logger.InfoLogger.Println("Successfully create order")
	}
}

func (h *orderHandler) GetOrders(c *gin.Context) {
	orders, err := h.orderService.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":   http.StatusBadRequest,
			"errors": err,
		})
		logger.ErrorLogger.Println(err)
	} else if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    nil,
			"message": "There is no data order",
		})
		logger.InfoLogger.Println("There is no data order")
	} else {
		var ordersResponse []order.OrderResponse
		for _, o := range orders {
			orderResponse := order.ConvertToOrderResponse(o)
			ordersResponse = append(ordersResponse, orderResponse)
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    ordersResponse,
			"message": "Successfully get data order",
		})
		logger.InfoLogger.Println("Successfully get data order")
	}
}

func (h *orderHandler) GetOrder(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	o, err := h.orderService.GetById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":   http.StatusBadRequest,
			"errors": err,
		})
		logger.ErrorLogger.Println(err)
	} else if o.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    nil,
			"message": "There is no data order",
		})
		logger.InfoLogger.Println("There is no data order")
	} else {
		orderResponse := order.ConvertToOrderResponse(o)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    orderResponse,
			"message": "Successfully get data order",
		})
		logger.InfoLogger.Println("Successfully get data order")
	}

}

func (h *orderHandler) UpdateOrder(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var orderRequest order.OrderRequest

	err := c.ShouldBindJSON(&orderRequest)

	o, err := h.orderService.Update(id, orderRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":   http.StatusBadRequest,
			"errors": err,
		})
		logger.ErrorLogger.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    order.ConvertToOrderResponse(o),
			"message": "Successfully update order",
		})
		logger.InfoLogger.Println("Successfully update order")
	}
}
