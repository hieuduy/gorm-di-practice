package entity

import (
	"database/sql"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name 		 string 	`json:"name"`
	Email        *string	`gorm:"unique"`
	Age          uint8 		`gorm:"->:false;<-:create;default:100"` // Create-only (disabled read from db)
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedTs 	 int64 		`gorm:"autoUpdateTime"`	 // Use unix seconds as creating time
	UpdatedTs	 int		`gorm:"autoCreateTime:nano"`// Use unix nano seconds as updating time
	Hobbies		 datatypes.JSON
}
