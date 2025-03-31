package model

type CreateTagsRequest struct {
	Pid  int64  `json:"pid"`
	Tags string `json:"tags"`
}

type CreateTagsResponse struct {
	Status int64  `json:"status"`
	Msg    string `json:"message"`
}

type DelTagsRequest struct {
	Pid  int64  `json:"pid"`
	Tags string `json:"tags"`
}

type DelTagsResponse struct {
	Status int64  `json:"status"`
	Msg    string `json:"message"`
}
