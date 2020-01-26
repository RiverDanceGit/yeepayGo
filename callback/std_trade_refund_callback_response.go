package callback

type StdTradeRefundCallbackResponse struct {
	DisAccountAmount             string `json:"disAccountAmount"`
	ResidualAmount               string `json:"residualAmount"`
	OrderId                      string `json:"orderId"`
	AccountDivided               string `json:"accountDivided"`
	BizSystemNo                  string `json:"bizSystemNo"`
	Description                  string `json:"description"`
	UniqueOrderNo                string `json:"uniqueOrderNo"`
	RefundRequestId              string `json:"refundRequestId"`
	RefundSuccessDate            string `json:"refundSuccessDate"`
	YeepayRefundPromotionDTOList string `json:"yeepayRefundPromotionDTOList"`
	UniqueRefundNo               string `json:"uniqueRefundNo"`
	ParentMerchantNo             string `json:"parentMerchantNo"`
	RefundRequestDate            string `json:"refundRequestDate"`
	MerchantNo                   string `json:"merchantNo"`
	RefundAmount                 string `json:"refundAmount"`
	Status                       string `json:"status"`
}
