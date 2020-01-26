package sdk

import (
	"encoding/json"
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/enum"
	"github.com/RiverDanceGit/yeepayGo/params"
	"github.com/RiverDanceGit/yeepayGo/request"
	"github.com/RiverDanceGit/yeepayGo/response"
	"github.com/RiverDanceGit/yeepayGo/util"
	"strconv"
)

func NewApiStdTradeOrder(config yeepayGo.YeepayConfig, logger yeepayGo.YeepayLoggerInterface) ApiStdTradeOrder {
	return ApiStdTradeOrder{config, logger}
}

type ApiStdTradeOrder struct {
	config yeepayGo.YeepayConfig
	logger yeepayGo.YeepayLoggerInterface
}

// 创建订单
// https://open.yeepay.com/docs/e-commerceprotocols/rest__v1.0__std__trade__order.html
func (obj ApiStdTradeOrder) GetRequest(orderId, orderAmount, goodsName, goodsDesc, notifyUrl string) request.StdTradeOrderRequest {
	goodsParamExt := params.GoodsParamExt{goodsName, goodsDesc}
	goodsParamExtBody, _ := json.Marshal(goodsParamExt)

	req := request.NewStdTradeOrderRequest()
	req.SetHttpMethod(enum.HTTP_METHOD_POST)
	req.SetMethod("/rest/v1.0/std/trade/order")
	req.SetVersion("1.0")
	req.SetParentMerchantNo(obj.config.GetParentMerchantNo())
	req.SetMerchantNo(obj.config.GetMerchantNo())
	req.SetOrderId(orderId)
	req.SetOrderAmount(orderAmount)
	req.SetTimeoutExpress("86400")
	req.SetRequestDate(util.GetNowDatetime())
	req.SetRedirectUrl("")
	req.SetNotifyUrl(notifyUrl)
	req.SetGoodsParamExt(string(goodsParamExtBody))
	req.SetPaymentParamExt("{}")
	req.SetIndustryParamExt("")
	req.SetMemo("")
	req.SetRiskParamExt("")
	req.SetCsUrl("")
	req.SetFundProcessType("REAL_TIME")
	req.SetDivideDetail("")
	req.SetDivideNotifyUrl("")
	return req
}

func (obj ApiStdTradeOrder) GetResponse(req request.StdTradeOrderRequest) (response.StdTradeOrderResponse, error) {
	headers := util.YeepayGetHeaders(req, obj.config, obj.logger)
	url := obj.config.GetGateway() + req.GetMethod()

	var resp response.StdTradeOrderResponse
	httpResp, err := util.Post(url, nil, req.GetBizContent(), headers, obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiStdTradeOrder,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiStdTradeOrder,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiStdTradeOrder,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiStdTradeOrder," + resp.Result.Code + ":" + resp.Result.Message)
	}
	return resp, nil
}
