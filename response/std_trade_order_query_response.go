package response

type StdTradeOrderQueryResponse struct {
	Result struct {
		Code                 string  `json:"code"`
		Message              string  `json:"message"`
		ParentMerchantNo     string  `json:"parentMerchantNo"`
		MerchantNo           string  `json:"merchantNo"`
		OrderId              string  `json:"orderId"`
		UniqueOrderNo        string  `json:"uniqueOrderNo"`
		Status               string  `json:"status"`
		OrderAmount          float64 `json:"orderAmount"`
		RequestDate          string  `json:"requestDate"`
		GoodsParamExt        string  `json:"goodsParamExt"`
		Token                string  `json:"token"`
		FundProcessType      string  `json:"fundProcessType"`
		HaveAccounted        bool    `json:"haveAccounted"`
		ResidualDivideAmount string  `json:"residualDivideAmount"`
		ParentMerchantName   string  `json:"parentMerchantName"`
		MerchantName         string  `json:"merchantName"`
		YpSettleAmount       float64 `json:"ypSettleAmount"`
	} `json:"result"`
}

func (resp *StdTradeOrderQueryResponse) IsSuccess() bool {
	if "OPR00000" == resp.Result.Code {
		return true
	}
	return false
}
