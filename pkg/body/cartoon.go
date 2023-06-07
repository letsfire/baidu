package body

import (
	"github.com/letsfire/baidu/pkg"
)

type CartoonRequest struct {
	Image string `json:"image"`
}

func (r *CartoonRequest) ToMap() map[string]string {
	return map[string]string{
		"type":  "anime",
		"image": pkg.ImageBase64(r.Image),
	}
}

type CartoonResponse struct {
	pkg.BaseResponse
	Image string `json:"image"`
}
