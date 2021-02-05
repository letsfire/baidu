package pkg

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type BaseResponse struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	LogId     int    `json:"log_id"`
	Timestamp int    `json:"timestamp"`
}

func (br *BaseResponse) Error() error {
	if br.ErrorCode == 0 {
		return nil
	}
	return fmt.Errorf(
		"[%d] - %s, LogId = %d, Timestamp: %d",
		br.ErrorCode, br.ErrorMsg, br.LogId, br.Timestamp,
	)
}

// QPSError 是否QPS超限
func (br *BaseResponse) QPSError() bool {
	return br.ErrorCode == 18
}

// BoolString 转字符串
type BoolString bool

func (bs BoolString) MarshalJSON() ([]byte, error) {
	if bs {
		return json.Marshal("true")
	}
	return json.Marshal("false")
}

func ImageBase64(image string) string {
	isUrl := strings.HasPrefix(image, "http://") ||
		strings.HasPrefix(image, "https://")
	if isUrl == false {
		return image
	}
	if res, err := http.Get(image); err != nil {
		panic(err)
	} else {
		defer res.Body.Close()
		bts, _ := ioutil.ReadAll(res.Body)
		return base64.StdEncoding.EncodeToString(bts)
	}
}
