package model

import "gorm.io/gorm"

type TssTransfer struct {
	gorm.Model  `json:"-"`
	ChainId     string `gorm:"type:varchar(100);not null;column:chain_id" json:"chain_id"`
	FromAddress string `gorm:"type:char(42);not null;column:from_address" json:"from_address"`
	Amount      string `gorm:"type:varchar(100);not null;column:amount" json:"amount"`
	TxHash      string `gorm:"type:varchar(66);not null;column:tx_hash;unique" json:"tx_hash"`
	MintType    int    `gorm:"not null;column:mint_type" json:"mint_type"`
	Nonce       int    `gorm:"not null;column:nonce" json:"nonce"`
	OrderStatus int    `gorm:"not null;column:order_status" json:"order_status"`
}

func (TssTransfer) TableName() string {
	return "tss_transfer"
}

type PostOrderRequest struct {
	ChainId     string `json:"chain_id"`
	TxHash      string `json:"tx_hash"`
	FromAddress string `json:"from_address"`
	MintType    int    `json:"mint_type"`
	Nonce       int    `json:"nonce"`
}

type GetOrderResponse struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	TssTransfer TssTransfer `json:"tss_transfer"`
	TokenId     []string    `json:"token_id"`
}
