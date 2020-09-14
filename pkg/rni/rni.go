package rni

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

var client = resty.New()

const (
	h5SASSGenTokenURL = "https://aip.baidubce.com/rpc/2.0/brain/solution/faceprint/verifyToken/generate"
	h5SASSRedirectURL = "https://brain.baidu.com/face/print/?token=%s&successUrl=%s&failedUrl=%s"
	h5SASSResultURL   = "https://aip.baidubce.com/rpc/2.0/brain/solution/faceprint/result/detail"
)

// SDK
type SDK struct {
	token string
	error error
}

func (sdk SDK) H5SASSToken(planId string) (string, error) {
	if sdk.error != nil {
		return "", sdk.error
	}
	var res = new(H5SASSTokenResponse)
	var req = H5SASSTokenRequest{PlanId: planId}
	_, err := client.R().SetResult(res).SetBody(req).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("%s?access_token=%s", h5SASSGenTokenURL, sdk.token))
	if err == nil && res.Success == false {
		err = errors.New("get h5 sass verify token failed")
	}
	return res.Result.VerifyToken, err
}

func (sdk SDK) H5SASSResult(token string) (*H5SASSResultResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(H5SASSResultResponse)
	var req = H5SASSResultRequest{VerifyToken: token}
	_, err := client.R().SetResult(res).SetBody(req).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("%s?access_token=%s", h5SASSResultURL, sdk.token))
	return res, err
}

func NewSDK(token string, err error) SDK {
	return SDK{token: token, error: err}
}

func H5SASSRedirect(token, doneUrl, failUrl string) string {
	return fmt.Sprintf(h5SASSRedirectURL, token, url.QueryEscape(doneUrl), url.QueryEscape(failUrl))
}
