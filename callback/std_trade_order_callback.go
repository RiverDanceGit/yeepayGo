package callback

import (
	"encoding/json"
	"errors"
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/util"
)

type StdTradeOrderCallback struct {
	Response               string `form:"response" binding:"required"`
	CustomerIdentification string `form:"customerIdentification" binding:"required"`
}

func OrderCallback(response string, yeepayConfig yeepayGo.YeepayConfig) (StdTradeOrderCallbackResponse, error) {
	var resp StdTradeOrderCallbackResponse

	_, sourceData, err := util.YeepayCallback(response, yeepayConfig)
	if err != nil {
		return resp, err
	}
	err2 := json.Unmarshal([]byte(sourceData), &resp)
	if err2 != nil {
		return resp, errors.New("sourceData," + err2.Error() + "," + sourceData)
	}
	return resp, nil
}
