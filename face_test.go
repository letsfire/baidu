package baidu

import (
	"encoding/base64"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ai-sdk/baidu/pkg/face"
)

func TestSDK_Face(t *testing.T) {
	bts, _ := ioutil.ReadFile("./face_glasses.jpg")
	bs64 := base64.StdEncoding.EncodeToString(bts)
	res1, err1 := sdk.Face().Detect(&face.DetectRequest{
		Image:     bs64,
		ImageType: "BASE64",
		FaceField: "quality,mask,glasses",
	})
	assert.Nil(t, err1)
	assert.Equal(t, 1, res1.Result.FaceNum)
	var faceToken = res1.Result.FaceList[0].FaceToken
	_, err2 := sdk.Face().UserAdd(&face.UserAddRequest{
		Image:     faceToken,
		ImageType: "FACE_TOKEN",
		GroupId:   "TEST",
		UserId:    "TEST",
		UserInfo:  "TEST",
	})
	assert.Nil(t, err2)
	time.Sleep(time.Second * 5) // 延缓查询
	res3, err3 := sdk.Face().Search(&face.SearchRequest{
		Image:       bs64,
		ImageType:   "BASE64",
		GroupIdList: "TEST",
	})
	assert.Nil(t, err3)
	assert.EqualValues(t, 1, len(res3.Result.UserList))

	_, err4 := sdk.Face().UserDel(&face.UserDelRequest{
		UserId:    "TEST",
		GroupId:   "TEST",
	})
	assert.Nil(t, err4)
}

func TestSDK_Face_GROUP(t *testing.T) {
	_, err := sdk.Face().GroupDel(&face.GroupDelRequest{GroupId: "ab5ql0mytw4f"})
	assert.Nil(t, err)
}
