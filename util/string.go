package util

import (
	"fmt"
	"net/url"
	"strings"
	"time"
	"unicode"

	uuid "github.com/satori/go.uuid"
)

func Uuid() string {
	return uuid.NewV4().String()
}

func Uniqid(prefix string) string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", prefix, now.Unix(), now.UnixNano()%0x100000)
}

func Urlencode(str string) string {
	return url.QueryEscape(str)
}

func RawUrlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

func Strtolower(str string) string {
	return strings.ToLower(str)
}

func Strtr(haystack string, params ...interface{}) string {
	ac := len(params)
	if ac == 1 {
		pairs := params[0].(map[string]string)
		length := len(pairs)
		if length == 0 {
			return haystack
		}
		oldnew := make([]string, length*2)
		for o, n := range pairs {
			if o == "" {
				return haystack
			}
			oldnew = append(oldnew, o, n)
		}
		return strings.NewReplacer(oldnew...).Replace(haystack)
	} else if ac == 2 {
		from := params[0].(string)
		to := params[1].(string)
		trlen, lt := len(from), len(to)
		if trlen > lt {
			trlen = lt
		}
		if trlen == 0 {
			return haystack
		}

		str := make([]uint8, len(haystack))
		var xlat [256]uint8
		var i int
		var j uint8
		if trlen == 1 {
			for i = 0; i < len(haystack); i++ {
				if haystack[i] == from[0] {
					str[i] = to[0]
				} else {
					str[i] = haystack[i]
				}
			}
			return string(str)
		}
		// trlen != 1
		for {
			xlat[j] = j
			if j++; j == 0 {
				break
			}
		}
		for i = 0; i < trlen; i++ {
			xlat[from[i]] = to[i]
		}
		for i = 0; i < len(haystack); i++ {
			str[i] = xlat[haystack[i]]
		}
		return string(str)
	}

	return haystack
}

func Rtrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimRightFunc(str, unicode.IsSpace)
	}
	return strings.TrimRight(str, characterMask[0])
}

func Substr(str string, start uint, length int) string {
	if start < 0 || length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}

func Strrpos(haystack, needle string, offset int) int {
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}
	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = strings.LastIndex(haystack, needle)
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

func Strrchr(haystack string, needle string) string {
	start := Strrpos(haystack, needle, 0)
	if start < 0 {
		return ""
	}
	end := len(haystack)
	return haystack[start:end]
}
