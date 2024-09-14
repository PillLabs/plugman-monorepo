package controller

import (
	"github.com/PillLabs/plugman-monorepo/metadata/src/common"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common/log"
	"github.com/PillLabs/plugman-monorepo/metadata/src/service/api"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Whitelist struct {
}

// GetAddressWhitelist godoc
// @Summary      get whitelist
// @Description  get whitelist
// @Tags         whitelist
// @Accept       json
// @Produce      json
// @Param		 address path string true "minter address"
// @Success 20000 {object} common.Response{data=api.GetWhitelistResponse} "A successful response with ReturnCode' = 20000"
// @Failure 20001 {object} common.Response "invalid param"
// @Failure 20015 {object} common.Response "Server internal errors, just blame backend engineer"
// @Failure 20024 {object} common.Response "not have whitelist yet"
// @Router       /whitelist/{address} [get]
func (w *Whitelist) GetAddressWhitelist(c *gin.Context) {
	appG := common.Gin{C: c}
	_address, _ := c.Params.Get("address")
	ok := common2.IsHexAddress(_address)
	if !ok {
		log.Info("[Whitelist Controller - GetAddressWhitelist] %v is not valid address", _address)
		appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
		return
	}
	resp, rtCode, err := api.GetWhitelist(_address)
	if err != nil {
		log.Error("[Whitelist Controller - GetAddressWhitelist] get address whitelist: %v", err)
	}
	if rtCode != common.SUCCESS {
		log.Info("[Whitelist Controller - GetAddressWhitelist] get address %v whitelist failed, "+
			"return code is %v", _address, rtCode)
		appG.Response(http.StatusOK, rtCode, nil)
		return
	}
	log.Debug("[Whitelist Controller - GetAddressWhitelist] get  address %v whitelist successfully, "+
		"response '%s'", _address, *resp)
	appG.SuccessResponse(*resp)
}
