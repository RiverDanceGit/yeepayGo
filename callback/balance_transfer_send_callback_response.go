package callback

type BalanceTransferSendCallbackResponse struct {
	CustomerNumber     string      `json:"customerNumber"`
	BatchNo            string      `json:"batchNo"`
	OrderId            string      `json:"orderId"`
	TransferStatusCode string      `json:"transferStatusCode"`
	BankTrxStatusCode  string      `json:"bankTrxStatusCode"`
	SuccessAmount      string      `json:"successAmount"`
	FinishDate         string      `json:"finishDate"`
	FailAmount         interface{} `json:"failAmount"`
	RefundAmount       interface{} `json:"refundAmount"`
	BankMsg            string      `json:"bankMsg"`
}

func (resp BalanceTransferSendCallbackResponse) IsSuccess() bool {
	if "S" == resp.BankTrxStatusCode && "0026" == resp.TransferStatusCode {
		return true
	}
	return false
}
