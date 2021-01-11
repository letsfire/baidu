package ugc

import (
	"github.com/letsfire/baidu/pkg"
)

// BaseResponse
type BaseResponse struct{ pkg.BaseResponse }

type Conclusion struct {
	Conclusion     string `json:"conclusion"`
	ConclusionType int    `json:"conclusionType"`
}

func (c Conclusion) IsPassed() bool {
	return c.ConclusionType == 1
}

func (c Conclusion) IsFailed() bool {
	return c.ConclusionType == 4
}

type DataResponse struct {
	ErrorCode   int     `json:"error_code"`
	ErrorMsg    string  `json:"error_msg"`
	Type        int     `json:"type"`
	SubType     int     `json:"subType"`
	Msg         string  `json:"msg"`
	Probability float64 `json:"probability"`
	DatasetName float64 `json:"datasetName"`
	Stars       []struct {
		Name        string  `json:"name"`
		Probability float64 `json:"probability"`
		DatasetName float64 `json:"datasetName"`
	} `json:"stars"`
	Hits []struct {
		Probability float64  `json:"probability"`
		DatasetName string   `json:"datasetName"`
		Words       []string `json:"words"`
	} `json:"hits"`
}

type DataResponses []DataResponse
