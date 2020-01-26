package response

type BalanceTransferQueryResponse struct {
	Result struct {
		ErrorCode     string `json:"errorCode"`
		ErrorMsg      string `json:"errorMsg"`
		PageNo        int    `json:"pageNo"`
		PageSize      int    `json:"pageSize"`
		TotalPageSize int    `json:"totalPageSize"`
		TotalCount    int    `json:"totalCount"`
		List          []struct {
			OrderId            string  `json:"orderId"`
			BatchNo            string  `json:"batchNo"`
			TransferStatusCode string  `json:"transferStatusCode"`
			BankTrxStatusCode  string  `json:"bankTrxStatusCode"`
			AccountName        string  `json:"accountName"`
			AccountNumber      string  `json:"accountNumber"`
			BankCode           string  `json:"bankCode"`
			BankName           string  `json:"bankName"`
			Amount             float64 `json:"amount"`
			Fee                float64 `json:"fee"`
			FeeType            string  `json:"feeType"`
			Urgency            bool    `json:"urgency"`
			LeaveWord          string  `json:"leaveWord"`
		} `json:"list"`
		ExtInfos struct {
			TOTALSRCFEE         float64 `json:"TOTALSRCFEE"`
			REFUNDEDCOUNT       int     `json:"REFUNDEDCOUNT"`
			TOTALSRCAMOUNT      float64 `json:"TOTALSRCAMOUNT"`
			TOTALTARGETAMOUNT   float64 `json:"TOTALTARGETAMOUNT"`
			GENERATEBATCHAMOUNT float64 `json:"GENERATEBATCHAMOUNT"`
			TOTALTARGETFEE      float64 `json:"TOTALTARGETFEE"`
			GENERATEBATCHCOUNT  int     `json:"GENERATEBATCHCOUNT"`
			REFUNDEDAMOUNT      float64 `json:"REFUNDEDAMOUNT"`
			TOTALCOUNT          int     `json:"TOTALCOUNT"`
		} `json:"extInfos"`
	} `json:"result"`
}

func (resp *BalanceTransferQueryResponse) IsSuccess() bool {
	if "BAC001" == resp.Result.ErrorCode {
		return true
	}
	return false
}
