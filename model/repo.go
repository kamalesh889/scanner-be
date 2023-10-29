package model

import "log"

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
