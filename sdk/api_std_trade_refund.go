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

func NewApiStdTradeRefund(config yeepayGo.YeepayConfig, logger yeepayGo.YeepayLoggerInterface) ApiStdTradeRefund {
	return ApiStdTradeRefund{config, logger}
}

type ApiStdTradeRefund struct {
	config yeepayGo.YeepayConfig
	logger yeepayGo.YeepayLoggerInterface
}

// 退款请求
// https://open.yeepay.com/docs/e-commerceprotocols/rest__v1.0__std__trade__refund.html
func (obj ApiStdTradeRefund) GetRequest(orderId, uniqueOrderNo, refundRequestId, refundAmount, description, memo string) request.StdTradeRefundRequest {
	req := request.NewStdTradeRefundRequest()
	req.SetHttpMethod(enum.HTTP_METHOD_POST)
	req.SetMethod("/rest/v1.0/std/trade/refund")
	req.SetVersion("1.0")
	req.SetParentMerchantNo(obj.config.GetParentMerchantNo())
	req.SetMerchantNo(obj.config.GetMerchantNo())
	req.SetOrderId(orderId)
	req.SetUniqueOrderNo(uniqueOrderNo)
	req.SetRefundRequestId(refundRequestId)
	req.SetRefundAmount(refundAmount)
	req.SetDescription(description)
	req.SetMemo(memo)
	req.SetNotifyUrl(obj.config.GetRefundNotifyUrl())
	req.SetAccountDivided("")
	req.SetRefundStrategy("")
	return req
}

func (obj ApiStdTradeRefund) GetResponse(req request.StdTradeRefundRequest) (response.StdTradeRefundResponse, error) {
	headers := util.YeepayGetHeaders(req, obj.config, obj.logger)
	url := obj.config.GetGateway() + req.GetMethod()

	var resp response.StdTradeRefundResponse
	code, bodyBytes, err := util.Post(url, nil, req.GetBizContent(), headers, obj.logger)
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
		return resp, errors.New("refund fail," + resp.Result.Message)
	}
	return resp, nil
}
