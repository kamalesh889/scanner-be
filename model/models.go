package model

type Store struct {
	Id       uint64 `gorm:"primaryKey"`
	Name     string
	Location string
	Schema   string
}
