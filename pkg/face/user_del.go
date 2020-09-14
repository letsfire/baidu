package face

type UserDelRequest struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`
}

type UserDelResponse struct{ BaseResponse }
