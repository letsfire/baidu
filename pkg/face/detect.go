package face

type DetectRequest struct {
	Image           string `json:"image"`
	ImageType       string `json:"image_type"`
	FaceField       string `json:"face_field,omitempty"`
	MaxFaceNum      uint32 `json:"max_face_num,omitempty"`
	FaceType        string `json:"face_type,omitempty"`
	LivenessControl string `json:"liveness_control,omitempty"`
}

type DetectResponse struct {
	BaseResponse
	Result struct {
		FaceNum  int           `json:"face_num"`
		FaceList []*DetectFace `json:"face_list"`
	} `json:"result"`
}

type DetectFace struct {
	FaceToken string `json:"face_token"`
	Location  struct {
		Left     float64 `json:"left"`
		Top      float64 `json:"top"`
		Width    float64 `json:"width"`
		Height   float64 `json:"height"`
		Rotation float64 `json:"rotation"`
	} `json:"location"`
	FaceProbability float64 `json:"face_probability"`
	Angle           struct {
		Yaw   float64 `json:"yaw"`
		Pitch float64 `json:"pitch"`
		Roll  float64 `json:"roll"`
	} `json:"angle"`
	Age        float64 `json:"age"`
	Beauty     float64 `json:"beauty"`
	Expression struct {
		Type        string  `json:"type"`
		Probability float64 `json:"probability"`
	} `json:"expression"`
	Gender struct {
		Type        string  `json:"type"`
		Probability float64 `json:"probability"`
	} `json:"gender"`
	Glasses struct {
		Type        string  `json:"type"`
		Probability float64 `json:"probability"`
	} `json:"glasses"`
	Emotion struct {
		Type        string  `json:"type"`
		Probability float64 `json:"probability"`
	} `json:"emotion"`
	Mask struct {
		Type        int     `json:"type"`
		Probability float64 `json:"probability"`
	} `json:"mask"`
	Quality struct {
		Occlusion struct {
			LeftEye     float64 `json:"left_eye"`
			RightEye    float64 `json:"right_eye"`
			Nose        float64 `json:"nose"`
			Mouth       float64 `json:"mouth"`
			LeftCheek   float64 `json:"left_cheek"`
			RightCheek  float64 `json:"right_cheek"`
			ChinContour float64 `json:"chin_contour"`
		} `json:"occlusion"`
		Blur         float64 `json:"blur"`
		Illumination float64 `json:"illumination"`
		Completeness int8    `json:"completeness"`
	} `json:"quality"`
}

func (df *DetectFace) HasMask() bool {
	return df.Mask.Type == 1 &&
		df.Mask.Probability >= 0.75
}

func (df *DetectFace) HasGlasses() bool {
	return df.Glasses.Type != "none" &&
		df.Glasses.Probability >= 0.75
}
