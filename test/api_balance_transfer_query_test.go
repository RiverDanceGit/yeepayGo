package test

import (
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/sdk"
	"testing"
)

func TestApiBalanceTransferQuery(t *testing.T) {
	config := GetYeepayConfig()
	logger := yeepayGo.NewYeepayLogger()
	apiBalanceTransferQuery := sdk.NewApiBalanceTransferQuery(config, logger)
	//req := apiBalanceTransferQuery.GetRequest("1000000000000000", "TEST_001", 1, 10)
	req := apiBalanceTransferQuery.GetRequest("1000000000000008", "TEST_008", 1, 10)
	resp, err := apiBalanceTransferQuery.GetResponse(req)
	if err != nil {
		t.Error(err)
		return
	}
	if !resp.IsSuccess() {
		t.Error("resp.Result.ErrorCode", resp.Result.ErrorCode)
		t.Error("resp.Result.ErrorMsg", resp.Result.ErrorMsg)
		return
	}
	t.Log("resp.Result.ErrorCode", resp.Result.ErrorCode)
	t.Log("resp.Result.PageNo", resp.Result.PageNo)
	t.Log("resp.Result.PageSize", resp.Result.PageSize)
	t.Log("resp.Result.TotalPageSize", resp.Result.TotalPageSize)
	t.Log("resp.Result.TotalCount", resp.Result.TotalCount)
	for key, item := range resp.Result.List {
		t.Log("resp.Result.List->", key)
		t.Log("item.OrderId", item.OrderId)
		t.Log("item.BatchNo", item.BatchNo)
		t.Log("item.TransferStatusCode", item.TransferStatusCode)
		t.Log("item.BankTrxStatusCode", item.BankTrxStatusCode)
		t.Log("item.AccountName", item.AccountName)
		t.Log("item.AccountNumber", item.AccountNumber)
		t.Log("item.BankCode", item.BankCode)
		t.Log("item.BankName", item.BankName)
		t.Log("item.Amount", item.Amount)
		t.Log("item.Fee", item.Fee)
		t.Log("item.FeeType", item.FeeType)
		t.Log("item.Urgency", item.Urgency)
		t.Log("item.LeaveWord", item.LeaveWord)
		t.Log()
	}
}
