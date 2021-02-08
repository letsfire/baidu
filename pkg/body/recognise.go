package body

import (
	"github.com/letsfire/baidu/pkg"
)

// RecogniseRequest
type RecogniseRequest struct {
	Image    string `json:"image"`
	BaikeNum int    `json:"baike_num"`
}

func (r *RecogniseRequest) ToMap() map[string]string {
	return map[string]string{
		"image": pkg.ImageBase64(r.Image),
	}
}

// RecogniseResponse
type RecogniseResponse struct {
	pkg.BaseResponse
	ResultNum int `json:"result_num"`
	Result    []struct {
		Keyword   string  `json:"keyword"`
		Score     float64 `json:"score"`
		Root      string  `json:"root"`
		BaikeInfo struct {
			BaikeUrl    string `json:"baike_url"`
			ImageUrl    string `json:"image_url"`
			Description string `json:"description"`
		} `json:"baike_info"`
	} `json:"result"`
}
