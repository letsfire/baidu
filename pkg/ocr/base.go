package ocr

import (
	"github.com/ai-sdk/baidu/pkg"
)

// BaseResponse
type BaseResponse struct{ pkg.BaseResponse }

// UrlError 图片下载失败
func (br *BaseResponse) UrlError() bool {
	return br.ErrorCode == 282112
}
