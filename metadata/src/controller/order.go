package controller

import (
	"encoding/json"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common/log"
	"github.com/PillLabs/plugman-monorepo/metadata/src/model"
	"github.com/PillLabs/plugman-monorepo/metadata/src/service/api"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Order struct {
}

// PostOrder godoc
// @Summary      post order
// @Description  post order
// @Tags         order
// @Accept       json
// @Produce      json
// @Param		 address path string true "minter address"
// @Param		 request body model.PostOrderRequest true "post the transfer* tx hash back to server, let server tracking orders"
// @Success 20000 {object} common.Response "A successful response with ReturnCode' = 20000"
// @Failure 20001 {object} common.Response "invalid param"
// @Failure 20015 {object} common.Response "Server internal errors, just blame backend engineer"
// @Failure 20024 {object} common.Response "transaction not found"
// @Failure 20027 {object} common.Response "this chain id is not supported"
// @Router       /order/{address} [post]
func (o *Order) PostOrder(c *gin.Context) {
	appG := common.Gin{C: c}
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("[Order - PostOrder] read request body failed", err)
		appG.Response(http.StatusOK, common.INTERNAL_ERROR, nil)
		return
	}
	defer c.Request.Body.Close()

	var request model.PostOrderRequest
	err = json.Unmarshal(bodyBytes, &request)
	if err != nil {
		log.Info("[Order - PostOrder] unmarshal request body failed", err)
		appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
		return
	}

	rtCode, err := api.PostOrder(request)
	if err != nil {
		log.Error("[Order - PostOrder] post order error: %s", err)
	}
	if rtCode != common.SUCCESS {
		log.Info("[Order - PostOrder]  post order failed, return code is %v", rtCode)
		appG.Response(http.StatusOK, rtCode, nil)
		return
	}
	log.Debug("[Order - PostOrder] post order for %s successfully", request.FromAddress)
	appG.SuccessResponse(gin.H{})
}

// GetOrders godoc
// @Summary      get orders
// @Description  get orders
// @Tags         order
// @Accept       json
// @Produce      json
// @Param		 address path string true "minter address"
// @Success 20000 {object} common.Response{data=model.GetOrderResponse} "A successful response with ReturnCode' = 20000"
// @Failure 20001 {object} common.Response "invalid param"
// @Failure 20015 {object} common.Response "Server internal errors, just blame backend engineer"
// @Failure 20026 {object} common.Response "order not found"
// @Router       /order/{address} [get]
func (o *Order) GetOrders(c *gin.Context) {
	appG := common.Gin{C: c}
	address, ok := c.Params.Get("address")
	if !ok {
		log.Info("[Order - GetOrders] get address failed")
		appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
		return
	}
	if !common2.IsHexAddress(address) {
		log.Info("[Order - GetOrders] address %s is not valid")
		appG.Response(http.StatusOK, common.INVALID_PARAM, nil)
		return
	}
	resp, rtCode, err := api.GetOrders(address)
	if err != nil {
		log.Error("[Order - GetOrders] get order error: %s", err)
	}
	if rtCode != common.SUCCESS {
		log.Info("[Order - GetOrders] get order failed, return code is %d", rtCode)
		appG.Response(http.StatusOK, rtCode, nil)
		return
	}
	log.Debug("[Order - GetOrders] get order for %s successfully", address)
	appG.SuccessResponse(resp)
}
