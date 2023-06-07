package body

import (
	"github.com/letsfire/baidu/pkg"
)

type SeparateRequest struct {
	Image string `json:"image"`
	Type  string `json:"type"`
}

func (r *SeparateRequest) ToMap() map[string]string {
	if r.Type == "" {
		r.Type = "foreground"
	}
	return map[string]string{
		"type":  r.Type,
		"image": pkg.ImageBase64(r.Image),
	}
}

type SeparateResponse struct {
	pkg.BaseResponse
	Labelmap   string `json:"labelmap"`
	Scoremap   string `json:"scoremap"`
	Foreground string `json:"foreground"`
	PersonNum  int    `json:"person_num"`
	PersonInfo []struct {
		Top    float64 `json:"top"`
		Left   float64 `json:"left"`
		Width  float64 `json:"width"`
		Height float64 `json:"height"`
		Score  float64 `json:"score"`
	} `json:"person_info"`
}
