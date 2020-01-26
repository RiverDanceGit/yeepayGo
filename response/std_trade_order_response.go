package response

type StdTradeOrderResponse struct {
	Result struct {
		Code               string `json:"code"`
		Message            string `json:"message"`
		ParentMerchantNo   string `json:"parentMerchantNo"`
		MerchantNo         string `json:"merchantNo"`
		OrderId            string `json:"orderId"`
		UniqueOrderNo      string `json:"uniqueOrderNo"`
		GoodsParamExt      string `json:"goodsParamExt"`
		Token              string `json:"token"`
		FundProcessType    string `json:"fundProcessType"`
		ParentMerchantName string `json:"parentMerchantName"`
		MerchantName       string `json:"merchantName"`
	} `json:"result"`
}

func (resp *StdTradeOrderResponse) IsSuccess() bool {
	if "OPR00000" == resp.Result.Code {
		return true
	}
	return false
}

//type StdTradeOrderResponse struct {
//	RequestID  string `json:"requestId"`
//	Code       string `json:"code"`
//	Message    string `json:"message"`
//	SubCode    string `json:"subCode"`
//	SubMessage string `json:"subMessage"`
//}
