package order

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Order, error)
	Save(order Order) (Order, error)
	GetById(id int) (Order, error)
	Update(order Order) (Order, error)
	GetProductById(id int) (Product, error)
	GetStock(id int) (Inventory, error)
	UpdateStock(inv Inventory) (Inventory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Order, error) {
	var orders []Order

	err := r.db.Model(orders).
		Preload("OrderStatus").
		Preload("Customer").
		Preload("OrderDetails").
		Find(&orders).Error

	return orders, err
}

func (r *repository) Save(order Order) (Order, error) {
	err := r.db.Create(&order).Error
	return order, err
}

func (r *repository) GetById(id int) (Order, error) {
	var order Order
	err := r.db.
		Preload("OrderStatus").
		Preload("Customer").
		Preload("OrderDetails").
		Find(&order, id).Error
	return order, err
}

func (r *repository) Update(order Order) (Order, error) {
	err := r.db.Model(Order{}).Where("id = ?", order.Id).Updates(Order{OrderStatusId: order.OrderStatusId}).Error
	return order, err
}

func (r *repository) GetProductById(id int) (Product, error) {
	var product Product
	err := r.db.Find(&product, id).Error
	return product, err
}

func (r *repository) GetStock(id int) (Inventory, error) {
	var inventory Inventory
	err := r.db.Find(&inventory, id).Error
	return inventory, err
}

func (r *repository) UpdateStock(inv Inventory) (Inventory, error) {
	err := r.db.Save(&inv).Error
	return inv, err
}
