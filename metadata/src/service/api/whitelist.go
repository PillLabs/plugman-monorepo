package api

import (
	"github.com/PillLabs/plugman-monorepo/metadata/src/common"
	"github.com/PillLabs/plugman-monorepo/metadata/src/model"
	"github.com/PillLabs/plugman-monorepo/metadata/src/premise"
	"gorm.io/gorm"
)

type SetWhitelistRequest struct {
	Whitelists []model.AddressWL `json:"whitelists"`
}

type GetWhitelistResponse struct {
	OGCount  uint `json:"og_count"`
	OGMinted uint `json:"og_minted"`
	WLCount  uint `json:"wl_count"`
	WLMinted uint `json:"wl_minted"`
}

func GetWhitelist(address string) (*GetWhitelistResponse, int, error) {
	var nonce model.Nonce
	err := premise.RuntimeVars.Dao.GetNonce(address, &nonce)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, common.INTERNAL_ERROR, nil
	} else if err == gorm.ErrRecordNotFound {
		return nil, common.NOT_HAVE_WHITELIST, nil
	}
	return &GetWhitelistResponse{
		OGCount:  nonce.OGCount,
		OGMinted: nonce.OGMinted,
		WLCount:  nonce.WLCount,
		WLMinted: nonce.WLMinted,
	}, common.SUCCESS, nil
}
