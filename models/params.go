package models

// 定義請求的參數結構體

type ParamSigUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
