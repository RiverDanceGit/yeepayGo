package test

import (
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/sdk"
	"testing"
)

func TestApiNccashierapiApiPayMiniApp(t *testing.T) {
	token := "E58FC33398FC02962822FDBCBCF7E8DAD0BA0B34C73A1FE8D16C07119EE38A82"
	config := GetYeepayConfig()
	logger := yeepayGo.NewYeepayLogger()
	apiNccashierapiApiPay := sdk.NewApiNccashierapiApiPay(config, logger)
	req := apiNccashierapiApiPay.GetMiniAppRequest(token, "1", "1", "115.216.124.169", "")
	resp, err := apiNccashierapiApiPay.GetResponse(req)
	if err != nil {
		t.Error(err)
		return
	}
	if !resp.IsSuccess() {
		t.Error("resp.Result.Code", resp.Result.Code)
		t.Error("resp.Result.Message", resp.Result.Message)
		return
	}
	t.Log("resp.Result.Code", resp.Result.Code)
	t.Log("resp.Result.Message", resp.Result.Message)
	t.Log("resp.Result.PayTool", resp.Result.PayTool)
	t.Log("resp.Result.PayType", resp.Result.PayType)
	t.Log("resp.Result.MerchantNo", resp.Result.MerchantNo)
	t.Log("resp.Result.Token", resp.Result.Token)
	t.Log("resp.Result.ResultType", resp.Result.ResultType)
	t.Log("resp.Result.ResultData", resp.Result.ResultData)

	resultData, err := apiNccashierapiApiPay.GetResultData(resp)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("resultData.AppId", resultData.AppId)
	t.Log("resultData.NonceStr", resultData.NonceStr)
	t.Log("resultData.Package", resultData.Package)
	t.Log("resultData.PaySign", resultData.PaySign)
	t.Log("resultData.SignType", resultData.SignType)
	t.Log("resultData.TimeStamp", resultData.TimeStamp)
}

func TestApiNccashierapiApiPayWechatOpen(t *testing.T) {
	token := "08598D98A74241D26C609860A17036EBEEEAC27DE70A0078F1E1F06978456BC2"
	config := GetYeepayConfig()
	logger := yeepayGo.NewYeepayLogger()
	apiNccashierapiApiPay := sdk.NewApiNccashierapiApiPay(config, logger)
	req := apiNccashierapiApiPay.GetWechatOpenRequest(token, "1", "1", "115.216.124.169", "")
	resp, err := apiNccashierapiApiPay.GetResponse(req)
	if err != nil {
		t.Error(err)
		return
	}
	if !resp.IsSuccess() {
		t.Error("resp.Result.Code", resp.Result.Code)
		t.Error("resp.Result.Message", resp.Result.Message)
		return
	}
	t.Log("resp.Result.Code", resp.Result.Code)
	t.Log("resp.Result.Message", resp.Result.Message)
	t.Log("resp.Result.PayTool", resp.Result.PayTool)
	t.Log("resp.Result.PayType", resp.Result.PayType)
	t.Log("resp.Result.MerchantNo", resp.Result.MerchantNo)
	t.Log("resp.Result.Token", resp.Result.Token)
	t.Log("resp.Result.ResultType", resp.Result.ResultType)
	t.Log("resp.Result.ResultData", resp.Result.ResultData)

	resultData, err := apiNccashierapiApiPay.GetResultData(resp)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("resultData.AppId", resultData.AppId)
	t.Log("resultData.NonceStr", resultData.NonceStr)
	t.Log("resultData.Package", resultData.Package)
	t.Log("resultData.PaySign", resultData.PaySign)
	t.Log("resultData.SignType", resultData.SignType)
	t.Log("resultData.TimeStamp", resultData.TimeStamp)
}
