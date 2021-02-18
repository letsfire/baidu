package body

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

var client = resty.New()

const (
	cartoonURL   = "https://aip.baidubce.com/rest/2.0/image-process/v1/selfie_anime"
	separateURL  = "https://aip.baidubce.com/rest/2.0/image-classify/v1/body_seg"
	recogniseURL = "https://aip.baidubce.com/api/v1/solution/direct/imagerecognition/combination"
)

type SDK struct {
	token string
	error error
}

// Cartoon 人像卡通
// https://ai.baidu.com/ai-doc/IMAGEPROCESS/Mk4i6olx5
func (sdk SDK) Cartoon(req *CartoonRequest) (*CartoonResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(CartoonResponse)
	_, err := client.R().SetResult(res).
		SetFormData(req.ToMap()).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post(fmt.Sprintf("%s?access_token=%s", cartoonURL, sdk.token))
	return res, err
}

// Separate 人像分割
// https://ai.baidu.com/ai-doc/BODY/Fk3cpyxua
func (sdk SDK) Separate(req *SeparateRequest) (*SeparateResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(SeparateResponse)
	_, err := client.R().SetResult(res).
		SetFormData(req.ToMap()).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post(fmt.Sprintf("%s?access_token=%s", separateURL, sdk.token))
	return res, err
}

// Recognise 图像识别
// https://ai.baidu.com/ai-doc/IMAGERECOGNITION/Kkbg3gxs7
func (sdk SDK) Recognise(req *RecogniseRequest) (*RecogniseResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(RecogniseResponse)
	_, err := client.R().SetBody(req).SetResult(res).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("%s?access_token=%s", recogniseURL, sdk.token))
	return res, err
}

func NewSDK(token string, err error) SDK {
	return SDK{token: token, error: err}
}
