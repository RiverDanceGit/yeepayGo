package response

type BalanceTransferSendResponse struct {
	Result struct {
		ErrorCode          string `json:"errorCode"`
		ErrorMsg           string `json:"errorMsg"`
		OrderId            string `json:"orderId"`
		BatchNo            string `json:"batchNo"`
		TransferStatusCode string `json:"transferStatusCode"`
		Urgency            bool   `json:"urgency"`
	} `json:"result"`
}

func (resp *BalanceTransferSendResponse) IsSuccess() bool {
	if "BAC001" == resp.Result.ErrorCode {
		return true
	}
	return false
}
