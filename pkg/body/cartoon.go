package body

import (
	"github.com/letsfire/baidu/pkg"
)

// CartoonRequest
type CartoonRequest struct {
	Image string `json:"image"`
}

func (r *CartoonRequest) ToMap() map[string]string {
	return map[string]string{
		"type":  "anime",
		"image": pkg.ImageBase64(r.Image),
	}
}

// CartoonResponse
type CartoonResponse struct {
	pkg.BaseResponse
	Image string `json:"image"`
}
