package util

import (
	"encoding/base64"
)

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(data string) ([]byte, error) {
	switch len(data) % 4 {
	case 2:
		data += "=="
	case 3:
		data += "="
	}
	return base64.StdEncoding.DecodeString(data)
}

func YeepayBase64Encode(data []byte, usePadding bool) string {
	if 0 == len(data) {
		return ""
	}
	encoded := Strtr(Base64Encode(data), "+/", "-_")
	if usePadding {
		return encoded
	}
	str := Rtrim(encoded, "=")
	return str
}

func YeepayBase64Decode(data string) ([]byte, error) {
	return Base64Decode(Strtr(data, "-_", "+/"))
}
