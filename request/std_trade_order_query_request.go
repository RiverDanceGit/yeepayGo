package request

func NewStdTradeOrderQueryRequest() StdTradeOrderQueryRequest {
	req := StdTradeOrderQueryRequest{}
	req.Init()
	return req
}

type StdTradeOrderQueryRequest struct {
	httpMethod       string
	method           string
	version          string
	parentMerchantNo string
	merchantNo       string
	orderId          string
	uniqueOrderNo    string
	bizContent       map[string]string
}

func (req *StdTradeOrderQueryRequest) Init() *StdTradeOrderQueryRequest {
	req.bizContent = make(map[string]string)
	return req
}

func (req *StdTradeOrderQueryRequest) SetHttpMethod(httpMethod string) *StdTradeOrderQueryRequest {
	req.httpMethod = httpMethod
	return req
}

func (req StdTradeOrderQueryRequest) GetHttpMethod() string {
	return req.httpMethod
}

func (req *StdTradeOrderQueryRequest) SetMethod(method string) *StdTradeOrderQueryRequest {
	req.method = method
	return req
}

func (req StdTradeOrderQueryRequest) GetMethod() string {
	return req.method
}

func (req *StdTradeOrderQueryRequest) SetVersion(version string) *StdTradeOrderQueryRequest {
	req.version = version
	return req
}

func (req *StdTradeOrderQueryRequest) SetParentMerchantNo(parentMerchantNo string) *StdTradeOrderQueryRequest {
	req.parentMerchantNo = parentMerchantNo
	req.bizContent["parentMerchantNo"] = parentMerchantNo
	return req
}

func (req *StdTradeOrderQueryRequest) SetMerchantNo(merchantNo string) *StdTradeOrderQueryRequest {
	req.merchantNo = merchantNo
	req.bizContent["merchantNo"] = merchantNo
	return req
}

func (req *StdTradeOrderQueryRequest) SetOrderId(orderId string) *StdTradeOrderQueryRequest {
	req.orderId = orderId
	req.bizContent["orderId"] = orderId
	return req
}

func (req *StdTradeOrderQueryRequest) SetUniqueOrderNo(uniqueOrderNo string) *StdTradeOrderQueryRequest {
	req.uniqueOrderNo = uniqueOrderNo
	req.bizContent["uniqueOrderNo"] = uniqueOrderNo
	return req
}

func (req StdTradeOrderQueryRequest) GetBizContent() map[string]string {
	return req.bizContent
}
