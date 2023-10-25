package model

type Product struct {
	Id       uint64 `gorm:"primaryKey"`
	Barcode  string
	Name     string
	Price    float64
	Discount float64
}
