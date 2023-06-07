package face

type GroupDelRequest struct {
	GroupId string `json:"group_id"`
}

type GroupDelResponse struct{ BaseResponse }
