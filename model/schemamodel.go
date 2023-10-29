package model

import (
	"fmt"
	"log"
)

type Product struct {
	Id       uint64 `gorm:"primaryKey"`
	Barcode  string
	Name     string
	Mrp      float64
	Discount float64
}

type Employee struct {
	Id       uint64 `gorm:"primaryKey"`
	Name     string
	Age      uint64
	Username string
	Password string
}

type Transaction struct {
	Id            uint64 `gorm:"primaryKey"`
	Value         float64
	NoofProducts  int
	PaymentMethod string
}

func (r *Database) SchemaMigration(schema string) error {

	query := fmt.Sprintf("CREATE SCHEMA %s AUTHORIZATION postgres", schema)

	result := r.DbConn.Exec(query)
	if result.Error != nil {
		return nil
	} else {
		log.Printf("Schema %s created successfully for ", schema)
	}

	table1 := fmt.Sprintf("%s.products", schema)
	table2 := fmt.Sprintf("%s.employees", schema)
	table3 := fmt.Sprintf("%s.transactions", schema)

	r.DbConn.Table(table1).AutoMigrate(&Product{})
	r.DbConn.Table(table2).AutoMigrate(&Employee{})
	r.DbConn.Table(table3).AutoMigrate(&Transaction{})

	log.Printf("Tables created successfully for schema %s", schema)

	return nil

}
