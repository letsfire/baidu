package body

import (
	"github.com/letsfire/baidu/pkg"
)

// RecogniseRequest
type RecogniseRequest struct {
	Image     string                            `json:"image,omitempty"`
	ImgUrl    string                            `json:"imgUrl,omitempty"`
	Scenes    []string                          `json:"scenes"`
	SceneConf map[string]map[string]interface{} `json:"sceneConf"`
}

func (r *RecogniseRequest) setImage(image string) *RecogniseRequest {
	if !pkg.IsURL(image) {
		r.Image = image
	} else {
		r.ImgUrl = image
	}
	return r
}

// NewPlantRecognise 植物识别
func NewPlantRecognise(image string, baikeNum int) *RecogniseRequest {
	req := RecogniseRequest{
		Scenes: []string{"plant"},
		SceneConf: map[string]map[string]interface{}{
			"plant": {"baike_num": baikeNum},
		},
	}
	return req.setImage(image)
}

// NewAnimalRecognise 动物识别
func NewAnimalRecognise(image string, baikeNum int) *RecogniseRequest {
	req := RecogniseRequest{
		Scenes: []string{"animal"},
		SceneConf: map[string]map[string]interface{}{
			"animal": {"top_num": baikeNum, "baike_num": baikeNum},
		},
	}
	return req.setImage(image)
}

// NewAdvancedGeneralRecognise 通用识别
func NewAdvancedGeneralRecognise(image string, baikeNum int) *RecogniseRequest {
	req := RecogniseRequest{
		Scenes: []string{"advanced_general"},
		SceneConf: map[string]map[string]interface{}{
			"advanced_general": {"baike_num": baikeNum},
		},
	}
	return req.setImage(image)
}

// RecogniseResponse
type RecogniseResponse struct {
	pkg.BaseResponse
	Result struct {
		Plant           PlantResponse           `json:"plant"`
		Animal          AnimalResponse          `json:"animal"`
		AdvancedGeneral AdvancedGeneralResponse `json:"advanced_general"`
	} `json:"result"`
}

// BaikeInfo
type BaikeInfo struct {
	BaikeUrl    string `json:"baike_url"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

// PlantResponse
type PlantResponse struct {
	pkg.BaseResponse
	Result []struct {
		Name      string    `json:"name"`
		Score     float64   `json:"score"`
		BaikeInfo BaikeInfo `json:"baike_info"`
	} `json:"result"`
}

// AnimalResponse
type AnimalResponse struct {
	pkg.BaseResponse
	Result []struct {
		Name      string    `json:"name"`
		Score     string    `json:"score"`
		BaikeInfo BaikeInfo `json:"baike_info"`
	} `json:"result"`
}

// AdvancedGeneralResponse
type AdvancedGeneralResponse struct {
	pkg.BaseResponse
	ResultNum int `json:"result_num"`
	Result    []struct {
		Keyword   string    `json:"keyword"`
		Score     float64   `json:"score"`
		Root      string    `json:"root"`
		BaikeInfo BaikeInfo `json:"baike_info"`
	} `json:"result"`
}
