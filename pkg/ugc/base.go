package ugc

import (
	"fmt"

	"github.com/letsfire/baidu/pkg"
)

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
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	Conclusion
	Msg         string  `json:"msg"`
	Type        int     `json:"type"`
	SubType     int     `json:"subType"`
	Probability float64 `json:"probability"`
	DatasetName string  `json:"datasetName"`
	Stars       []struct {
		Name        string  `json:"name"`
		Probability float64 `json:"probability"`
		DatasetName string  `json:"datasetName"`
	} `json:"stars"`
	Hits []struct {
		Probability float64  `json:"probability"`
		DatasetName string   `json:"datasetName"`
		Words       []string `json:"words"`
	} `json:"hits"`
}

func (dr *DataResponse) Words() []string {
	words := make([]string, 0)
	for _, hit := range dr.Hits {
		if len(hit.Words) == 0 {
			continue
		}
		words = append(words, hit.Words...)
	}
	return words
}

func (dr *DataResponse) Error() error {
	if dr.ErrorCode == 0 {
		return nil
	}
	return fmt.Errorf("[%d] - %s", dr.ErrorCode, dr.ErrorMsg)
}

type DataResponses []DataResponse

func (drs DataResponses) Error() error {
	for _, dr := range drs {
		if dr.Error() != nil {
			return dr.Error()
		}
	}
	return nil
}
