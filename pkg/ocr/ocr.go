package ocr

import (
	"github.com/go-resty/resty/v2"
)

var client = resty.New()

// SDK
type SDK struct {
	token string
	error error
}

func NewSDK(token string, err error) SDK {
	return SDK{token: token, error: err}
}
