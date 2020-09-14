package baidu

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ai-sdk/baidu/pkg/ocr"
)

func TestSDK_Ocr(t *testing.T) {
	bts, _ := ioutil.ReadFile("./ocr.jpg")
	bs64 := base64.StdEncoding.EncodeToString(bts)
	res, err := sdk.Ocr().BasicAccurate(&ocr.BasicAccurateRequest{
		Image: bs64, LanguageType: "ENG", Probability: true,
	})
	assert.Nil(t, err)
	fmt.Println(res)
}
