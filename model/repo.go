package model

import (
	"fmt"
	"log"
)

func (r *Database) GetUser(req *User) (*User, error) {

	user := &User{}

	err := r.DbConn.First(&user, "username=? And password=?", req.Username, req.Password).Error
	if err != nil {
		log.Println("Error in Fetching user details", err)
		return nil, err
	}

	return user, nil

}

func (r *Database) CreateStore(req *Store) (*Store, error) {

	result := r.DbConn.Create(req)

	if result.Error != nil {
		return nil, result.Error
	}

	// Create schema for that store

	err := r.SchemaMigration(req.Schema)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (r *Database) GetStore(storeid uint64) (*Store, error) {

	store := Store{}

	err := r.DbConn.First(&store, "id=?", storeid).Error
	if err != nil {
		log.Println("Error in Fetching store details", err)
		return nil, err
	}

	return &store, nil
}

func (r *Database) GetProduct(code string, schema string) (*Product, error) {

	product := &Product{}

	query := fmt.Sprintf("select * from %s.products where barcode = ?", schema)

	err := r.DbConn.Raw(query, code).Scan(product).Error
	if err != nil {
		log.Println("Error in Fetching product details", err)
		return nil, err
	}

	return product, nil

}

func (r *Database) GetOrders(userid uint64) ([]UserOrder, error) {

	orders := []UserOrder{}

	query := "select * from user_order where user_id = ? limit 10"

	err := r.DbConn.Raw(query, userid).Scan(orders).Error
	if err != nil {
		log.Println("Error in Fetching User Order details", err)
		return nil, err
	}

	return orders, nil

}
