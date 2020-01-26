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

func NewApiBalanceTransferQuery(config yeepayGo.YeepayConfig, logger yeepayGo.YeepayLoggerInterface) ApiBalanceTransferQuery {
	return ApiBalanceTransferQuery{config, logger}
}

type ApiBalanceTransferQuery struct {
	config yeepayGo.YeepayConfig
	logger yeepayGo.YeepayLoggerInterface
}

// 代付代发--出款查询
// https://open.yeepay.com/docs/remit001/rest__v1.0__balance__transfer_query.html
func (obj ApiBalanceTransferQuery) GetRequest(batchNo, orderId string, pageNo, pageSize int) request.BalanceTransferQueryRequest {
	req := request.NewBalanceTransferQueryRequest()
	req.SetHttpMethod(enum.HTTP_METHOD_POST)
	req.SetMethod("/rest/v1.0/balance/transfer_query")
	req.SetVersion("1.0")
	req.SetCustomerNumber(obj.config.GetMerchantNo())
	req.SetGroupNumber(obj.config.GetMerchantNo())
	req.SetBatchNo(batchNo)
	req.SetOrderId(orderId)
	req.SetProduct("WTJS")
	req.SetQueryMode("1")
	req.SetPageNo(pageNo)
	req.SetPageSize(pageSize)
	return req
}

func (obj ApiBalanceTransferQuery) GetResponse(req request.BalanceTransferQueryRequest) (response.BalanceTransferQueryResponse, error) {
	headers := util.YeepayGetHeaders(req, obj.config, obj.logger)
	url := obj.config.GetGateway() + req.GetMethod()

	var resp response.BalanceTransferQueryResponse
	httpResp, err := util.Post(url, req.GetBizContent(), nil, headers, obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiBalanceTransferQuery,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiBalanceTransferQuery,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiBalanceTransferQuery,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiBalanceTransferQuery," + resp.Result.ErrorCode + ":" + resp.Result.ErrorMsg)
	}
	return resp, nil
}
