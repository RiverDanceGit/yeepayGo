package test

import "github.com/RiverDanceGit/yeepayGo"

func GetPrivateKeyFile() string {
	return ""
}

func GetYeepayConfig() yeepayGo.YeepayConfig {
	return yeepayGo.NewYeepayConfig("", "", "", "", GetPrivateKeyFile())
}
