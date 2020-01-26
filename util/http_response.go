package util

import (
	"fmt"
	"time"
)

type HttpResponse struct {
	code      int
	bytes     []byte
	startTime time.Time
}

func (resp *HttpResponse) SetCode(code int) *HttpResponse {
	resp.code = code
	return resp
}

func (resp HttpResponse) GetCode() int {
	return resp.code
}

func (resp HttpResponse) IsOk() bool {
	return 200 == resp.code
}

func (resp *HttpResponse) SetBytes(bytes []byte) *HttpResponse {
	resp.bytes = bytes
	return resp
}

func (resp HttpResponse) GetBytes() []byte {
	return resp.bytes
}

func (resp *HttpResponse) SetStartTime(startTime time.Time) *HttpResponse {
	resp.startTime = startTime
	return resp
}

func (resp HttpResponse) GetLatency() float64 {
	latency := time.Since(resp.startTime)
	return latency.Seconds()
}

func (resp HttpResponse) GetLatencyStr() string {
	return fmt.Sprintf("%.2fs", resp.GetLatency())
}
