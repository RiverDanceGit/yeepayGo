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

func NewApiNccashierapiApiPay(config yeepayGo.YeepayConfig, logger yeepayGo.YeepayLoggerInterface) ApiNccashierapiApiPay {
	return ApiNccashierapiApiPay{config, logger}
}

type ApiNccashierapiApiPay struct {
	config yeepayGo.YeepayConfig
	logger yeepayGo.YeepayLoggerInterface
}

// 聚合API收银台
// https://open.yeepay.com/docs/e-commerceprotocols/rest__v1.0__nccashierapi__api__pay.html
func (obj ApiNccashierapiApiPay) GetRequest(token, userNo, userType, userIp, appId, openId, payTool, payType string) request.NccashierapiApiPayRequest {
	extParamMap := params.ExtParamMap{"XIANXIA"}
	extParamMapBody, _ := json.Marshal(extParamMap)

	req := request.NewNccashierapiApiPayRequest()
	req.SetHttpMethod(enum.HTTP_METHOD_POST)
	req.SetMethod("/rest/v1.0/nccashierapi/api/pay")
	req.SetVersion("1.0")
	req.SetToken(token)
	req.SetPayTool(payTool)
	req.SetPayType(payType)
	req.SetUserNo(userNo)
	req.SetUserType(userType)
	req.SetUserIp(userIp)
	req.SetAppId(appId)
	req.SetOpenId(openId)
	req.SetPayEmpowerNo("")
	req.SetMerchantTerminalId("")
	req.SetMerchantStoreNo("")
	req.SetExtParamMap(string(extParamMapBody))
	return req
}

func (obj ApiNccashierapiApiPay) GetResponse(req request.NccashierapiApiPayRequest) (response.NccashierapiApiPayResponse, error) {
	headers := util.YeepayGetHeaders(req, obj.config, obj.logger)
	url := obj.config.GetGateway() + req.GetMethod()

	var resp response.NccashierapiApiPayResponse
	httpResp, err := util.Post(url, req.GetBizContent(), nil, headers, obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiNccashierapiApiPay,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiNccashierapiApiPay,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiNccashierapiApiPay,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiNccashierapiApiPay," + resp.Result.Code + ":" + resp.Result.Message)
	}
	return resp, nil
}

func (obj ApiNccashierapiApiPay) GetResultData(resp response.NccashierapiApiPayResponse) (response.NccashierapiApiPayResponseResultData, error) {
	var resultData response.NccashierapiApiPayResponseResultData
	bytes := []byte(resp.Result.ResultData)
	err := json.Unmarshal(bytes, &resultData)
	if err != nil {
		return resultData, util.ErrorWrap(err, "ApiNccashierapiApiPay ResultData,json decode fail")
	}
	return resultData, nil
}
