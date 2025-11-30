package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          uint      `gorm:"primaryKey;autoIncrement; not null"`
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();not null"`
	Name        string    `gorm:"type:varchar(100); not null"`
	Username    string    `gorm:"type:varchar(50); not null;unique"`
	Photo       string    `gorm:"type:varchar(255);"`
	Email       string    `gorm:"type:varchar(100); not null;unique"`
	Password    string    `gorm:"type:varchar(255); not null"`
	PhoneNumber string    `gorm:"type:varchar(15); not null;unique"`
	Gender      string    `gorm:"type:varchar(10); not null"`
	DateOfBirth string    `gorm:"type:varchar(10); not null"` // format YYYY-MM-DD
	IsVerified  bool      `gorm:"type:boolean;default:false;index:idx_users_is_verified"`
	Lat         string    `gorm:"varchar(50); not null"`
	Lng         string    `gorm:"varchar(50); not null"`
	Address     string    `gorm:"type:text; not null"`
	RoleId      uint      `gorm:"type:uint;not null"`
	Created_At  *time.Time
	Updated_At  *time.Time
	Role        []Role `gorm:"many2many:user_role;"`
}
