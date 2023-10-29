package model

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
