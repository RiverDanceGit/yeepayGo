package sdk

import (
	"encoding/json"
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/enum"
	"github.com/RiverDanceGit/yeepayGo/request"
	"github.com/RiverDanceGit/yeepayGo/response"
	"github.com/RiverDanceGit/yeepayGo/util"
	"strconv"
)

func NewApiBalanceTransferSend(config yeepayGo.YeepayConfig, logger yeepayGo.YeepayLoggerInterface) ApiBalanceTransferSend {
	return ApiBalanceTransferSend{config, logger}
}

type ApiBalanceTransferSend struct {
	config yeepayGo.YeepayConfig
	logger yeepayGo.YeepayLoggerInterface
}

// 代付代发--单笔出款请求
// https://open.yeepay.com/docs/remit001/rest__v1.0__balance__transfer_send.html
func (obj ApiBalanceTransferSend) GetRequest(batchNo, orderId, amount, accountName, accountNumber, bankCode, desc, leaveWord string) request.BalanceTransferSendRequest {
	req := request.NewBalanceTransferSendRequest()
	req.SetHttpMethod(enum.HTTP_METHOD_POST)
	req.SetMethod("/rest/v1.0/balance/transfer_send")
	req.SetVersion("1.0")
	req.SetCustomerNumber(obj.config.GetMerchantNo())
	req.SetGroupNumber(obj.config.GetMerchantNo())
	req.SetBatchNo(batchNo) //15-20 位，仅数字
	req.SetOrderId(orderId)
	req.SetAmount(amount)
	req.SetProduct("WTJS")
	req.SetUrgency("1")
	req.SetAccountName(accountName)
	req.SetAccountNumber(accountNumber)
	req.SetBankCode(bankCode)
	req.SetBankName("")
	req.SetBankBranchName("")
	req.SetProvinceCode("")
	req.SetCityCode("")
	req.SetFeeType("SOURCE")
	req.SetDesc(desc)
	req.SetLeaveWord(leaveWord)
	req.SetAbstractInfo("")
	return req
}

// 一分钟内，一张银行卡号只能发起一次请求
func (obj ApiBalanceTransferSend) GetResponse(req request.BalanceTransferSendRequest) (response.BalanceTransferSendResponse, error) {
	headers := util.YeepayGetHeaders(req, obj.config, obj.logger)
	url := obj.config.GetGateway() + req.GetMethod()

	var resp response.BalanceTransferSendResponse
	httpResp, err := util.Post(url, nil, req.GetBizContent(), headers, obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiBalanceTransferSend,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiBalanceTransferSend,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiBalanceTransferSend,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiBalanceTransferSend," + resp.Result.ErrorCode + ":" + resp.Result.ErrorMsg)
	}
	return resp, nil
}
