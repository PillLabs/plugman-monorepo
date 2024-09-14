package model

import "gorm.io/gorm"

type Rarity struct {
	gorm.Model
	Name     string `gorm:"type:varchar(32);not null;gorm:name;unique" json:"name"`
	TraitKey string `gorm:"type:varchar(32);not null;column:trait_key" json:"trait_key"`
	Rarity   string `gorm:"type:varchar(32);not null;column:rarity" json:"rarity"`
}

func (Rarity) TableName() string {
	return "rarity"
}
