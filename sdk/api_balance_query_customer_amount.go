package sdk

import (
	"encoding/json"
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/enum"
	"github.com/RiverDanceGit/yeepayGo/request"
	"github.com/RiverDanceGit/yeepayGo/response"
	"github.com/RiverDanceGit/yeepayGo/util"
	"strconv"
)

func NewApiBalanceQueryCustomerAmount(config yeepayGo.YeepayConfig, logger yeepayGo.YeepayLoggerInterface) ApiBalanceQueryCustomerAmount {
	return ApiBalanceQueryCustomerAmount{config, logger}
}

type ApiBalanceQueryCustomerAmount struct {
	config yeepayGo.YeepayConfig
	logger yeepayGo.YeepayLoggerInterface
}

// 代付代发--可用余额查询
// https://open.yeepay.com/docs/remit001/rest__v1.0__balance__query_customer_amount.html
func (obj ApiBalanceQueryCustomerAmount) GetRequest() request.BalanceQueryCustomerAmountRequest {
	req := request.NewBalanceQueryCustomerAmountRequest()
	req.SetHttpMethod(enum.HTTP_METHOD_POST)
	req.SetMethod("/rest/v1.0/balance/query_customer_amount")
	req.SetVersion("1.0")
	req.SetCustomerNumber(obj.config.GetMerchantNo())
	return req
}

func (obj ApiBalanceQueryCustomerAmount) GetResponse(req request.BalanceQueryCustomerAmountRequest) (response.BalanceQueryCustomerAmountResponse, error) {
	headers := util.YeepayGetHeaders(req, obj.config, obj.logger)
	url := obj.config.GetGateway() + req.GetMethod()

	var resp response.BalanceQueryCustomerAmountResponse
	httpResp, err := util.Post(url, req.GetBizContent(), nil, headers, obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiBalanceQueryCustomerAmount,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiBalanceQueryCustomerAmount,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiBalanceQueryCustomerAmount,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiBalanceQueryCustomerAmount," + resp.Result.ErrorCode + ":" + resp.Result.ErrorMsg)
	}
	return resp, nil
}
