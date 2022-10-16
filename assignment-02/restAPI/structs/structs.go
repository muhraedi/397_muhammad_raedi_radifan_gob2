package structs

import (
	"time"
)

type Order struct {
	ID           int       `json:"orderId" gorm:"primaryKey;autoIncrement"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID;references:ID"`
}

type Item struct {
	ID          int    `json:"lineItemId" gorm:"primaryKey;autoIncrement"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"orderId"`
}
