package face

type UserAddRequest struct {
	Image           string `json:"image"`
	ImageType       string `json:"image_type"`
	GroupId         string `json:"group_id"`
	UserId          string `json:"user_id"`
	UserInfo        string `json:"user_info,omitempty"`
	QualityControl  string `json:"quality_control,omitempty"`
	LivenessControl string `json:"liveness_control,omitempty"`
	ActionType      string `json:"action_type,omitempty"`
}

type UserAddResponse struct {
	BaseResponse
	Result struct {
		FaceToken string `json:"face_token"`
		Location  struct {
			Left     float64 `json:"left"`
			Top      float64 `json:"top"`
			Width    float64 `json:"width"`
			Height   float64 `json:"height"`
			Rotation float64 `json:"rotation"`
		} `json:"location"`
	} `json:"result"`
}
