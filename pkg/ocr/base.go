package ocr

import (
	"github.com/letsfire/baidu/pkg"
)

type BaseResponse struct{ pkg.BaseResponse }

// UrlError 图片下载失败
func (br *BaseResponse) UrlError() bool {
	return br.ErrorCode == 282112
}
