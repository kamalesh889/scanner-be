package model

import "gorm.io/gorm"

type Database struct {
	DbConn *gorm.DB
}
