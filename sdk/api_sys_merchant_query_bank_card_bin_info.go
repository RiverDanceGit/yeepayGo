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

func NewApiSysMerchantQueryBankCardBinInfo(config yeepayGo.YeepayConfig, logger yeepayGo.YeepayLoggerInterface) ApiSysMerchantQueryBankCardBinInfo {
	return ApiSysMerchantQueryBankCardBinInfo{config, logger}
}

type ApiSysMerchantQueryBankCardBinInfo struct {
	config yeepayGo.YeepayConfig
	logger yeepayGo.YeepayLoggerInterface
}

// 查询银行卡卡bin信息
// https://open.yeepay.com/docs/retail000001/rest__v1.0__sys__merchant__query-bank-card-bin-info.html
func (obj ApiSysMerchantQueryBankCardBinInfo) GetRequest(bankCardNo string) request.SysMerchantQueryBankCardBinInfoRequest {
	req := request.NewSysMerchantQueryBankCardBinInfoRequest()
	req.SetHttpMethod(enum.HTTP_METHOD_POST)
	req.SetMethod("/rest/v1.0/sys/merchant/query-bank-card-bin-info")
	req.SetVersion("1.0")
	req.SetBankCardNo(bankCardNo)
	return req
}

func (obj ApiSysMerchantQueryBankCardBinInfo) GetResponse(req request.SysMerchantQueryBankCardBinInfoRequest) (response.SysMerchantQueryBankCardBinInfoResponse, error) {
	headers := util.YeepayGetHeaders(req, obj.config, obj.logger)
	url := obj.config.GetGateway() + req.GetMethod()

	var resp response.SysMerchantQueryBankCardBinInfoResponse
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
		return resp, errors.New("bank card query fail," + resp.Result.ReturnMsg)
	}
	return resp, nil
}
