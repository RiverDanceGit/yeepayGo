package request

func NewBalanceQueryCustomerAmountRequest() BalanceQueryCustomerAmountRequest {
	req := BalanceQueryCustomerAmountRequest{}
	req.Init()
	return req
}

type BalanceQueryCustomerAmountRequest struct {
	httpMethod     string
	method         string
	version        string
	customerNumber string
	bizContent     map[string]string
}

func (req *BalanceQueryCustomerAmountRequest) Init() *BalanceQueryCustomerAmountRequest {
	req.bizContent = make(map[string]string)
	return req
}

func (req *BalanceQueryCustomerAmountRequest) SetHttpMethod(httpMethod string) *BalanceQueryCustomerAmountRequest {
	req.httpMethod = httpMethod
	return req
}

func (req BalanceQueryCustomerAmountRequest) GetHttpMethod() string {
	return req.httpMethod
}

func (req *BalanceQueryCustomerAmountRequest) SetMethod(method string) *BalanceQueryCustomerAmountRequest {
	req.method = method
	return req
}

func (req BalanceQueryCustomerAmountRequest) GetMethod() string {
	return req.method
}

func (req *BalanceQueryCustomerAmountRequest) SetVersion(version string) *BalanceQueryCustomerAmountRequest {
	req.version = version
	return req
}

func (req *BalanceQueryCustomerAmountRequest) SetCustomerNumber(customerNumber string) *BalanceQueryCustomerAmountRequest {
	req.customerNumber = customerNumber
	req.bizContent["customerNumber"] = customerNumber
	return req
}

func (req BalanceQueryCustomerAmountRequest) GetCustomerNumber() string {
	return req.customerNumber
}

func (req BalanceQueryCustomerAmountRequest) GetBizContent() map[string]string {
	return req.bizContent
}
