package main

import (
	"e-commerce/handler"
	"e-commerce/logger"
	"e-commerce/order"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "sa:admin123@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection failed")
		logger.ErrorLogger.Println("DB connection failed")
	} else {
		fmt.Println("Connection success")
		logger.InfoLogger.Println("Connection success")
	}

	orderRepository := order.NewRepository(db)
	orderService := order.NewService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)

	router := gin.Default()
	v1 := router.Group("/v1")

	done := make(chan bool)
	go func() {
		v1.GET("/Orders", orderHandler.GetOrders)
		done <- true
	}()
	go func() {
		v1.GET("/Order/:id", orderHandler.GetOrder)
		done <- true
	}()
	go func() {
		v1.POST("/Orders", orderHandler.PostOrder)
		done <- true
	}()
	go func() {
		v1.PUT("/Order/:id", orderHandler.UpdateOrder)
		done <- true
	}()

	<-done

	router.Run()
}
