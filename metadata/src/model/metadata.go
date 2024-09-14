package model

import (
	"gorm.io/gorm"
)

type Metadata struct {
	gorm.Model
	OriginalId uint   `gorm:"not null;column:original_id;unique" json:"original_id"`
	OneOfOne   bool   `gorm:"not null;column:one_of_one;index:address_one_idx,priority:2" json:"one_of_one"`
	Body       string `gorm:"type:varchar(32);not null;column:body" json:"body"`
	BodyColor  string `gorm:"type:varchar(32);not null;column:body_color" json:"body_color"`
	Background string `gorm:"type:varchar(32);not null;column:background" json:"background"`
	Front      string `gorm:"type:varchar(32);not null;column:front" json:"front"`
	Side       string `gorm:"type:varchar(32);not null;column:side" json:"side"`
	Top        string `gorm:"type:varchar(32);not null;column:top" json:"top"`
	Address    string `gorm:"type:varchar(42);not null;column:address;index:address_one_idx,priority:1" json:"address"`
	Nonce      uint   `gorm:"not null;column:nonce;index:address_one_idx,priority:3" json:"nonce"`
	MintType   uint   `gorm:"not null;column:mint_type" json:"mint_type"`
	Timestamp  uint64 `gorm:"type:bigint;not null;column:timestamp" json:"timestamp"`
	TxHash     string `gorm:"type:varchar(100);not null;column:tx_hash;index:address_one_idx,priority:4" json:"tx_hash"`
	TokenId    string `gorm:"type:varchar(42);not null;column:token_id" json:"token_id"`
}

func (Metadata) TableName() string {
	return "metadata"
}

// SortMetadata impl sort.Interface
type SortMetadata []*Metadata

func (s SortMetadata) Len() int {
	return len(s)
}

func (s SortMetadata) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (s SortMetadata) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
