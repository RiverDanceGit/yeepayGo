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

func NewApiStdTradeRefundQuery(config yeepayGo.YeepayConfig, logger yeepayGo.YeepayLoggerInterface) ApiStdTradeRefundQuery {
	return ApiStdTradeRefundQuery{config, logger}
}

type ApiStdTradeRefundQuery struct {
	config yeepayGo.YeepayConfig
	logger yeepayGo.YeepayLoggerInterface
}

// 退款查询
// https://open.yeepay.com/docs/e-commerceprotocols/rest__v1.0__std__trade__refundquery.html
func (obj ApiStdTradeRefundQuery) GetRequest(orderId, refundRequestId, uniqueRefundNo string) request.StdTradeRefundQueryRequest {
	req := request.NewStdTradeRefundQueryRequest()
	req.SetHttpMethod(enum.HTTP_METHOD_POST)
	req.SetMethod("/rest/v1.0/std/trade/refundquery")
	req.SetVersion("1.0")
	req.SetParentMerchantNo(obj.config.GetMerchantNo())
	req.SetMerchantNo(obj.config.GetMerchantNo())
	req.SetOrderId(orderId)
	req.SetRefundRequestId(refundRequestId)
	if "" != uniqueRefundNo {
		req.SetUniqueRefundNo(uniqueRefundNo)
	}
	return req
}

func (obj ApiStdTradeRefundQuery) GetResponse(req request.StdTradeRefundQueryRequest) (response.StdTradeRefundQueryResponse, error) {
	headers := util.YeepayGetHeaders(req, obj.config, obj.logger)
	url := obj.config.GetGateway() + req.GetMethod()

	var resp response.StdTradeRefundQueryResponse
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
		return resp, errors.New("refund query fail," + resp.Result.Message)
	}
	return resp, nil
}
