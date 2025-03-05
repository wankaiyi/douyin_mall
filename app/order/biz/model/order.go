package model

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"strings"
	"time"
)

const (
	OrderStatusUnpaid = iota
	OrderStatusPaid
	OrderStatusCanceled
)

type Order struct {
	OrderID       string                `gorm:"primarykey;type:varchar(64)"`
	UserID        int32                 `gorm:"not null;type:int;index:idx_user_id_deleted_at_created_at,priority:1"`
	TotalCost     float64               `gorm:"not null;type:decimal(10,2)"`
	Name          string                `gorm:"not null;type:varchar(64)"`
	PhoneNumber   string                `gorm:"not null;type:char(11)"`
	Province      string                `gorm:"not null;type:varchar(64)"`
	City          string                `gorm:"not null;type:varchar(64)"`
	Region        string                `gorm:"not null;type:varchar(64)"`
	DetailAddress string                `gorm:"not null;type:varchar(255)"`
	Status        int32                 `gorm:"not null;type:int;default:0;index:idx_status_created_at_deleted_at,priority:1"`
	CreatedAt     time.Time             `gorm:"index:idx_user_id_deleted_at_created_at,priority:3;index:idx_status_created_at_deleted_at,priority:2"`
	DeletedAt     soft_delete.DeletedAt `gorm:"index:idx_user_id_deleted_at_created_at,priority:2;index:idx_status_created_at_deleted_at,priority:3"`
	OrderItems    []OrderItem           `gorm:"foreignKey:OrderID;references:OrderID"`
}

type OrderInfo struct {
	Order
	orderItems []OrderItem
}

func (o *Order) TableName() string {
	return "tb_order"
}

func CreateOrder(ctx context.Context, db *gorm.DB, order *Order) error {
	return db.WithContext(ctx).Create(order).Error
}

func GetOrdersByUserId(ctx context.Context, db *gorm.DB, id int32) (orderList []Order, err error) {
	err = db.WithContext(ctx).Model(&Order{}).
		Where(&Order{UserID: id}).
		Order("created_at DESC").
		Preload("OrderItems").
		Find(&orderList).Error
	return
}

func CheckOrderUnpaid(ctx context.Context, db *gorm.DB, orderId string) (bool, error) {
	var order Order
	err := db.WithContext(ctx).Model(&Order{}).Select("status").Where(&Order{OrderID: orderId}).First(&order).Error
	if err != nil {
		return false, err
	} else {
		return order.Status == OrderStatusUnpaid, nil
	}
}

func CancelOrder(ctx context.Context, db *gorm.DB, orderId string) (int64, error) {
	tx := db.WithContext(ctx).Model(&Order{}).
		Where(&Order{OrderID: orderId, Status: OrderStatusUnpaid}).
		Update("status", OrderStatusCanceled)
	return tx.RowsAffected, tx.Error
}

func CancelOrderList(ctx context.Context, db *gorm.DB, orderIdList []string) (int64, error) {
	tx := db.WithContext(ctx).Model(&Order{}).
		Where("order_id IN ? AND status = ?", orderIdList, OrderStatusUnpaid).
		Update("status", OrderStatusCanceled)
	return tx.RowsAffected, tx.Error
}

func GetOverdueOrder(ctx context.Context, db *gorm.DB, placeTime time.Time) (orderIdList []string, err error) {
	err = db.WithContext(ctx).Model(&Order{}).
		Where("status = ? and created_at < ?", OrderStatusUnpaid, placeTime).
		Order("created_at ASC").
		Pluck("order_id", &orderIdList).Error
	return
}

func MarkOrderPaid(ctx context.Context, db *gorm.DB, orderId string) (int64, error) {
	tx := db.WithContext(ctx).Model(&Order{}).
		Where(&Order{
			OrderID: orderId,
			Status:  OrderStatusUnpaid,
		}).Update("status", OrderStatusPaid)
	return tx.RowsAffected, tx.Error
}

func GetOrder(ctx context.Context, db *gorm.DB, orderId string) (order *Order, err error) {
	err = db.WithContext(ctx).Model(&Order{}).
		Where(&Order{OrderID: orderId}).
		Preload("OrderItems").
		First(&order).Error
	return
}

func SmartSearchOrder(ctx context.Context, db *gorm.DB, userId int32, terms []string, startTime string, endTime string) (orderList []Order, err error) {
	var likeConditions []string
	var args []interface{}
	for _, keyword := range terms {
		likeConditions = append(likeConditions, "oi.product_name LIKE ? OR oi.product_description LIKE ?")
		args = append(args, "%"+keyword+"%", "%"+keyword+"%")
	}

	query := strings.Join(likeConditions, " OR ")

	tx := db.WithContext(ctx).Model(&Order{}).
		Select("tb_order.*").
		Joins("INNER JOIN tb_order_item oi ON oi.order_id = tb_order.order_id").
		Where("tb_order.user_id = ?", userId)

	if startTime != "" && endTime != "" {
		tx = tx.Where("tb_order.created_at between ? and ?", startTime, endTime)
	}
	if terms != nil && len(terms) > 0 {
		tx = tx.Where(query, args...)
	}
	err = tx.Preload("OrderItems").
		Find(&orderList).Error
	return
}

func SelectCanceledSuccessOrders(ctx context.Context, db *gorm.DB, orderIdList []string) (canceledOrderIdList []string, err error) {
	err = db.WithContext(ctx).Model(&Order{}).
		Select("order_id").
		Where("order_id IN ? AND status = ?", orderIdList, OrderStatusCanceled).
		Pluck("order_id", &canceledOrderIdList).
		Error
	return
}
