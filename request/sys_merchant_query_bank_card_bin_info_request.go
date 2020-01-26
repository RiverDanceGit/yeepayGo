package request

func NewSysMerchantQueryBankCardBinInfoRequest() SysMerchantQueryBankCardBinInfoRequest {
	req := SysMerchantQueryBankCardBinInfoRequest{}
	req.Init()
	return req
}

type SysMerchantQueryBankCardBinInfoRequest struct {
	httpMethod string
	method     string
	version    string
	bankCardNo string
	bizContent map[string]string
}

func (req *SysMerchantQueryBankCardBinInfoRequest) Init() *SysMerchantQueryBankCardBinInfoRequest {
	req.bizContent = make(map[string]string)
	return req
}

func (req *SysMerchantQueryBankCardBinInfoRequest) SetHttpMethod(httpMethod string) *SysMerchantQueryBankCardBinInfoRequest {
	req.httpMethod = httpMethod
	return req
}

func (req SysMerchantQueryBankCardBinInfoRequest) GetHttpMethod() string {
	return req.httpMethod
}

func (req *SysMerchantQueryBankCardBinInfoRequest) SetMethod(method string) *SysMerchantQueryBankCardBinInfoRequest {
	req.method = method
	return req
}

func (req SysMerchantQueryBankCardBinInfoRequest) GetMethod() string {
	return req.method
}

func (req *SysMerchantQueryBankCardBinInfoRequest) SetVersion(version string) *SysMerchantQueryBankCardBinInfoRequest {
	req.version = version
	return req
}

func (req *SysMerchantQueryBankCardBinInfoRequest) SetBankCardNo(bankCardNo string) *SysMerchantQueryBankCardBinInfoRequest {
	req.bankCardNo = bankCardNo
	req.bizContent["bankCardNo"] = bankCardNo
	return req
}

func (req SysMerchantQueryBankCardBinInfoRequest) GetBankCardNo() string {
	return req.bankCardNo
}

func (req SysMerchantQueryBankCardBinInfoRequest) GetBizContent() map[string]string {
	return req.bizContent
}
