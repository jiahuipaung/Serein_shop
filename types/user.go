package types

type UserRegisterReq struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type UserTokenData struct {
	User         interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}

type UserLoginReq struct {
	Email    string `form:"user_name" json:"email"`
	Password string `form:"password" json:"password"`
}

type UserInfoUpdateReq struct {
	// NickName string `form:"nick_name" json:"nick_name"`
	
}

type UserInfoShowReq struct {
}

type SendEmailServiceReq struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	// OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}

type UserInfoResp struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}
