package ocr

import (
	"encoding/json"

	"github.com/ai-sdk/baidu/pkg"
)

// BasicAccurateRequest
type BasicAccurateRequest struct {
	Image           string         `json:"image,omitempty"`
	LanguageType    string         `json:"language_type,omitempty"`
	DetectDirection pkg.BoolString `json:"detect_direction,omitempty"`
	Paragraph       pkg.BoolString `json:"paragraph,omitempty"`
	Probability     pkg.BoolString `json:"probability,omitempty"`
}

func (req *BasicAccurateRequest) ToMap() map[string]string {
	var smp map[string]string
	bts, _ := json.Marshal(req)
	_ = json.Unmarshal(bts, &smp)
	return smp
}

// BasicAccurateResponse
type BasicAccurateResponse struct {
	BaseResponse
	Direction   int `json:"direction"`
	WordsResult []struct {
		Words       string `json:"words"`
		Probability struct {
			Average  float64 `json:"average"`
			Min      float64 `json:"min"`
			Variance float64 `json:"variance"`
		} `json:"probability"`
	} `json:"words_result"`
	WordsResultNum   int `json:"words_result_num"`
	ParagraphsResult []struct {
		WordsResultIdx []int `json:"words_result_idx"`
	} `json:"paragraphs_result"`
	ParagraphsResultNum int `json:"paragraphs_result_num"`
	Language            int `json:"language"`
}
