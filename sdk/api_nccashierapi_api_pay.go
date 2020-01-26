package sdk

import (
	"encoding/json"
	"errors"
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
func (obj ApiNccashierapiApiPay) GetRequest(token, userNo, userType, userIp, appid, openId, payTool, payType string) request.NccashierapiApiPayRequest {
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
	req.SetAppId(appid)
	req.SetOpenId(openId)
	req.SetPayEmpowerNo("")
	req.SetMerchantTerminalId("")
	req.SetMerchantStoreNo("")
	req.SetExtParamMap(string(extParamMapBody))
	return req
}

// 小程序支付
func (obj ApiNccashierapiApiPay) GetMiniAppRequest(token, userNo, userType, userIp, openId string) request.NccashierapiApiPayRequest {
	return obj.GetRequest(token, userNo, userType, userIp, obj.config.GetMiniAppId(), openId, "MINI_PROGRAM", "WECHAT")
}

//公众号 H5 支付
func (obj ApiNccashierapiApiPay) GetWechatOpenRequest(token, userNo, userType, userIp, openId string) request.NccashierapiApiPayRequest {
	return obj.GetRequest(token, userNo, userType, userIp, obj.config.GetMpAppId(), openId, "WECHAT_OPENID", "WECHAT")
}

func (obj ApiNccashierapiApiPay) GetResponse(req request.NccashierapiApiPayRequest) (response.NccashierapiApiPayResponse, error) {
	headers := util.YeepayGetHeaders(req, obj.config, obj.logger)
	url := obj.config.GetGateway() + req.GetMethod()

	var resp response.NccashierapiApiPayResponse
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
		return resp, errors.New("pay fail," + resp.Result.Message)
	}
	return resp, nil
}

func (obj ApiNccashierapiApiPay) GetResultData(resp response.NccashierapiApiPayResponse) (response.NccashierapiApiPayResponseResultData, error) {
	var resultData response.NccashierapiApiPayResponseResultData
	bytes := []byte(resp.Result.ResultData)
	err := json.Unmarshal(bytes, &resultData)
	if err != nil {
		return resultData, err
	}
	return resultData, nil
}