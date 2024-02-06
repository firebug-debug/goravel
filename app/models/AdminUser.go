package models

import (
	"github.com/goravel/framework/database/orm"
)

type AdminUser struct {
	orm.Model
	Id       int16 `gorm:"primaryKey"`
	Username string
	Name     string
	Avatar   string
	Password string
	Balance  float32
	orm.SoftDeletes
}
