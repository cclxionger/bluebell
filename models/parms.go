package models

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
	// binding里面要求的字段名不是json字段名，是结构体里面的字段名
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ParamPostVote 对帖子投票
type ParamPostVote struct {
	PostID    int  `json:"post_id,string" binding:"required"`
	Direction int8 `json:"direction,string" binding:"required,oneof= 1 0 -1"` // 对应赞成，默认，反对
}
