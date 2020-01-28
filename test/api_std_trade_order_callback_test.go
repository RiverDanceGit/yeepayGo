package test

import (
	"encoding/json"
	"github.com/RiverDanceGit/yeepayGo/callback"
	"github.com/RiverDanceGit/yeepayGo/util"
	"strings"
	"testing"
)

func TestApiStdTradeOrderCallback(t *testing.T) {
	response := ""
	response = "Qg954y5vMlwmFh0qMyBEeq8HjXXgBkj7irGhhCnzRNTRwRV4Pt205EoFcsu5a1NF8bwMocaWFdEg1-DWZfNKqmRt5WqmJPEDBIWjGs8a1bhzlzEsfQkyM2SG3GlP3_eVDjWD459X5Hw0-tCOHZmuhgYdkCkvpArwZyVjp0qoqNQU55VObM6xgIhIXujnpF7Kj_Uwu-4kxAOwVa2wbpqiTAZe1UNvH6DhK5ve975rfBX6KKm1IAjFMlzn4rYTWTVypP5XmsBrKsX7yBBjYp1dxtjCpWBhvCsiqC_EDEe3sIyrK3a8eszO3xVNfGBd6UMZfeGYnhp1lm4jPV4Z4qerfg$ZZ0VdBKsPfhD3cK8_twnzO46MgIyyTIGVech5TYvo3Cv5UUG7-IdYEfvFGe_bT-hwQPfykuynfyO-VQREtDaL3KbQ6sNiV0eJcWkwjRs6zZ5PtgCRd16OBvEHPk6eZ4ToJs0HVjjJhqLeZFjVg8f3ryKItKl90_o1Fo39_03EwfvgEgrDRzkmMRWhq28UG52IIK3Fce2dPnthzHFfFEIB2dsrhOasJd3eutKde0gpA5HRv8VJgy9Pcqm26wJHXNvBtWhKVPGrUrPD6WOuECvZg6MT6MZWQSrWsxIAwJ0M3HkXX9biUh0M0EOpmoyYvn0TvJRPOEEhjAUX6uPVLwnb8ttb95OEnq8xYgjA7A3bAu5Pe5Tnx1F-hXSYrHPMu2obn_40UFB_Po5aizEP9QGw5rFyYu7slgzKB5bzJXd_nLjaBDo8WreV-TKry5UF7_81DWvz-tQeLBJosCfhrPjY5sumWqFxJRIcLtYC2rEagj5GJuMtwtjghj_AVPKObmBUsrT_KD48fjINKnMXQmmedhG9ANJIoO9Asr4s6ntyKK3Z69EOSVbdcDBt9SUx0cPLFknjvAzRyaB9PQJs9EQHuE3N6PxKvG2M5GJfX2zgnZnpOS6wcFCPIuCKqeQfE9ZX7RUWjONNNWzRwpT9QhKLprFyYu7slgzKB5bzJXd_nLO_CEl9Z-9Ic6YeYNBpIZ1IPNKSItP6Wf-H5saDiqiqlcN3LLPm2ZiLsnmJpJESq_CDhLHAxyP4h4wMHCRZbZ-6PEVY-J0dFqBrw7-NsEgSPJ81LuMdrzX4o3dDE8mzRUFMsbWuCBIZRfTOcO07q0AHcmcgleuK6rLVTieKLfig1lDJqP8yicUZGcGClenv7BERSt5dvIX3ne6YlD0SjpOgO3RkBxzYUU8tLW5aM-RduruOdfZiD-En_rmRk-nm3EPYl5SxQ5P9pvMILLNuqWI6MyORUto-ZRQJ2A52ZVm0sjjvynyKZ-mK13GKnYL-18L0z03Swj8tfQiAHPP7WIKf1twqG1h6TIgcH8wkBRJGdNulTClM5xNALq_eeRCTS0vVlD57iHYpC8V64iS-5g08trzB_x1KXeHiZFn0tevbiMzZ3B6fOsBiarefDiT7PbRqy-nyClioV1uYJO5d-zTIsuv12__0Jdt2rI1_XndBwGIbVp-MxOsPUbG_DWq2cjuaOGrA0PjePCshfSm1_MOI6Nhqvf7SK-p0PngcD0EY5XVtyPcAqYiwe0rNJl9Q1kFJEiFsEXhfZtsaTLL14pwNaKFic0Ikvgv82n-VCgE6VzjfiagoxDfpQcpleJ9tRArUXq9O8M4cERsZvAE9zFoTAvmBstWD56zdnay6qibaQ$AES$SHA256"
	stdTradeOrderCallback := callback.StdTradeOrderCallback{
		response,
		"OPR:10028946649",
	}

	args := strings.Split(stdTradeOrderCallback.Response, "$")
	t.Log("args[0]", args[0])
	t.Log("args[1]", args[1])
	t.Log("args[2]", args[2])
	t.Log("args[3]", args[3])

	encryptedRandomKeyToBase64, _ := util.YeepayBase64Decode(args[0])
	encryptedDataToBase64, _ := util.YeepayBase64Decode(args[1])
	//symmetricEncryptAlg := args[2]
	//digestAlg := args[3]

	randomKey, err := util.RsaPrivateDecode(encryptedRandomKeyToBase64, GetPrivateKeyFile())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("randomKey", string(randomKey))

	encryptedData := util.Aes128ECBDecode(encryptedDataToBase64, randomKey)
	t.Log("encryptedData", string(encryptedData))

	signToBase64 := util.Substr(util.Strrchr(string(encryptedData), "$"), 1, len(encryptedData)-1)
	sourceData := util.Substr(string(encryptedData), 0, len(encryptedData)-len(signToBase64)-1)

	var resp callback.StdTradeOrderCallbackResponse
	err = json.Unmarshal([]byte(sourceData), &resp)
	if err != nil {
		t.Error(err)
		return
	}
	if !resp.IsSuccess() {
		t.Error("resp.Status", resp.Status)
		t.Error(sourceData)
		return
	}

	t.Log("resp.Status", resp.Status)
	t.Log("resp.BankTrxId", resp.BankTrxId)
	t.Log("resp.PreAuthAccountTime", resp.PreAuthAccountTime)
	t.Log("resp.OrderId", resp.OrderId)
	t.Log("resp.BankOrderId", resp.BankOrderId)
	t.Log("resp.OpenId", resp.OpenId)
	t.Log("resp.PaySuccessDate", resp.PaySuccessDate)
	t.Log("resp.CustomerFee", resp.CustomerFee)
	t.Log("resp.PlatformType", resp.PlatformType)
	t.Log("resp.CardType", resp.CardType)
	t.Log("resp.BankPaySuccessDate", resp.BankPaySuccessDate)
	t.Log("resp.UniqueOrderNo", resp.UniqueOrderNo)
	t.Log("resp.YpSettleAmount", resp.YpSettleAmount)
	t.Log("resp.FundProcessType", resp.FundProcessType)
	t.Log("resp.BankId", resp.BankId)
	t.Log("resp.OrderAmount", resp.OrderAmount)
	t.Log("resp.PayAmount", resp.PayAmount)
	t.Log("resp.RequestDate", resp.RequestDate)
	t.Log("resp.ParentMerchantNo", resp.ParentMerchantNo)
	t.Log("resp.YeepayPromotionDTOList", resp.YeepayPromotionDTOList)
	t.Log("resp.PaymentProduct", resp.PaymentProduct)
	t.Log("resp.MerchantNo", resp.MerchantNo)
}
