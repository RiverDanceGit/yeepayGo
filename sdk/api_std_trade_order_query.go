package sdk

import (
	"encoding/json"
	"errors"
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/enum"
	"github.com/RiverDanceGit/yeepayGo/request"
	"github.com/RiverDanceGit/yeepayGo/response"
	"github.com/RiverDanceGit/yeepayGo/util"
	"strconv"
)

func NewApiStdTradeOrderQuery(config yeepayGo.YeepayConfig, logger yeepayGo.YeepayLoggerInterface) ApiStdTradeOrderQuery {
	return ApiStdTradeOrderQuery{config, logger}
}

type ApiStdTradeOrderQuery struct {
	config yeepayGo.YeepayConfig
	logger yeepayGo.YeepayLoggerInterface
}

// 查询订单
// https://open.yeepay.com/docs/e-commerceprotocols/rest__v1.0__std__trade__orderquery.html
func (obj ApiStdTradeOrderQuery) GetRequest(orderId, uniqueOrderNo string) request.StdTradeOrderQueryRequest {
	req := request.NewStdTradeOrderQueryRequest()
	req.SetHttpMethod(enum.HTTP_METHOD_POST)
	req.SetMethod("/rest/v1.0/std/trade/orderquery")
	req.SetVersion("1.0")
	req.SetParentMerchantNo(obj.config.GetMerchantNo())
	req.SetMerchantNo(obj.config.GetMerchantNo())
	req.SetOrderId(orderId)
	if "" != uniqueOrderNo {
		req.SetUniqueOrderNo(uniqueOrderNo)
	}
	return req
}

func (obj ApiStdTradeOrderQuery) GetResponse(req request.StdTradeOrderQueryRequest) (response.StdTradeOrderQueryResponse, error) {
	headers := util.YeepayGetHeaders(req, obj.config, obj.logger)
	url := obj.config.GetGateway() + req.GetMethod()

	var resp response.StdTradeOrderQueryResponse
	code, bodyBytes, err := util.Post(url, req.GetBizContent(), nil, headers, obj.logger)
	if err != nil {
		return resp, err
	} else if code != 200 {
		return resp, errors.New(req.GetMethod() + " code:" + strconv.Itoa(code))
	}
	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil {
		return resp, err
	}
	if !resp.IsSuccess() {
		return resp, errors.New("order query fail," + resp.Result.Message)
	}
	return resp, nil
}
