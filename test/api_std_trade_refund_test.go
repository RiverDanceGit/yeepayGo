package test

import (
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/sdk"
	"testing"
)

func TestApiStdTradeRefund(t *testing.T) {
	config := GetYeepayConfig()
	logger := yeepayGo.NewYeepayLogger()
	apiStdTradeRefund := sdk.NewApiStdTradeRefund(config, logger)
	//req := apiStdTradeRefund.GetRequest("20191022144904272121988263183551", "1001201910220000001214000713", util.Uuid(), "0.01", "退款说明", "备注")
	//req := apiStdTradeRefund.GetRequest("20191028170337273005262917469375", "1001201910280000001229080435", "TEST_1", "0.01", "退款说明", "备注")
	req := apiStdTradeRefund.GetRequest("20191106124518274283857384178879", "1001201911060000001253553715", "TEST_2", "0.01", "退款说明", "备注", "")
	resp, err := apiStdTradeRefund.GetResponse(req)
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
	t.Log("resp.Result.ParentMerchantNo", resp.Result.ParentMerchantNo)
	t.Log("resp.Result.MerchantNo", resp.Result.MerchantNo)
	t.Log("resp.Result.OrderId", resp.Result.OrderId)
	t.Log("resp.Result.RefundRequestId", resp.Result.RefundRequestId)
	t.Log("resp.Result.RefundAmount", resp.Result.RefundAmount)
	t.Log("resp.Result.Description", resp.Result.Description)
}
