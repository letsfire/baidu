package baidu

import (
	"encoding/base64"
	"fmt"
	"github.com/letsfire/baidu/pkg/ugc"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestSDK_Ugc_Image(t *testing.T) {
	bts, _ := ioutil.ReadFile("./ugc_error.jpg")
	bs64 := base64.StdEncoding.EncodeToString(bts)
	res, err := sdk.Ugc().Image(&ugc.ImageRequest{Image: bs64})
	fmt.Println(res)
	assert.Nil(t, err)
	assert.True(t, res.IsFailed())
	assert.False(t, res.IsPassed())
}

func TestSDK_Ugc_Content(t *testing.T) {
	res, err := sdk.Ugc().Content(&ugc.ContentRequest{
		Text: "中国共产党",
	})
	fmt.Println(res)
	assert.Nil(t, err)
	assert.True(t, res.IsFailed())
	assert.False(t, res.IsPassed())
}

func TestSDK_Ugc_Video(t *testing.T) {
	res, err := sdk.Ugc().Video(&ugc.VideoRequest{
		Name:     "色情片",
		VideoURL: "",
	})
	fmt.Println(res)
	assert.Nil(t, err)
}
