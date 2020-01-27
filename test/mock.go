package test

import (
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/util"
)

func GetPrivateKeyFile() (string, error) {
	return util.FileGetContents("../cert/cfca_private_key.pem")
}

func getYopPublicKey() (string, error) {
	return util.FileGetContents("../cert/cfca_public_key.txt")
}

func GetYeepayConfig() (cfg yeepayGo.YeepayConfig, err error) {
	privFile, err := GetPrivateKeyFile()
	if err != nil {
		return cfg, err
	}
	pubKey, err := GetPrivateKeyFile()
	if err != nil {
		return cfg, err
	}
	cfg = yeepayGo.NewYeepayConfig("10028946649", pubKey, privFile)
	return cfg, nil
}
