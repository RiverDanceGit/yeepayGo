package response

type StdTradeRefundQueryResponse struct {
	Result struct {
		Code              string  `json:"code"`
		Message           string  `json:"message"`
		ParentMerchantNo  string  `json:"parentMerchantNo"`
		MerchantNo        string  `json:"merchantNo"`
		OrderId           string  `json:"orderId"`
		RefundRequestId   string  `json:"refundRequestId"`
		UniqueOrderNo     string  `json:"uniqueOrderNo"`
		UniqueRefundNo    string  `json:"uniqueRefundNo"`
		RefundAmount      float64 `json:"refundAmount"`
		Status            string  `json:"status"`
		RefundRequestDate string  `json:"refundRequestDate"`
		DisAccountAmount  float64 `json:"disAccountAmount"`
	} `json:"result"`
}

func (resp *StdTradeRefundQueryResponse) IsSuccess() bool {
	if "OPR00000" == resp.Result.Code {
		return true
	}
	return false
}
