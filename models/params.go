package models

// 定義請求的參數結構體

type ParamSigUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID    int64 `json:"post_id,string" binding:"required"`                // 帖子 id
	Direction int8  `json:"direction,string" binding:"required,oneof=-1 0 1"` // 贊成1;反對-1;取消0
}
