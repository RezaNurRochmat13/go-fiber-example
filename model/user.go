package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
