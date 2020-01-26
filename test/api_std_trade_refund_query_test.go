package test

import (
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/sdk"
	"github.com/RiverDanceGit/yeepayGo/util"
	"testing"
)

func TestApiStdTradeRefundQuery(t *testing.T) {
	config := GetYeepayConfig()
	logger := yeepayGo.NewYeepayLogger()
	apiStdTradeRefundQuery := sdk.NewApiStdTradeRefundQuery(config, logger)
	req := apiStdTradeRefundQuery.GetRequest("", util.Uuid(), "")
	resp, err := apiStdTradeRefundQuery.GetResponse(req)
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
	t.Log("resp.Result.UniqueOrderNo", resp.Result.UniqueOrderNo)
	t.Log("resp.Result.UniqueRefundNo", resp.Result.UniqueRefundNo)
	t.Log("resp.Result.RefundAmount", resp.Result.RefundAmount)
	t.Log("resp.Result.Status", resp.Result.Status)
	t.Log("resp.Result.RefundRequestDate", resp.Result.RefundRequestDate)
	t.Log("resp.Result.DisAccountAmount", resp.Result.DisAccountAmount)
}
