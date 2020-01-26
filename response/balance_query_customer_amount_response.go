package response

type BalanceQueryCustomerAmountResponse struct {
	Result struct {
		ErrorCode       string  `json:"errorCode"`
		ErrorMsg        string  `json:"errorMsg"`
		CustomerNumber  string  `json:"customerNumber"`
		AccountAmount   float64 `json:"accountAmount"`
		RjtValidAmount  float64 `json:"rjtValidAmount"`
		WtjsValidAmount float64 `json:"wtjsValidAmount"`
	} `json:"result"`
}

func (resp *BalanceQueryCustomerAmountResponse) IsSuccess() bool {
	if "BAC001" == resp.Result.ErrorCode {
		return true
	}
	return false
}
