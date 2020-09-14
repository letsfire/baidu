package face

// MultiSearchRequest
type MultiSearchRequest struct {
	Image           string  `json:"image"`
	ImageType       string  `json:"image_type"`
	GroupIdList     string  `json:"group_id_list"`
	QualityControl  string  `json:"quality_control,omitempty"`
	LivenessControl string  `json:"liveness_control,omitempty"`
	MatchThreshold  float64 `json:"match_threshold,omitempty"`
	MaxFaceNum      int     `json:"max_face_num,omitempty"`
	MaxUserNum      int     `json:"max_user_num,omitempty"`
}

// MultiSearchResponse
type MultiSearchResponse struct {
	BaseResponse
	Result struct {
		FaceNum  int `json:"face_num"`
		FaceList []struct {
			FaceToken string `json:"face_token"`
			Location  struct {
				Left     float64 `json:"left"`
				Top      float64 `json:"top"`
				Width    float64 `json:"width"`
				Height   float64 `json:"height"`
				Rotation float64 `json:"rotation"`
			} `json:"location"`
			UserList []MatchUser `json:"user_list"`
		} `json:"face_list"`
	} `json:"result"`
}
