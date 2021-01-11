package ugc

import (
	"encoding/json"
	"fmt"
)

const (
	imageCheckURl = "https://aip.baidubce.com/rest/2.0/solution/v1/img_censor/v2/user_defined"
)

type ImageRequest struct {
	Image   string `json:"image"`
	ImgUrl  string `json:"imgUrl"`
	ImgType int    `json:"imgType"`
}

func (req *ImageRequest) ToMap() map[string]string {
	var smp map[string]string
	bts, _ := json.Marshal(req)
	_ = json.Unmarshal(bts, &smp)
	return smp
}

type ImageResponse struct {
	BaseResponse
	Conclusion
	Data DataResponses `json:"data"`
}

func (sdk SDK) Image(req *ImageRequest) (*ImageResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(ImageResponse)
	_, err := client.R().SetResult(res).
		SetFormData(req.ToMap()).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post(fmt.Sprintf("%s?access_token=%s", imageCheckURl, sdk.token))
	return res, err
}
