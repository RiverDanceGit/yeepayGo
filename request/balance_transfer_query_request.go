package request

import "strconv"

func NewBalanceTransferQueryRequest() BalanceTransferQueryRequest {
	req := BalanceTransferQueryRequest{}
	req.Init()
	return req
}

type BalanceTransferQueryRequest struct {
	httpMethod     string
	method         string
	version        string
	customerNumber string
	groupNumber    string
	batchNo        string
	orderId        string
	product        string
	queryMode      string
	pageNo         int
	pageSize       int
	bizContent     map[string]string
}

func (req *BalanceTransferQueryRequest) Init() *BalanceTransferQueryRequest {
	req.bizContent = make(map[string]string)
	return req
}

func (req *BalanceTransferQueryRequest) SetHttpMethod(httpMethod string) *BalanceTransferQueryRequest {
	req.httpMethod = httpMethod
	return req
}

func (req BalanceTransferQueryRequest) GetHttpMethod() string {
	return req.httpMethod
}

func (req *BalanceTransferQueryRequest) SetMethod(method string) *BalanceTransferQueryRequest {
	req.method = method
	return req
}

func (req BalanceTransferQueryRequest) GetMethod() string {
	return req.method
}

func (req *BalanceTransferQueryRequest) SetVersion(version string) *BalanceTransferQueryRequest {
	req.version = version
	return req
}

func (req *BalanceTransferQueryRequest) SetCustomerNumber(customerNumber string) *BalanceTransferQueryRequest {
	req.customerNumber = customerNumber
	req.bizContent["customerNumber"] = customerNumber
	return req
}

func (req *BalanceTransferQueryRequest) SetGroupNumber(groupNumber string) *BalanceTransferQueryRequest {
	req.groupNumber = groupNumber
	req.bizContent["groupNumber"] = groupNumber
	return req
}

func (req *BalanceTransferQueryRequest) SetBatchNo(batchNo string) *BalanceTransferQueryRequest {
	req.batchNo = batchNo
	req.bizContent["batchNo"] = batchNo
	return req
}

func (req *BalanceTransferQueryRequest) SetOrderId(orderId string) *BalanceTransferQueryRequest {
	req.orderId = orderId
	req.bizContent["orderId"] = orderId
	return req
}

func (req *BalanceTransferQueryRequest) SetProduct(product string) *BalanceTransferQueryRequest {
	req.product = product
	req.bizContent["product"] = product
	return req
}

func (req *BalanceTransferQueryRequest) SetQueryMode(queryMode string) *BalanceTransferQueryRequest {
	req.queryMode = queryMode
	req.bizContent["queryMode"] = queryMode
	return req
}

func (req *BalanceTransferQueryRequest) SetPageNo(pageNo int) *BalanceTransferQueryRequest {
	req.pageNo = pageNo
	req.bizContent["pageNo"] = strconv.Itoa(pageNo)
	return req
}

func (req *BalanceTransferQueryRequest) SetPageSize(pageSize int) *BalanceTransferQueryRequest {
	req.pageSize = pageSize
	req.bizContent["pageSize"] = strconv.Itoa(pageSize)
	return req
}

func (req BalanceTransferQueryRequest) GetBizContent() map[string]string {
	return req.bizContent
}
