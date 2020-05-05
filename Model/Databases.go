package Model

import "github.com/jinzhu/gorm"

type Databases struct {
	gorm.Model
	Database string `gorm:"column:Database"`
}
