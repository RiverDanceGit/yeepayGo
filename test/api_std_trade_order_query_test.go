package test

import (
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/sdk"
	"testing"
)

func TestApiStdTradeOrderQuery(t *testing.T) {
	config, err := GetYeepayConfig()
	if err != nil {
		t.Error(err)
		return
	}
	logger := yeepayGo.NewYeepayLogger()
	apiStdTradeOrderQuery := sdk.NewApiStdTradeOrderQuery(config, logger)
	req := apiStdTradeOrderQuery.GetRequest("20191022144904272121988263183551", "")
	resp, err := apiStdTradeOrderQuery.GetResponse(req)
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
	t.Log("resp.Result.UniqueOrderNo", resp.Result.UniqueOrderNo)
	t.Log("resp.Result.Status", resp.Result.Status)
	t.Log("resp.Result.OrderAmount", resp.Result.OrderAmount)
	t.Log("resp.Result.RequestDate", resp.Result.RequestDate)
	t.Log("resp.Result.GoodsParamExt", resp.Result.GoodsParamExt)
	t.Log("resp.Result.Token", resp.Result.Token)
	t.Log("resp.Result.FundProcessType", resp.Result.FundProcessType)
	t.Log("resp.Result.HaveAccounted", resp.Result.HaveAccounted)
	t.Log("resp.Result.ResidualDivideAmount", resp.Result.ResidualDivideAmount)
	t.Log("resp.Result.ParentMerchantName", resp.Result.ParentMerchantName)
	t.Log("resp.Result.MerchantName", resp.Result.MerchantName)
	t.Log("resp.Result.YpSettleAmount", resp.Result.YpSettleAmount)
}
