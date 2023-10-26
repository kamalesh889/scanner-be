package model

import "log"

func (r *Database) GetUser(req *User) (*User, error) {

	user := &User{}

	err := r.DbConn.First(&user, "username=? And password=?", req.UserName, req.PassWord).Error
	if err != nil {
		log.Println("Error in Fetching user details", err)
		return nil, err
	}

	return user, nil

}
