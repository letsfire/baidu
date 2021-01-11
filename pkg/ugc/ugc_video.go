package ugc

import (
	"encoding/json"
	"fmt"
)

const (
	videoCheckURl = "https://aip.baidubce.com/rest/2.0/solution/v1/video_censor/v2/user_defined"
)

type VideoRequest struct {
	Name     string `json:"name"`
	VideoURL string `json:"videoUrl"`
	ExtId    string `json:"extId"`
	ExtInfo  []struct {
		Subject string `json:"subject"`
		Fields  []struct {
			Title string `json:"title"`
			Value string `json:"value"`
		} `json:"fields"`
	} `json:"extInfo"`
}

func (req *VideoRequest) ToMap() map[string]string {
	var smp map[string]string
	bts, _ := json.Marshal(req)
	_ = json.Unmarshal(bts, &smp)
	return smp
}

type VideoResponse struct {
	BaseResponse
	Conclusion
	Msg      string `json:"msg"`
	IsHitMd5 bool   `json:"isHitMd5"`
	Frames   []struct {
		Conclusion
		FrameTimeStamp    int64         `json:"frameTimeStamp"`
		FrameUrl          string        `json:"frameUrl"`
		FrameThumbnailUrl string        `json:"frameThumbnailUrl"`
		Data              DataResponses `json:"data"`
	} `json:"frames"`
}

func (sdk SDK) Video(req *VideoRequest) (*VideoResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(VideoResponse)
	_, err := client.R().SetResult(res).
		SetFormData(req.ToMap()).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post(fmt.Sprintf("%s?access_token=%s", videoCheckURl, sdk.token))
	return res, err
}
