package response

type SysMerchantQueryBankCardBinInfoResponse struct {
	Result struct {
		ReturnMsg    string `json:"returnMsg"`
		ReturnCode   string `json:"returnCode"`
		BankID       string `json:"bankId"`
		BankCode     string `json:"bankCode"`
		BankName     string `json:"bankName"`
		CardName     string `json:"cardName"`
		CardLength   int    `json:"cardLength"`
		VerifyLength int    `json:"verifyLength"`
		VerifyCode   string `json:"verifyCode"`
		CardType     string `json:"cardType"`
	} `json:"result"`
}

func (resp *SysMerchantQueryBankCardBinInfoResponse) IsSuccess() bool {
	if "REG00000" == resp.Result.ReturnCode {
		return true
	}
	return false
}
