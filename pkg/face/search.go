package face

type SearchRequest struct {
	Image           string `json:"image"`
	ImageType       string `json:"image_type"`
	GroupIdList     string `json:"group_id_list"`
	QualityControl  string `json:"quality_control,omitempty"`
	LivenessControl string `json:"liveness_control,omitempty"`
	UserId          string `json:"user_id,omitempty"`
	MaxUserNum      int    `json:"max_user_num,omitempty"`
	MatchThreshold  int    `json:"match_threshold,omitempty"`
}

type SearchResponse struct {
	BaseResponse
	Result struct {
		FaceToken string      `json:"face_token"`
		UserList  []MatchUser `json:"user_list"`
	} `json:"result"`
}

type MatchUser struct {
	GroupId  string  `json:"group_id"`
	UserId   string  `json:"user_id"`
	UserInfo string  `json:"user_info"`
	Score    float64 `json:"score"`
}
