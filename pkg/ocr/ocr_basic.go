package ocr

import (
	"fmt"
)

const (
	basicGeneralUrl  = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
	basicAccurateUrl = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"
)

// BasicGeneral 通用文字识别
// https://ai.baidu.com/ai-doc/OCR/zk3h7xz52
func (sdk SDK) BasicGeneral(req *BasicGeneralRequest) (*BasicGeneralResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(BasicGeneralResponse)
	_, err := client.R().SetResult(res).
		SetFormData(req.ToMap()).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post(fmt.Sprintf("%s?access_token=%s", basicGeneralUrl, sdk.token))
	return res, err
}

// BasicAccurate 通用文字识别(高精度)
// https://ai.baidu.com/ai-doc/OCR/1k3h7y3db
func (sdk SDK) BasicAccurate(req *BasicAccurateRequest) (*BasicAccurateResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(BasicAccurateResponse)
	_, err := client.R().SetResult(res).
		SetFormData(req.ToMap()).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post(fmt.Sprintf("%s?access_token=%s", basicAccurateUrl, sdk.token))
	return res, err
}
