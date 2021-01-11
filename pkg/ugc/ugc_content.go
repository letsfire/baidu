package ugc

import "fmt"

const (
	contentCheckURL = "https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined"
)

type ContentRequest struct {
	Text string `json:"text"`
}

func (r *ContentRequest) ToMap() map[string]string {
	return map[string]string{"text": r.Text}
}

type ContentResponse struct {
	BaseResponse
	Conclusion
	Data DataResponses `json:"data"`
}

func (sdk SDK) Content(req *ContentRequest) (*ContentResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(ContentResponse)
	_, err := client.R().SetResult(res).
		SetFormData(req.ToMap()).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post(fmt.Sprintf("%s?access_token=%s", contentCheckURL, sdk.token))
	return res, err
}
