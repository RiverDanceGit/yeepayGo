package request

func NewStdTradeRefundQueryRequest() StdTradeRefundQueryRequest {
	req := StdTradeRefundQueryRequest{}
	req.Init()
	return req
}

type StdTradeRefundQueryRequest struct {
	httpMethod       string
	method           string
	version          string
	parentMerchantNo string
	merchantNo       string
	orderId          string
	refundRequestId  string
	uniqueRefundNo   string
	bizContent       map[string]string
}

func (req *StdTradeRefundQueryRequest) Init() *StdTradeRefundQueryRequest {
	req.bizContent = make(map[string]string)
	return req
}

func (req *StdTradeRefundQueryRequest) SetHttpMethod(httpMethod string) *StdTradeRefundQueryRequest {
	req.httpMethod = httpMethod
	return req
}

func (req StdTradeRefundQueryRequest) GetHttpMethod() string {
	return req.httpMethod
}

func (req *StdTradeRefundQueryRequest) SetMethod(method string) *StdTradeRefundQueryRequest {
	req.method = method
	return req
}

func (req StdTradeRefundQueryRequest) GetMethod() string {
	return req.method
}

func (req *StdTradeRefundQueryRequest) SetVersion(version string) *StdTradeRefundQueryRequest {
	req.version = version
	return req
}

func (req *StdTradeRefundQueryRequest) SetParentMerchantNo(parentMerchantNo string) *StdTradeRefundQueryRequest {
	req.parentMerchantNo = parentMerchantNo
	req.bizContent["parentMerchantNo"] = parentMerchantNo
	return req
}

func (req *StdTradeRefundQueryRequest) SetMerchantNo(merchantNo string) *StdTradeRefundQueryRequest {
	req.merchantNo = merchantNo
	req.bizContent["merchantNo"] = merchantNo
	return req
}

func (req *StdTradeRefundQueryRequest) SetOrderId(orderId string) *StdTradeRefundQueryRequest {
	req.orderId = orderId
	req.bizContent["orderId"] = orderId
	return req
}

func (req *StdTradeRefundQueryRequest) SetRefundRequestId(refundRequestId string) *StdTradeRefundQueryRequest {
	req.refundRequestId = refundRequestId
	req.bizContent["refundRequestId"] = refundRequestId
	return req
}

func (req *StdTradeRefundQueryRequest) SetUniqueRefundNo(uniqueRefundNo string) *StdTradeRefundQueryRequest {
	req.uniqueRefundNo = uniqueRefundNo
	req.bizContent["uniqueRefundNo"] = uniqueRefundNo
	return req
}

func (req StdTradeRefundQueryRequest) GetBizContent() map[string]string {
	return req.bizContent
}
