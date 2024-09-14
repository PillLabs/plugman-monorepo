package model

import "gorm.io/gorm"

type Block struct {
	gorm.Model
	ChainName string `gorm:"type:varchar(30);not null;column:chain_name;unique" json:"chain_name"`
	ChainId   string `gorm:"type:varchar(20);not null;default:'0';column:chain_id" json:"chain_id"`
	Number    string `gorm:"type:varchar(42);not null;column:number" json:"number"`
}

func (Block) TableName() string {
	return "block"
}
