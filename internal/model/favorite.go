package model

type AddDelFavoriteRequest struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}
