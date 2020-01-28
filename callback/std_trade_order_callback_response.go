package callback

type StdTradeOrderCallbackResponse struct {
	BankTrxId              string `json:"bankTrxId"`
	PreAuthAccountTime     string `json:"preAuthAccountTime"`
	OrderId                string `json:"orderId"`
	BankOrderId            string `json:"bankOrderId"`
	OpenId                 string `json:"openID"`
	PaySuccessDate         string `json:"paySuccessDate"`
	CustomerFee            string `json:"customerFee"`
	PlatformType           string `json:"platformType"`
	CardType               string `json:"cardType"`
	BankPaySuccessDate     string `json:"bankPaySuccessDate"`
	UniqueOrderNo          string `json:"uniqueOrderNo"`
	YpSettleAmount         string `json:"ypSettleAmount"`
	FundProcessType        string `json:"fundProcessType"`
	BankId                 string `json:"bankId"`
	OrderAmount            string `json:"orderAmount"`
	PayAmount              string `json:"payAmount"`
	RequestDate            string `json:"requestDate"`
	ParentMerchantNo       string `json:"parentMerchantNo"`
	YeepayPromotionDTOList string `json:"yeepayPromotionDTOList"`
	PaymentProduct         string `json:"paymentProduct"`
	MerchantNo             string `json:"merchantNo"`
	Status                 string `json:"status"`
}

func (resp StdTradeOrderCallbackResponse) IsSuccess() bool {
	if "SUCCESS" == resp.Status {
		return true
	}
	return false
}
