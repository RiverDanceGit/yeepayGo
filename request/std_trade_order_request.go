package request

func NewStdTradeOrderRequest() StdTradeOrderRequest {
	req := StdTradeOrderRequest{}
	req.Init()
	return req
}

type StdTradeOrderRequest struct {
	httpMethod       string
	method           string
	version          string
	parentMerchantNo string
	merchantNo       string
	orderId          string //订单号
	orderAmount      string //金额，单位:元
	timeoutExpress   string //订单有效期
	requestDate      string
	redirectUrl      string
	notifyUrl        string
	goodsParamExt    string
	paymentParamExt  string
	industryParamExt string
	memo             string
	riskParamExt     string
	csUrl            string
	fundProcessType  string
	divideDetail     string
	divideNotifyUrl  string
	bizContent       map[string]string
}

func (req *StdTradeOrderRequest) Init() *StdTradeOrderRequest {
	req.bizContent = make(map[string]string)
	return req
}

func (req *StdTradeOrderRequest) SetHttpMethod(httpMethod string) *StdTradeOrderRequest {
	req.httpMethod = httpMethod
	return req
}

func (req StdTradeOrderRequest) GetHttpMethod() string {
	return req.httpMethod
}

func (req *StdTradeOrderRequest) SetMethod(method string) *StdTradeOrderRequest {
	req.method = method
	return req
}

func (req StdTradeOrderRequest) GetMethod() string {
	return req.method
}

func (req *StdTradeOrderRequest) SetVersion(version string) *StdTradeOrderRequest {
	req.version = version
	return req
}

func (req *StdTradeOrderRequest) SetParentMerchantNo(parentMerchantNo string) *StdTradeOrderRequest {
	req.parentMerchantNo = parentMerchantNo
	req.bizContent["parentMerchantNo"] = parentMerchantNo
	return req
}

func (req *StdTradeOrderRequest) SetMerchantNo(merchantNo string) *StdTradeOrderRequest {
	req.merchantNo = merchantNo
	req.bizContent["merchantNo"] = merchantNo
	return req
}

func (req *StdTradeOrderRequest) SetOrderId(orderId string) *StdTradeOrderRequest {
	req.orderId = orderId
	req.bizContent["orderId"] = orderId
	return req
}

func (req *StdTradeOrderRequest) SetOrderAmount(orderAmount string) *StdTradeOrderRequest {
	req.orderAmount = orderAmount
	req.bizContent["orderAmount"] = orderAmount
	return req
}

func (req *StdTradeOrderRequest) SetTimeoutExpress(timeoutExpress string) *StdTradeOrderRequest {
	req.timeoutExpress = timeoutExpress
	req.bizContent["timeoutExpress"] = timeoutExpress
	return req
}

func (req *StdTradeOrderRequest) SetRequestDate(requestDate string) *StdTradeOrderRequest {
	req.requestDate = requestDate
	req.bizContent["requestDate"] = requestDate
	return req
}

func (req *StdTradeOrderRequest) SetRedirectUrl(redirectUrl string) *StdTradeOrderRequest {
	req.redirectUrl = redirectUrl
	req.bizContent["redirectUrl"] = redirectUrl
	return req
}

func (req *StdTradeOrderRequest) SetNotifyUrl(notifyUrl string) *StdTradeOrderRequest {
	req.notifyUrl = notifyUrl
	req.bizContent["notifyUrl"] = notifyUrl
	return req
}

func (req *StdTradeOrderRequest) SetGoodsParamExt(goodsParamExt string) *StdTradeOrderRequest {
	req.goodsParamExt = goodsParamExt
	req.bizContent["goodsParamExt"] = goodsParamExt
	return req
}

func (req *StdTradeOrderRequest) SetPaymentParamExt(paymentParamExt string) *StdTradeOrderRequest {
	req.paymentParamExt = paymentParamExt
	req.bizContent["paymentParamExt"] = paymentParamExt
	return req
}

func (req *StdTradeOrderRequest) SetIndustryParamExt(industryParamExt string) *StdTradeOrderRequest {
	req.industryParamExt = industryParamExt
	req.bizContent["industryParamExt"] = industryParamExt
	return req
}

func (req *StdTradeOrderRequest) SetMemo(memo string) *StdTradeOrderRequest {
	req.memo = memo
	req.bizContent["memo"] = memo
	return req
}

func (req *StdTradeOrderRequest) SetRiskParamExt(riskParamExt string) *StdTradeOrderRequest {
	req.riskParamExt = riskParamExt
	req.bizContent["riskParamExt"] = riskParamExt
	return req
}

func (req *StdTradeOrderRequest) SetCsUrl(csUrl string) *StdTradeOrderRequest {
	req.csUrl = csUrl
	req.bizContent["csUrl"] = csUrl
	return req
}

func (req *StdTradeOrderRequest) SetFundProcessType(fundProcessType string) *StdTradeOrderRequest {
	req.fundProcessType = fundProcessType
	req.bizContent["fundProcessType"] = fundProcessType
	return req
}

func (req *StdTradeOrderRequest) SetDivideDetail(divideDetail string) *StdTradeOrderRequest {
	req.divideDetail = divideDetail
	req.bizContent["divideDetail"] = divideDetail
	return req
}

func (req *StdTradeOrderRequest) SetDivideNotifyUrl(divideNotifyUrl string) *StdTradeOrderRequest {
	req.divideNotifyUrl = divideNotifyUrl
	req.bizContent["divideNotifyUrl"] = divideNotifyUrl
	return req
}

func (req StdTradeOrderRequest) GetBizContent() map[string]string {
	return req.bizContent
}
