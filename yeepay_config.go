package yeepayGo

func NewYeepayConfig(id, yopPublicKey, privateKeyFile string) YeepayConfig {
	yeepayConfig := YeepayConfig{}
	yeepayConfig.SetSdkLang("go")
	yeepayConfig.SetSdkVersion("3.2.11")
	yeepayConfig.SetGateway("https://openapi.yeepay.com/yop-center")
	yeepayConfig.SetAppKey("OPR:" + id)
	yeepayConfig.SetMerchantNo(id)
	yeepayConfig.SetParentMerchantNo(id)
	yeepayConfig.SetYopPublicKey(yopPublicKey)
	yeepayConfig.SetPrivateKeyFile(privateKeyFile)
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
	privateKeyFile   string
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

func (req *YeepayConfig) SetPrivateKeyFile(privateKeyFile string) *YeepayConfig {
	req.privateKeyFile = privateKeyFile
	return req
}

func (req YeepayConfig) GetPrivateKeyFile() string {
	return req.privateKeyFile
}
