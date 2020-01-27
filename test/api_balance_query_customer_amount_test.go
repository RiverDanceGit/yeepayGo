package test

import (
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/sdk"
	"testing"
)

func TestApiBalanceQueryCustomerAmount(t *testing.T) {
	config, err := GetYeepayConfig()
	if err != nil {
		t.Error(err)
		return
	}
	logger := yeepayGo.NewYeepayLogger()
	apiBalanceQueryCustomerAmount := sdk.NewApiBalanceQueryCustomerAmount(config, logger)
	req := apiBalanceQueryCustomerAmount.GetRequest()
	resp, err := apiBalanceQueryCustomerAmount.GetResponse(req)
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
	t.Log("resp.Result.CustomerNumber", resp.Result.CustomerNumber)   // 商户编号
	t.Log("resp.Result.AccountAmount", resp.Result.AccountAmount)     // 账户可用余额
	t.Log("resp.Result.RjtValidAmount", resp.Result.RjtValidAmount)   // 日结通可用余额
	t.Log("resp.Result.WtjsValidAmount", resp.Result.WtjsValidAmount) // 代付代发可用余额
}
