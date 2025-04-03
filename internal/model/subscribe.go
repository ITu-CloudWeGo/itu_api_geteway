package model

type AddDelSubscribeRequest struct {
	Uid      int64 `json:"uid"`
	AuthorId int64 `json:"author_id"`
}
