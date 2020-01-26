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
	httpResp, err := util.Post(url, req.GetBizContent(), nil, headers, obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiSysMerchantQueryBankCardBinInfo,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiSysMerchantQueryBankCardBinInfo,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiSysMerchantQueryBankCardBinInfo,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiSysMerchantQueryBankCardBinInfo," + resp.Result.ReturnCode + ":" + resp.Result.ReturnMsg)
	}
	return resp, nil
}
