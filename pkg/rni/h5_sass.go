package rni

type H5SASSTokenRequest struct {
	PlanId string `json:"plan_id"`
}

type H5SASSTokenResponse struct {
	Success bool `json:"success"`
	Result  struct {
		VerifyToken string `json:"verify_token"`
	} `json:"result"`
}

type H5SASSResultRequest struct {
	VerifyToken string `json:"verify_token"`
}

type H5SASSResultResponse struct {
	Success bool `json:"success"`
	Result  struct {
		IdcardOcrResult struct {
			Name           string `json:"name"`
			Address        string `json:"address"`
			Birthday       string `json:"birthday"`
			IdCardNumber   string `json:"id_card_number"`
			Gender         string `json:"gender"`
			Nation         string `json:"nation"`
			ExpireTime     string `json:"expire_time"`
			IssueAuthority string `json:"issue_authority"`
			IssueTime      string `json:"issue_time"`
		} `json:"idcard_ocr_result"`
		IdcardImages struct {
			FrontBase64 string `json:"front_base64"`
			BackBase64  string `json:"back_base64"`
		} `json:"idcard_images"`
		VerifyResult struct {
			LivenessScore float64 `json:"liveness_score"`
			Score         float64 `json:"score"`
			Spoofing      float64 `json:"spoofing"`
		} `json:"verify_result"`
		IdcardConfirm struct {
			Name         string `json:"name"`
			IdcardNumber string `json:"idcard_number"`
		} `json:"idcard_confirm"`
	} `json:"result"`
}
