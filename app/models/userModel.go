package models

import "time"

type User struct {
	ID        uint      `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"column:username;type:varchar;size:255;not null" json:"username"`
	Email     string    `gorm:"column:email;type:varchar;size:255;not null" json:"email"`
	Password  string    `gorm:"column:password;type:varchar;not null" json:"password"`
	RoleID    int       `json:"role_id"`
	Role      Role      `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserList struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleName string `json:"role,omitempty"`
}
