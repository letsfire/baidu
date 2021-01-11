package baidu

import (
	"encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/letsfire/baidu/pkg/ugc"
)

func TestSDK_Ugc_Image(t *testing.T) {
	bts, _ := ioutil.ReadFile("./ugc_error2.jpg")
	bs64 := base64.StdEncoding.EncodeToString(bts)
	res, err := sdk.Ugc().Image(&ugc.ImageRequest{Image: bs64})
	assert.Nil(t, err)
	assert.False(t, res.IsFailed())
	assert.False(t, res.IsPassed())
}

func TestSDK_Ugc_Content(t *testing.T) {
	res, err := sdk.Ugc().Content(&ugc.ContentRequest{
		Text: "中国共产党",
	})
	assert.Nil(t, err)
	assert.False(t, res.IsFailed())
	assert.False(t, res.IsPassed())
}

func TestSDK_Ugc_Video(t *testing.T) {
	res, err := sdk.Ugc().Video(&ugc.VideoRequest{
		Name:     "色情片",
		VideoURL: "https://ing-sports-dev.oss-cn-hangzhou.aliyuncs.com/ugc_error.mp4",
		ExtId:    "test",
	})
	assert.Nil(t, err)
	assert.False(t, res.IsFailed())
	assert.False(t, res.IsPassed())
}
