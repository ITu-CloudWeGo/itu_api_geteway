package model

type AddLikesRequest struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

type AddLikesResponse struct {
	Status int64  `json:"status"`
	Msg    string `json:"message"`
}

type DelLikesRequest struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

type DelLikesResponse struct {
	Status int64  `json:"status"`
	Msg    string `json:"message"`
}
