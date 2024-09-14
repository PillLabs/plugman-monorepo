package model

import "gorm.io/gorm"

type AddressWL struct {
	Address string `json:"address"`
	OGCount uint   `json:"og_count"`
	WLCount uint   `json:"wl_count"`
}

type Nonce struct {
	gorm.Model
	Address     string `gorm:"type:varchar(42);not null;column:address;unique" json:"address"`
	OGCount     uint   `gorm:"not null;column:og_count;" json:"og_count"`
	OGMinted    uint   `gorm:"not null;column:og_minted;" json:"og_minted"`
	OGNonce     uint   `gorm:"not null;column:og_nonce" json:"og_nonce"`
	WLCount     uint   `gorm:"not null;column:wl_count" json:"wl_count"`
	WLMinted    uint   `gorm:"not null;column:wl_minted" json:"wl_minted"`
	WLNonce     uint   `gorm:"not null;column:wl_nonce" json:"wl_nonce"`
	PublicNonce uint   `gorm:"not null;column:public_nonce" json:"public_nonce"`
}

func (Nonce) TableName() string {
	return "nonce"
}
