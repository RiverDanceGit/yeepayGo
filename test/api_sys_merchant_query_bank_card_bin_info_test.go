package test

import (
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/sdk"
	"testing"
)

func TestApiSysMerchantQueryBankCardBinInfo(t *testing.T) {
	config, err := GetYeepayConfig()
	if err != nil {
		t.Error(err)
		return
	}
	logger := yeepayGo.NewYeepayLogger()
	cardno := ""
	apiSysMerchantQueryBankCardBinInfo := sdk.NewApiSysMerchantQueryBankCardBinInfo(config, logger)
	req := apiSysMerchantQueryBankCardBinInfo.GetRequest(cardno)
	resp, err := apiSysMerchantQueryBankCardBinInfo.GetResponse(req)
	if err != nil {
		t.Error(err)
		return
	}
	if !resp.IsSuccess() {
		t.Error("resp.Result.ReturnCode", resp.Result.ReturnCode)
		t.Error("resp.Result.ReturnMsg", resp.Result.ReturnMsg)
		return
	}
	t.Log("resp.Result.ReturnCode", resp.Result.ReturnCode)
	t.Log("resp.Result.ReturnMsg", resp.Result.ReturnMsg)
	t.Log("resp.Result.BankID", resp.Result.BankID)
	t.Log("resp.Result.BankCode", resp.Result.BankCode)
	t.Log("resp.Result.BankName", resp.Result.BankName)
	t.Log("resp.Result.CardName", resp.Result.CardName)
	t.Log("resp.Result.CardLength", resp.Result.CardLength)
	t.Log("resp.Result.VerifyLength", resp.Result.VerifyLength)
	t.Log("resp.Result.VerifyCode", resp.Result.VerifyCode)
	t.Log("resp.Result.CardType", resp.Result.CardType)
}
