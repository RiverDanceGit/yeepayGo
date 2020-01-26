package util

import (
	"errors"
	"sort"
	"strings"
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/request"
)

func YeepayGetSortQuery(arr map[string]string) string {
	var params []string
	var query string
	for key, val := range arr {
		if "Authorization" == key {
			continue
		}
		query = key + "=" + RawUrlencode(val)
		params = append(params, query)
	}
	sort.Strings(params)
	return strings.Join(params, "&")
}

func yeepayGetSignedHeaderParams() []string {
	var headers []string
	headers = append(headers, "x-yop-request-id")
	return headers
}

func YeepayGetSignedHeader() string {
	headers := yeepayGetSignedHeaderParams()
	str := strings.Join(headers, ";")
	return Strtolower(str)
}

func YeepayGetHeader(arr map[string]string) string {
	allowed := yeepayGetSignedHeaderParams()

	var headers []string
	header := ""
	for key, val := range arr {
		if StrInArray(key, allowed) {
			header = RawUrlencode(Strtolower(key)) + ":" + RawUrlencode(val)
			headers = append(headers, header)
		}
	}
	sort.Strings(headers)
	return strings.Join(headers, "\n")
}

func YeepayEncrypt(str string, pKeyFile string) string {
	pKey, err := ParsePrivateFromFile(pKeyFile)
	if err != nil {
		return ""
	}
	signer, err := RsaPrivateSha256Encode(str, pKey)
	if err != nil {
		return ""
	}
	return string(signer)
}

func YeepayGetHeaders(req request.YeepayRequest, yeepayConfig yeepayGo.YeepayConfig, yeepayLogger yeepayGo.YeepayLoggerInterface) map[string]string {
	headers := make(map[string]string)
	headers["x-yop-appkey"] = yeepayConfig.GetAppKey()
	headers["x-yop-request-id"] = Uuid()

	protocolVersion := "yop-auth-v2"
	expiredSeconds := "1800"
	timestamp := GetISO8601TimeStamp()
	canonicalURI := req.GetMethod()
	canonicalQueryString := YeepayGetSortQuery(req.GetBizContent())
	canonicalHeader := YeepayGetHeader(headers)
	canonicalRequest := protocolVersion + "/" + yeepayConfig.GetAppKey() + "/" + timestamp + "/" + expiredSeconds + "\n" + req.GetHttpMethod() + "\n" + canonicalURI + "\n" + canonicalQueryString + "\n" + canonicalHeader

	encodeData := YeepayEncrypt(canonicalRequest, yeepayConfig.GetPrivateKeyFile())
	signToBase64 := YeepayBase64Encode([]byte(encodeData), false)

	headers["Authorization"] = "YOP-RSA2048-SHA256 " + protocolVersion + "/" + yeepayConfig.GetAppKey() + "/" + timestamp + "/" + expiredSeconds + "/" + YeepayGetSignedHeader() + "/" + signToBase64 + "$SHA256"
	headers["x-yop-sdk-langs"] = yeepayConfig.GetSdkLang()
	headers["x-yop-sdk-version"] = yeepayConfig.GetSdkVersion()

	if yeepayLogger != nil {
		yeepayLogger.Debug(canonicalQueryString, "|", canonicalHeader, "|", canonicalRequest, "|", headers["Authorization"])
	}
	return headers
}

func YeepayCallback(response string, yeepayConfig yeepayGo.YeepayConfig) (string, string, error) {
	args := strings.Split(response, "$")
	if 4 != len(args) {
		return "", "", errors.New("response has wrong args")
	}
	encryptedRandomKeyToBase64, err := YeepayBase64Decode(args[0])
	if err != nil {
		return "", "", errors.New("YeepayBase64Decode args[0] fail," + err.Error())
	}
	encryptedDataToBase64, err := YeepayBase64Decode(args[1])
	if err != nil {
		return "", "", errors.New("YeepayBase64Decode args[1] fail," + err.Error())
	}
	randomKey, err := RsaPrivateDecode(encryptedRandomKeyToBase64, yeepayConfig.GetPrivateKeyFile())
	if err != nil {
		return "", "", errors.New("RsaPrivateDecode fail," + err.Error())
	}
	encryptedData := Aes128ECBDecode(encryptedDataToBase64, randomKey)
	if 0 == len(encryptedData) {
		return "", "", errors.New("Aes128ECBDecode fail")
	}
	signToBase64 := Substr(Strrchr(string(encryptedData), "$"), 1, len(encryptedData)-1)
	sourceData := Substr(string(encryptedData), 0, len(encryptedData)-len(signToBase64)-1)
	return signToBase64, sourceData, nil
}
