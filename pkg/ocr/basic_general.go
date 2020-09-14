package ocr

import (
	"encoding/json"

	"github.com/letsfire/baidu/pkg"
)

// BasicGeneralRequest
type BasicGeneralRequest struct {
	Image           string         `json:"image,omitempty"`
	Url             string         `json:"url,omitempty"`
	LanguageType    string         `json:"language_type,omitempty"`
	DetectDirection pkg.BoolString `json:"detect_direction,omitempty"`
	DetectLanguage  pkg.BoolString `json:"detect_language,omitempty"`
	Paragraph       pkg.BoolString `json:"paragraph,omitempty"`
	Probability     pkg.BoolString `json:"probability,omitempty"`
}

func (req *BasicGeneralRequest) ToMap() map[string]string {
	var smp map[string]string
	bts, _ := json.Marshal(req)
	_ = json.Unmarshal(bts, &smp)
	return smp
}

// BasicGeneralResponse
type BasicGeneralResponse struct {
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
