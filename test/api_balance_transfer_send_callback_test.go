package test

import (
	"encoding/json"
	"github.com/RiverDanceGit/yeepayGo/callback"
	"github.com/RiverDanceGit/yeepayGo/util"
	"strings"
	"testing"
)

func TestApiBalanceTransferSendCallbackResponse(t *testing.T) {
	sourceData := ""
	sourceData = `{
"customerNumber": "10028946649",
"batchNo": "1000000000000008",
"orderId": "TEST_008",
"transferStatusCode": "0026",
"bankTrxStatusCode": "S",
"finishDate": "2019-10-31 12:30:51",
"successAmount": "",
"failAmount": "",
"refundAmount": "",
"bankMsg": ""
}`
	printApiBalanceTransferSendCallback(t, sourceData)
	sourceData = `{
"customerNumber": "10028946649",
"batchNo": "1000000000000006",
"orderId": "TEST_007",
"transferStatusCode": "0028",
"bankTrxStatusCode": "W",
"finishDate": "2019-10-30 17:41:42",
"refundAmount": 0.0100,
"bankMsg": "可用打款余额不足",
"leaveWord": "留言留言留言留言留言"
}`
	printApiBalanceTransferSendCallback(t, sourceData)
	sourceData = `{
"customerNumber": "10028946649",
"batchNo": "277518088889501312",
"orderId": "W277518088889435776",
"transferStatusCode": "0026",
"bankTrxStatusCode": "F",
"finishDate": "2019-11-28 20:18:25",
"failAmount": 496.0000,
"bankMsg": "您输入的身份验证信息有误，请确认后重试[1020305]",
"leaveWord": ""
}`
	printApiBalanceTransferSendCallback(t, sourceData)
}

