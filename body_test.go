package baidu

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"

	"github.com/letsfire/baidu/pkg/body"
)

func TestSDK_Body_Cartoon(t *testing.T) {
	bts, _ := ioutil.ReadFile("./ugc_error.jpg")
	bs64 := base64.StdEncoding.EncodeToString(bts)
	req := body.CartoonRequest{Image: bs64}
	res, _ := sdk.Body().Cartoon(&req)
	img, _ := base64.StdEncoding.DecodeString(res.Image)
	_ = ioutil.WriteFile("./ugc_error_car.png", img, os.ModePerm)
}

func TestSDK_Body_Separate(t *testing.T) {
	bts, _ := ioutil.ReadFile("./ugc_error.jpg")
	bs64 := base64.StdEncoding.EncodeToString(bts)
	req := body.SeparateRequest{Image: bs64}
	res, _ := sdk.Body().Separate(&req)
	img, _ := base64.StdEncoding.DecodeString(res.Foreground)
	_ = ioutil.WriteFile("./ugc_error_sep.png", img, os.ModePerm)
}
