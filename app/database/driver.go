package database

import (
	"gorm.io/gorm"
)

type Driver interface {
	Ping()
	Close() bool
	Driver() *gorm.DB
}
