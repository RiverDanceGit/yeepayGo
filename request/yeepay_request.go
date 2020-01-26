package request

type YeepayRequest interface {
	GetMethod() string
	GetHttpMethod() string
	GetBizContent() map[string]string
}
