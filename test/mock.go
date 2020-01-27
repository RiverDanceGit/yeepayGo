package test

import (
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/util"
)

func GetPrivateKeyFile() string {
	return "../cert/cfca_private_key.pem"
}

func getYopPublicKey() (string, error) {
	return util.FileGetContents("../cert/cfca_public_key.txt")
}

func GetYeepayConfig() (cfg yeepayGo.YeepayConfig, err error) {
	yopPublicKey, err := getYopPublicKey()
	if err != nil {
		return cfg, err
	}
	privateFile := GetPrivateKeyFile()
	cfg = yeepayGo.NewYeepayConfig("10028946649", yopPublicKey, privateFile)
	return cfg, nil
}
