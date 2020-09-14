package face

import (
	"github.com/ai-sdk/baidu/pkg"
)

// BaseResponse
type BaseResponse struct{ pkg.BaseResponse }

// NoFace 图片没有人脸
func (br *BaseResponse) NoFace() bool {
	return br.ErrorCode == 222202
}

// NoUser 用户不存在
func (br *BaseResponse) NoUser() bool {
	return br.ErrorCode == 223103
}

// NoMatch 未找到匹配的用户
func (br *BaseResponse) NoMatch() bool {
	return br.ErrorCode == 222207
}

// UrlError 图片下载失败
func (br *BaseResponse) UrlError() bool {
	return br.ErrorCode == 222204
}
