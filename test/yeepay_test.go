package test

import (
	"github.com/RiverDanceGit/yeepayGo/util"
	"testing"
)

func TestYeepayEncrypt(t *testing.T) {
	str := util.YeepayEncrypt("abc", GetPrivateKeyFile())
	t.Log(str)
	t.Log(util.Base64Encode([]byte(str)))
}
