package request

func NewNccashierapiApiPayRequest() NccashierapiApiPayRequest {
	req := NccashierapiApiPayRequest{}
	req.Init()
	return req
}

type NccashierapiApiPayRequest struct {
	httpMethod         string
	method             string
	version            string
	token              string
	payTool            string
	payType            string
	userNo             string
	userType           string
	appId              string
	openId             string
	payEmpowerNo       string
	merchantTerminalId string
	merchantStoreNo    string
	userIp             string
	extParamMap        string
	bizContent         map[string]string
}

func (req *NccashierapiApiPayRequest) Init() *NccashierapiApiPayRequest {
	req.bizContent = make(map[string]string)
	return req
}

func (req *NccashierapiApiPayRequest) SetHttpMethod(httpMethod string) *NccashierapiApiPayRequest {
	req.httpMethod = httpMethod
	return req
}

func (req NccashierapiApiPayRequest) GetHttpMethod() string {
	return req.httpMethod
}

func (req *NccashierapiApiPayRequest) SetMethod(method string) *NccashierapiApiPayRequest {
	req.method = method
	return req
}

func (req NccashierapiApiPayRequest) GetMethod() string {
	return req.method
}

func (req *NccashierapiApiPayRequest) SetVersion(version string) *NccashierapiApiPayRequest {
	req.version = version
	return req
}

func (req *NccashierapiApiPayRequest) SetToken(token string) *NccashierapiApiPayRequest {
	req.token = token
	req.bizContent["token"] = token
	return req
}

func (req *NccashierapiApiPayRequest) SetPayTool(payTool string) *NccashierapiApiPayRequest {
	req.payTool = payTool
	req.bizContent["payTool"] = payTool
	return req
}

func (req *NccashierapiApiPayRequest) SetPayType(payType string) *NccashierapiApiPayRequest {
	req.payType = payType
	req.bizContent["payType"] = payType
	return req
}

func (req *NccashierapiApiPayRequest) SetUserNo(userNo string) *NccashierapiApiPayRequest {
	req.userNo = userNo
	req.bizContent["userNo"] = userNo
	return req
}

func (req *NccashierapiApiPayRequest) SetUserType(userType string) *NccashierapiApiPayRequest {
	req.userType = userType
	req.bizContent["userType"] = userType
	return req
}

func (req *NccashierapiApiPayRequest) SetAppId(appId string) *NccashierapiApiPayRequest {
	req.appId = appId
	req.bizContent["appId"] = appId
	return req
}

func (req *NccashierapiApiPayRequest) SetOpenId(openId string) *NccashierapiApiPayRequest {
	req.openId = openId
	req.bizContent["openId"] = openId
	return req
}

func (req *NccashierapiApiPayRequest) SetPayEmpowerNo(payEmpowerNo string) *NccashierapiApiPayRequest {
	req.payEmpowerNo = payEmpowerNo
	req.bizContent["payEmpowerNo"] = payEmpowerNo
	return req
}

func (req *NccashierapiApiPayRequest) SetMerchantTerminalId(merchantTerminalId string) *NccashierapiApiPayRequest {
	req.merchantTerminalId = merchantTerminalId
	req.bizContent["merchantTerminalId"] = merchantTerminalId
	return req
}

func (req *NccashierapiApiPayRequest) SetMerchantStoreNo(merchantStoreNo string) *NccashierapiApiPayRequest {
	req.merchantStoreNo = merchantStoreNo
	req.bizContent["merchantStoreNo"] = merchantStoreNo
	return req
}

func (req *NccashierapiApiPayRequest) SetUserIp(userIp string) *NccashierapiApiPayRequest {
	req.userIp = userIp
	req.bizContent["userIp"] = userIp
	return req
}

func (req *NccashierapiApiPayRequest) SetExtParamMap(extParamMap string) *NccashierapiApiPayRequest {
	req.extParamMap = extParamMap
	req.bizContent["extParamMap"] = extParamMap
	return req
}

func (req NccashierapiApiPayRequest) GetBizContent() map[string]string {
	return req.bizContent
}
