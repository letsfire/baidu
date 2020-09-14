package face

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

var client = resty.New()

const (
	detectURL      = "https://aip.baidubce.com/rest/2.0/face/v3/detect"
	userAddURL     = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add"
	userDelURL     = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/delete"
	searchURL      = "https://aip.baidubce.com/rest/2.0/face/v3/search"
	multiSearchURL = "https://aip.baidubce.com/rest/2.0/face/v3/multi-search"
	groupDelURL    = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/delete"
)

type SDK struct {
	token string
	error error
}

// Detect 人脸检测
// https://ai.baidu.com/ai-doc/FACE/yk37c1u4t
func (sdk SDK) Detect(req *DetectRequest) (*DetectResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(DetectResponse)
	_, err := client.R().SetBody(req).SetResult(res).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("%s?access_token=%s", detectURL, sdk.token))
	return res, err
}

// UserAdd 注册人脸
// https://ai.baidu.com/ai-doc/FACE/7k37c1twu#%E4%BA%BA%E8%84%B8%E6%B3%A8%E5%86%8C
func (sdk SDK) UserAdd(req *UserAddRequest) (*UserAddResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(UserAddResponse)
	_, err := client.R().SetBody(req).SetResult(res).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("%s?access_token=%s", userAddURL, sdk.token))
	return res, err
}

// UserDel 删除人脸
// https://ai.baidu.com/ai-doc/FACE/7k37c1twu
func (sdk SDK) UserDel(req *UserDelRequest) (*UserDelResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(UserDelResponse)
	_, err := client.R().SetBody(req).SetResult(res).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("%s?access_token=%s", userDelURL, sdk.token))
	return res, err
}

// Search 人脸搜索
// https://ai.baidu.com/ai-doc/FACE/Gk37c1uzc
func (sdk SDK) Search(req *SearchRequest) (*SearchResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(SearchResponse)
	_, err := client.R().SetBody(req).SetResult(res).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("%s?access_token=%s", searchURL, sdk.token))
	return res, err
}

// GroupDel 删除人脸组
// https://ai.baidu.com/ai-doc/FACE/7k37c1twu#%E5%88%A0%E9%99%A4%E7%94%A8%E6%88%B7%E7%BB%84
func (sdk SDK) GroupDel(req *GroupDelRequest) (*GroupDelResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(GroupDelResponse)
	_, err := client.R().SetBody(req).SetResult(res).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("%s?access_token=%s", groupDelURL, sdk.token))
	return res, err
}

// MultiSearch M:N搜索
// https://ai.baidu.com/ai-doc/FACE/Gk37c1uzc#%E4%BA%BA%E8%84%B8%E6%90%9C%E7%B4%A2-mn-%E8%AF%86%E5%88%AB
func (sdk SDK) MultiSearch(req *MultiSearchRequest) (*MultiSearchResponse, error) {
	if sdk.error != nil {
		return nil, sdk.error
	}
	var res = new(MultiSearchResponse)
	_, err := client.R().SetBody(req).SetResult(res).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("%s?access_token=%s", multiSearchURL, sdk.token))
	return res, err
}

func NewSDK(token string, err error) SDK {
	return SDK{token: token, error: err}
}
