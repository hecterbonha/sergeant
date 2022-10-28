package model

import (
	"time"
)

type Credentials struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	PassKey   string    `json:"passKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updateAt"`
}
