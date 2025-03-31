package model

type Data struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
}

type SignRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
}

type SignResponse struct {
	Status int64  `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"user"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status int64  `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"user"`
}

type GetUserRequest struct {
	Uid      int64  `json:"uid"`
	UserName string `json:"username"`
}

type GetUserResponse struct {
	Status int64  `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"user"`
}

type UpdateUserRequest struct {
	Uid           int64  `json:"uid"`
	NewEmail      string `json:"new_email"`
	NewNickname   string `json:"new_nickname"`
	NewIcon       string `json:"new_icon"`
	Password      string `json:"password"`
	NewPassword   string `json:"new_password"`
	ConfirmPasswd string `json:"confirm_passwd"`
}

type UpdateUserResponse struct {
	Status int64  `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"user"`
}
