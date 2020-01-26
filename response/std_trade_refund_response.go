package response

type StdTradeRefundResponse struct {
	Result struct {
		Code             string `json:"code"`
		Message          string `json:"message"`
		ParentMerchantNo string `json:"parentMerchantNo"`
		MerchantNo       string `json:"merchantNo"`
		OrderId          string `json:"orderId"`
		RefundRequestId  string `json:"refundRequestId"`
		RefundAmount     string `json:"refundAmount"`
		Description      string `json:"description"`
	} `json:"result"`
}

func (resp *StdTradeRefundResponse) IsSuccess() bool {
	if "OPR00000" == resp.Result.Code {
		return true
	}
	return false
}
