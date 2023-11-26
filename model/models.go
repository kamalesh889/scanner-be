package model

import "time"

type Store struct {
	Id       uint64 `gorm:"primaryKey"`
	Name     string
	Location string
	Schema   string
}

type User struct {
	Id       uint64 `gorm:"primaryKey"`
	Username string
	Password string
}

type UserOrder struct {
	UserId        uint64  `gorm:"not null"`
	TotalCost     float64 `gorm:"not null"`
	PaymentMethod string
	OrderDetails  map[string]interface{} `gorm:"type:json;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

/*OrderDetails :=
map[string]interface{}{
	"product1": map[string]interface{}{
		"quantity": 2,
		"price":    50.0,
	},
	"product2": map[string]interface{}{
		"quantity": 3,
		"price":    30.0,
	},
}

fmt.Println(OrderDetails)
*/
