package controller

import (
	"github.com/PillLabs/plugman-monorepo/metadata/src/common"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common/log"
	"github.com/PillLabs/plugman-monorepo/metadata/src/conf"
	"github.com/PillLabs/plugman-monorepo/metadata/src/service/api"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Mint struct {
}

// GetSignature godoc
// @Summary      get signature
// @Description  get signature
// @Tags         mint
// @Accept       json
// @Produce      json
// @Param		 address path string true "minter address"
// @Param		 mint_type path string true "mint type can be 2(WL sale) 3(Public sale)"
// @Param		 count path string true "purchase quantity for this minting"
// @Success 20000 {object} common.Response{data=api.GetSignatureResponse} "A successful response with ReturnCode' = 20000"
// @Failure 20001 {object} common.Response "invalid param"
// @Failure 20015 {object} common.Response "Server internal errors, just blame backend engineer"
// @Failure 20021 {object} common.Response "insufficient quota for WL sale"
// @Failure 20022 {object} common.Response "invalid mint type, this param can be 2 or 3 only"
// @Failure 20023 {object} common.Response "all metadata has been used"
// @Router       /mint/{address}/{mint_type}/{count}/getSignature [get]
func (m *Mint) GetSignature(c *gin.Context) {
	appG := common.Gin{C: c}
	_address, ok1 := c.Params.Get("address")
	_mintType, ok2 := c.Params.Get("mint_type")
	_count, ok3 := c.Params.Get("count")
	if !ok1 || !ok2 || !ok3 {
		log.Info("[Mint Controller - GetSignature] param not found, status: %v, %v, %v", ok1, ok2, ok3)
		appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
		return
	}
	ok := common2.IsHexAddress(_address)
	if !ok {
		log.Info("[Mint Controller - GetSignature] %v is not valid address", _address)
		appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
		return
	}
	mintType, err := strconv.Atoi(_mintType)
	if err != nil {
		log.Info("[Mint Controller - GetSignature] mint type is not integer: %v", _mintType)
		appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
		return
	}
	count, err := strconv.Atoi(_count)
	if err != nil {
		log.Info("[Mint Controller - GetSignature] count is not integer: %v", _count)
		appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
		return
	}
	if count <= 0 || count > conf.GConfig.GetInt("mint.maxCount") || mintType <= 0 || mintType > 4 {
		log.Info("[Mint Controller - GetSignature] mint type or count is not valid, mint type: %v, count: %v", mintType, count)
		appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
		return
	}
	request := api.GetSignatureRequest{
		Address:  _address,
		MintType: uint(mintType),
		Count:    uint(count),
	}
	resp, rtCode, err := api.GetSignature(&request)
	if err != nil {
		log.Error("[Mint Controller - GetSignature] get signature error: %v", err)
	}
	if rtCode != common.SUCCESS {
		log.Info("[Mint Controller - GetSignature] get signature failed, "+
			"request data '%s', return code is %v", request, rtCode)
		appG.Response(http.StatusOK, rtCode, nil)
		return
	}
	log.Debug("[Mint Controller - GetSignature] get signature successfully request '%s', response '%s'",
		request, *resp)
	appG.SuccessResponse(*resp)
}
