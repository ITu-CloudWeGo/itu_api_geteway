package model

//type GenerateTokenRequest struct {
//	Uid uint64 `json:"uid"`
//}
//type GenerateTokenResponse struct {
//	AccessToken        string `json:"access_token"`
//	RefreshToken       string `json:"refresh_token"`
//	AccessTokenExpire  int64  `json:"access_token_expire"`
//	RefreshTokenExpire int64  `json:"refresh_token_expire"`
//}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken        string `json:"access_token"`
	RefreshToken       string `json:"refresh_token"`
	AccessTokenExpire  int64  `json:"access_token_expire"`
	RefreshTokenExpire int64  `json:"refresh_token_expire"`
}

type EmailRequest struct {
	Email string `json:"email"`
}

type EmailResponse struct {
	Captcha int64 `json:"email"`
}
