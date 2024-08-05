package models

import "time"

type Role struct {
	ID        uint      `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
	RoleName  string    `gorm:"column:role_name;type:varchar;size:255;not null" json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleRequest struct {
	RoleName string `json:"role_name"`
}

type RoleList struct {
	ID       uint   `json:"id"`
	RoleName string `json:"role_name"`
}
