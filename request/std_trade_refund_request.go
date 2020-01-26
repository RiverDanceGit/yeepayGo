package request

func NewStdTradeRefundRequest() StdTradeRefundRequest {
	req := StdTradeRefundRequest{}
	req.Init()
	return req
}

type StdTradeRefundRequest struct {
	httpMethod       string
	method           string
	version          string
	parentMerchantNo string
	merchantNo       string
	orderId          string
	refundRequestId  string
	uniqueOrderNo    string
	refundAmount     string
	description      string
	memo             string
	notifyUrl        string
	accountDivided   string
	refundStrategy   string
	bizContent       map[string]string
}

func (req *StdTradeRefundRequest) Init() *StdTradeRefundRequest {
	req.bizContent = make(map[string]string)
	return req
}

func (req *StdTradeRefundRequest) SetHttpMethod(httpMethod string) *StdTradeRefundRequest {
	req.httpMethod = httpMethod
	return req
}

func (req StdTradeRefundRequest) GetHttpMethod() string {
	return req.httpMethod
}

func (req *StdTradeRefundRequest) SetMethod(method string) *StdTradeRefundRequest {
	req.method = method
	return req
}

func (req StdTradeRefundRequest) GetMethod() string {
	return req.method
}

func (req *StdTradeRefundRequest) SetVersion(version string) *StdTradeRefundRequest {
	req.version = version
	return req
}

func (req *StdTradeRefundRequest) SetParentMerchantNo(parentMerchantNo string) *StdTradeRefundRequest {
	req.parentMerchantNo = parentMerchantNo
	req.bizContent["parentMerchantNo"] = parentMerchantNo
	return req
}

func (req *StdTradeRefundRequest) SetMerchantNo(merchantNo string) *StdTradeRefundRequest {
	req.merchantNo = merchantNo
	req.bizContent["merchantNo"] = merchantNo
	return req
}

func (req *StdTradeRefundRequest) SetOrderId(orderId string) *StdTradeRefundRequest {
	req.orderId = orderId
	req.bizContent["orderId"] = orderId
	return req
}

func (req *StdTradeRefundRequest) SetRefundRequestId(refundRequestId string) *StdTradeRefundRequest {
	req.refundRequestId = refundRequestId
	req.bizContent["refundRequestId"] = refundRequestId
	return req
}

func (req *StdTradeRefundRequest) SetUniqueOrderNo(uniqueOrderNo string) *StdTradeRefundRequest {
	req.uniqueOrderNo = uniqueOrderNo
	req.bizContent["uniqueOrderNo"] = uniqueOrderNo
	return req
}

func (req *StdTradeRefundRequest) SetRefundAmount(refundAmount string) *StdTradeRefundRequest {
	req.refundAmount = refundAmount
	req.bizContent["refundAmount"] = refundAmount
	return req
}

func (req *StdTradeRefundRequest) SetDescription(description string) *StdTradeRefundRequest {
	req.description = description
	req.bizContent["description"] = description
	return req
}

func (req *StdTradeRefundRequest) SetMemo(memo string) *StdTradeRefundRequest {
	req.memo = memo
	req.bizContent["memo"] = memo
	return req
}

func (req *StdTradeRefundRequest) SetNotifyUrl(notifyUrl string) *StdTradeRefundRequest {
	req.notifyUrl = notifyUrl
	req.bizContent["notifyUrl"] = notifyUrl
	return req
}

func (req *StdTradeRefundRequest) SetAccountDivided(accountDivided string) *StdTradeRefundRequest {
	req.accountDivided = accountDivided
	req.bizContent["accountDivided"] = accountDivided
	return req
}

func (req *StdTradeRefundRequest) SetRefundStrategy(refundStrategy string) *StdTradeRefundRequest {
	req.refundStrategy = refundStrategy
	req.bizContent["refundStrategy"] = refundStrategy
	return req
}

func (req StdTradeRefundRequest) GetBizContent() map[string]string {
	return req.bizContent
}
