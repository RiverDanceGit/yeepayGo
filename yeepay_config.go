package yeepayGo

func NewYeepayConfig(id, miniAppid, mpAppid, yopPublicKey, privateKeyFile string) YeepayConfig {
	yeepayConfig := YeepayConfig{}
	yeepayConfig.SetSdkLang("go")
	yeepayConfig.SetSdkVersion("3.2.11")
	yeepayConfig.SetGateway("https://openapi.yeepay.com/yop-center")
	yeepayConfig.SetAppKey("OPR:" + id)
	yeepayConfig.SetMerchantNo(id)
	yeepayConfig.SetParentMerchantNo(id)
	yeepayConfig.SetYopPublicKey(yopPublicKey)
	yeepayConfig.SetPrivateKeyFile(privateKeyFile)
	yeepayConfig.SetMiniAppId(miniAppid)
	yeepayConfig.SetMpAppId(mpAppid)
	return yeepayConfig
}

type YeepayConfig struct {
	sdkLang          string
	sdkVersion       string
	gateway          string
	appKey           string
	merchantNo       string
	parentMerchantNo string
	yopPublicKey     string
	privateKey       string
	privateKeyFile   string
	orderNotifyUrl   string
	refundNotifyUrl  string
	vipNotifyUrl     string
	redirectUrl      string
	miniAppId        string
	mpAppId          string
}

func (req *YeepayConfig) SetSdkLang(sdkLang string) *YeepayConfig {
	req.sdkLang = sdkLang
	return req
}

func (req YeepayConfig) GetSdkLang() string {
	return req.sdkLang
}

func (req *YeepayConfig) SetSdkVersion(sdkVersion string) *YeepayConfig {
	req.sdkVersion = sdkVersion
	return req
}

func (req YeepayConfig) GetSdkVersion() string {
	return req.sdkVersion
}

func (req *YeepayConfig) SetGateway(gateway string) *YeepayConfig {
	req.gateway = gateway
	return req
}

func (req YeepayConfig) GetGateway() string {
	return req.gateway
}

func (req *YeepayConfig) SetAppKey(appKey string) *YeepayConfig {
	req.appKey = appKey
	return req
}

func (req YeepayConfig) GetAppKey() string {
	return req.appKey
}

func (req *YeepayConfig) SetMerchantNo(merchantNo string) *YeepayConfig {
	req.merchantNo = merchantNo
	return req
}

func (req YeepayConfig) GetMerchantNo() string {
	return req.merchantNo
}

func (req *YeepayConfig) SetParentMerchantNo(parentMerchantNo string) *YeepayConfig {
	req.parentMerchantNo = parentMerchantNo
	return req
}

func (req YeepayConfig) GetParentMerchantNo() string {
	return req.parentMerchantNo
}

func (req *YeepayConfig) SetYopPublicKey(yopPublicKey string) *YeepayConfig {
	req.yopPublicKey = yopPublicKey
	return req
}

func (req YeepayConfig) GetYopPublicKey() string {
	return req.yopPublicKey
}

func (req *YeepayConfig) SetPrivateKey(privateKey string) *YeepayConfig {
	req.privateKey = privateKey
	return req
}

func (req YeepayConfig) GetPrivateKey() string {
	return req.privateKey
}

func (req *YeepayConfig) SetPrivateKeyFile(privateKeyFile string) *YeepayConfig {
	req.privateKeyFile = privateKeyFile
	return req
}

func (req YeepayConfig) GetPrivateKeyFile() string {
	return req.privateKeyFile
}

func (req *YeepayConfig) SetOrderNotifyUrl(orderNotifyUrl string) *YeepayConfig {
	req.orderNotifyUrl = orderNotifyUrl
	return req
}

func (req YeepayConfig) GetOrderNotifyUrl() string {
	return req.orderNotifyUrl
}

func (req *YeepayConfig) SetRefundNotifyUrl(refundNotifyUrl string) *YeepayConfig {
	req.refundNotifyUrl = refundNotifyUrl
	return req
}

func (req YeepayConfig) GetRefundNotifyUrl() string {
	return req.refundNotifyUrl
}

func (req *YeepayConfig) SetVipNotifyUrl(vipNotifyUrl string) *YeepayConfig {
	req.vipNotifyUrl = vipNotifyUrl
	return req
}

func (req YeepayConfig) GetVipNotifyUrl() string {
	return req.vipNotifyUrl
}

func (req *YeepayConfig) SetMiniAppId(miniAppId string) *YeepayConfig {
	req.miniAppId = miniAppId
	return req
}

func (req YeepayConfig) GetMiniAppId() string {
	return req.miniAppId
}

func (req *YeepayConfig) SetMpAppId(mpAppId string) *YeepayConfig {
	req.mpAppId = mpAppId
	return req
}

func (req YeepayConfig) GetMpAppId() string {
	return req.mpAppId
}