func TestApiBalanceTransferSendCallback(t *testing.T) {
	response := ""
	response = "Z4BBIB2W9q4tIIJFBjKGkp9bfxCPAlC7LLWgSffxvTH9pP7y8h0sLOrfkL2jkjalhsR4UdB2gFmHukbvGsfMyl_UvwL7_4JoCIkUQhH9EYquzcPCO7YCZzmo7J73C0sseFQ12vgElZZaQY-r8-VHib3lunq4haLGsxjDDjUx5weL3EULBwsCsAXDZU6J6Ozv5JmQN_wguIPj__2C6saPVVMGJ-vRmr0Qm2ioSt0inuzKEnmQk_XR0geQRRwWl_NFOQLAokWrUAa5yryBCMv0zDmdoEW9kOSJluzpJTjn8wCpSHBrNkRR_v64VkpZ-_VhUHotOI0_Dg1CukmCb5fsgw$j-wRDsTwFxGWeBxImUAfhdJVQlcprcgwjFoGlpN_nysUI712fbRgwjiSPMLgE-KtBsRQBOlgcZ07maK6iM0_gmxQXR2lsD1Pcq2TbT5IZBp86uAUl98TTiD27CDYvLyk9a9DUFjvAW8Jg4tJ2rFAotWH7USKLkfOwa5pLQqMSegcAQmQasUVb9mWaDaVjhTrSXtthvUXXBkycZwwtNhby9XbsB4nU_lwKqGxRRsIKIV_-Ziy4M1IZJMu96gi80UQzb8IEgkQmX5cdS7UnUcqSeflvZksJAebps3ZLLaHTgMiE3JICXN7PsaNCRRZNj2Swi7RUr22Eupa_56iEdckEK--T0j9reLcXNQMKdu0JYt-hMBTZovzYI-ttQ4xLyiGVlwqBXrWr54mf4rUxh5a67jwfmW8C5D5SI8Q3PC0Es4ag0pJnKotnlgRaPWJofofAMd01Zry45hSsFJZc5QZRqmSI3M8CdeZGLxsAJqZfuqhgPhbgwOqul04d2Q07-JJA9wjusuI_hCbbG-wLqE69INJ0CE2GuFynxtzo6rLzuKOVtvo-DwuJWJB_ob0ngOqfvBnIUqp_CAEJlt-c0qwgmrXKH2hQZjiPfE1HY_P6dW6fC03E8aTr8ViO1zIPKRkvzQQ2Cb_DswrRqI_HTFVZdshby0jeY3Ucxdw6ao1kA-zti9UsXI9eZTlu9U2A-_nhTEAhXtPgggZfhkTpllvHaIhRYNGJvJUIShnGMI3N3RucYogknn-gijCk2Z0cbV6sQMrezZKF4XHpykfciouyxt8TfKoLgRvifri-Ad0hGj-9Vef2GcS2SOR6SPkedb8$AES$SHA256"
	response = "qlNOAzHDV52XcJqTZ_fdhn-uVZxZiBul5jq6Ixw7tn8BZbxmBFeRzLIBSuoSW25M3zSSseVzuu0Ph-C3nMbzIeQ3DXH89-d8ccTbI6gSWAKi6GCHQKL25zKaBSaIYrhcqXcAhxVsR6R8ZGNBvXaMEQ9cfFQAe0yc_swrgvSk-D-Y_2vrRoq32inA2fj9cUyobXyoQuZREnDwdmovYjrKD9iOVOHKawA49QCvMfk_CwmJfmJEFbrgm0OB4hU-VZ6-zWnrgKOoM2HTYigdwcfhUaOdoVA0LQJ8WOZTXamwdIRuTj6dYdjPQyk-NbSFvT9rz48Hazs9-k4NsKHXR5FFwA$sZZjeoOsVcp_UaWicgSyM0N-y_-KHPPAwQRZ4y8PC4QCF4TYF6w0MBYnvHa3Vu-0YiFP5dH-lkvrhNDsgOUInv2HA9lWblzpXJZpzNAUUbiexgptaieNSMbolcevSKp-tyeyLywYLGCxJ8Jf19AGCnri7fO3LxDfnFKoPVSkEaa-qcNzvqfwOXgU41IYZFkQXXZFuf_iKyF6FSJteqiApju4Wd6_1lYgjQRT1txQ7M9Ui7mtD3zknP2dhXK4lfg5N_z2BF5OTcJ9qi-1nF1IbKWovC6JOutQf_zG07i0xmUpvUsLDtGtnQarrDLQgxdB1aeXUYW7zCRs_kO-y8IJL7qtfpczOnfZ736IWrV6sXBXYwkA_kDF2i926kTF_GJ3vwSpXw5r_whvg-SNFCuVHof2cNEnAH0YOM6J20Dj78qIkFxQT_fTyLZbWgXjOVhEoLtlwY3o5us-pVu1M2XNEdoZC7XGkGXPDm-tXXWCkQEVpV5yrOEtHdgeEukDgEUmIXI51065SMMSBOMHnWPIPBkhbdVtUeb4fbZQqxdwZr31KlXfHk5VAaMKdmaO_FPg2qMnELSejVc3Z0Yt0Oab_HDDgJYNIHul5xTSAsNTu-2hDsjaMS0742IsQ2yTC72qU3ltxxsA_bse7WwPR_fUV3khyLgqmJTqEG6x_1wip6FFiv0xl_gYDdwZYpt6APD_glWmlejAfHE2HWHgkUuBtORwb4Uo5127zx7JeFfcRvBuzHPprOvP8RwdlwpeTnf2HIucoqgvGj0WA4FPnCXsX_CdgCkV-QEjlkDJ6XSu6jM$AES$SHA256"
	balanceTransferSendCallback := callback.BalanceTransferSendCallback{
		response,
		"OPR:10028946649",
	}

	args := strings.Split(balanceTransferSendCallback.Response, "$")
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
	printApiBalanceTransferSendCallback(t, sourceData)
}

func printApiBalanceTransferSendCallback(t *testing.T, sourceData string) {
	var resp callback.BalanceTransferSendCallbackResponse
	err := json.Unmarshal([]byte(sourceData), &resp)
	if err != nil {
		t.Error(err)
		return
	}
	if !resp.IsSuccess() {
		t.Error("resp.BankTrxStatusCode", resp.BankTrxStatusCode)
		t.Error("resp.TransferStatusCode", resp.TransferStatusCode)
		t.Error(sourceData)
		return
	}

	t.Log("resp.TransferStatusCode", resp.TransferStatusCode)
	t.Log("resp.BankTrxStatusCode", resp.BankTrxStatusCode)
	t.Log("resp.CustomerNumber", resp.CustomerNumber)
	t.Log("resp.BatchNo", resp.BatchNo)
	t.Log("resp.OrderId", resp.OrderId)
	t.Log("resp.SuccessAmount", resp.SuccessAmount)
	t.Log("resp.FinishDate", resp.FinishDate)
	t.Log("resp.FailAmount", resp.FailAmount)
	t.Log("resp.RefundAmount", resp.RefundAmount)
	t.Log("resp.BankMsg", resp.BankMsg)
}
