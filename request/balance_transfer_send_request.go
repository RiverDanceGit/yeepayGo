package request

func NewBalanceTransferSendRequest() BalanceTransferSendRequest {
	req := BalanceTransferSendRequest{}
	req.Init()
	return req
}

type BalanceTransferSendRequest struct {
	httpMethod     string
	method         string
	version        string
	customerNumber string
	groupNumber    string
	batchNo        string
	orderId        string
	amount         string
	product        string
	urgency        string
	accountName    string
	accountNumber  string
	bankCode       string
	bankName       string
	bankBranchName string
	provinceCode   string
	cityCode       string
	feeType        string
	desc           string
	leaveWord      string
	abstractInfo   string
	bizContent     map[string]string
}

func (req *BalanceTransferSendRequest) Init() *BalanceTransferSendRequest {
	req.bizContent = make(map[string]string)
	return req
}

func (req *BalanceTransferSendRequest) SetHttpMethod(httpMethod string) *BalanceTransferSendRequest {
	req.httpMethod = httpMethod
	return req
}

func (req BalanceTransferSendRequest) GetHttpMethod() string {
	return req.httpMethod
}

func (req *BalanceTransferSendRequest) SetMethod(method string) *BalanceTransferSendRequest {
	req.method = method
	return req
}

func (req BalanceTransferSendRequest) GetMethod() string {
	return req.method
}

func (req *BalanceTransferSendRequest) SetVersion(version string) *BalanceTransferSendRequest {
	req.version = version
	return req
}

func (req *BalanceTransferSendRequest) SetCustomerNumber(customerNumber string) *BalanceTransferSendRequest {
	req.customerNumber = customerNumber
	req.bizContent["customerNumber"] = customerNumber
	return req
}

func (req *BalanceTransferSendRequest) SetGroupNumber(groupNumber string) *BalanceTransferSendRequest {
	req.groupNumber = groupNumber
	req.bizContent["groupNumber"] = groupNumber
	return req
}

func (req *BalanceTransferSendRequest) SetBatchNo(batchNo string) *BalanceTransferSendRequest {
	req.batchNo = batchNo
	req.bizContent["batchNo"] = batchNo
	return req
}

func (req *BalanceTransferSendRequest) SetOrderId(orderId string) *BalanceTransferSendRequest {
	req.orderId = orderId
	req.bizContent["orderId"] = orderId
	return req
}

func (req *BalanceTransferSendRequest) SetAmount(amount string) *BalanceTransferSendRequest {
	req.amount = amount
	req.bizContent["amount"] = amount
	return req
}

func (req *BalanceTransferSendRequest) SetProduct(product string) *BalanceTransferSendRequest {
	req.product = product
	req.bizContent["product"] = product
	return req
}

func (req *BalanceTransferSendRequest) SetUrgency(urgency string) *BalanceTransferSendRequest {
	req.urgency = urgency
	req.bizContent["urgency"] = urgency
	return req
}

func (req *BalanceTransferSendRequest) SetAccountName(accountName string) *BalanceTransferSendRequest {
	req.accountName = accountName
	req.bizContent["accountName"] = accountName
	return req
}

func (req *BalanceTransferSendRequest) SetAccountNumber(accountNumber string) *BalanceTransferSendRequest {
	req.accountNumber = accountNumber
	req.bizContent["accountNumber"] = accountNumber
	return req
}

func (req *BalanceTransferSendRequest) SetBankCode(bankCode string) *BalanceTransferSendRequest {
	req.bankCode = bankCode
	req.bizContent["bankCode"] = bankCode
	return req
}

func (req *BalanceTransferSendRequest) SetBankName(bankName string) *BalanceTransferSendRequest {
	req.bankName = bankName
	req.bizContent["bankName"] = bankName
	return req
}

func (req *BalanceTransferSendRequest) SetBankBranchName(bankBranchName string) *BalanceTransferSendRequest {
	req.bankBranchName = bankBranchName
	req.bizContent["bankBranchName"] = bankBranchName
	return req
}

func (req *BalanceTransferSendRequest) SetProvinceCode(provinceCode string) *BalanceTransferSendRequest {
	req.provinceCode = provinceCode
	req.bizContent["provinceCode"] = provinceCode
	return req
}

func (req *BalanceTransferSendRequest) SetCityCode(cityCode string) *BalanceTransferSendRequest {
	req.cityCode = cityCode
	req.bizContent["cityCode"] = cityCode
	return req
}

func (req *BalanceTransferSendRequest) SetFeeType(feeType string) *BalanceTransferSendRequest {
	req.feeType = feeType
	req.bizContent["feeType"] = feeType
	return req
}

func (req *BalanceTransferSendRequest) SetDesc(desc string) *BalanceTransferSendRequest {
	req.desc = desc
	req.bizContent["desc"] = desc
	return req
}

// 留言字数限制
// 农业银行 4
// 建设银行 20
// 中国银行 15
// 广发银行 20
// 中信银行 10
// 光大银行 64
// 浦发银行 30
// 招商银行 20
// 其他 不支持备注
func (req *BalanceTransferSendRequest) SetLeaveWord(leaveWord string) *BalanceTransferSendRequest {
	req.leaveWord = leaveWord
	req.bizContent["leaveWord"] = leaveWord
	return req
}

func (req *BalanceTransferSendRequest) SetAbstractInfo(abstractInfo string) *BalanceTransferSendRequest {
	req.abstractInfo = abstractInfo
	req.bizContent["abstractInfo"] = abstractInfo
	return req
}

func (req BalanceTransferSendRequest) GetBizContent() map[string]string {
	return req.bizContent
}
