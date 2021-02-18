package baidu

import (
	"encoding/base64"
	"fmt"
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

func TestSDK_Body_Recognise(t *testing.T) {
	req := body.NewAdvancedGeneralRecognise("http://imgsrc.baidu.com/baike/pic/item/91ef76c6a7efce1b27893518a451f3deb58f6546.jpg", 3)
	res, err := sdk.Body().Recognise(req)
	fmt.Println(fmt.Sprintf("%+v", res), err)
}
