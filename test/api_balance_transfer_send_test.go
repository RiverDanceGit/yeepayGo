package test

import (
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/sdk"
	"testing"
)

func TestApiBalanceTransferSend(t *testing.T) {
	config, err := GetYeepayConfig()
	if err != nil {
		t.Error(err)
		return
	}
	logger := yeepayGo.NewYeepayLogger()
	apiBalanceTransferSend := sdk.NewApiBalanceTransferSend(config, logger)
	//req := apiBalanceTransferSend.GetRequest("1000000000000001", "TEST_002", "0.01", "马哈哈", "6212261202026355109", "ICBC", "描述描述描述描述描述", "留言留言留言留言留言")
	req := apiBalanceTransferSend.GetRequest("1000000000000011", "TEST_011", "0.01", "韦世介", "6230580000248755808", "SZPA", "描述描述描述描述描述", "留言留言留言留言留言")
	resp, err := apiBalanceTransferSend.GetResponse(req)
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
	t.Log("resp.Result.OrderId", resp.Result.OrderId)
	t.Log("resp.Result.BatchNo", resp.Result.BatchNo)
	t.Log("resp.Result.TransferStatusCode", resp.Result.TransferStatusCode)
	t.Log("resp.Result.Urgency", resp.Result.Urgency)
}
