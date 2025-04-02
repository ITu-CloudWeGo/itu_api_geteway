package model

type PostResp struct {
	Pid          int64    `json:"pid"`
	AuthorId     int64    `json:"author_id"`
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	ImageUrls    []string `json:"image_urls"`
	Tags         []string `json:"tags"`
	CreateTime   string   `json:"create_time"`
	UpdateTime   string   `json:"update_time"`
	ViewCount    int64    `json:"view_count"`
	LikeCount    int64    `json:"like_count"`
	CommentCount int64    `json:"comment_count"`
	IsLiked      bool     `json:"is_liked"`
	IsFavorite   bool     `json:"is_favorite"`
}

type UpdatePostReq struct {
	Uid     int64  `json:"uid"`
	Pid     int64  `json:"pid"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DeletePostReq struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}
