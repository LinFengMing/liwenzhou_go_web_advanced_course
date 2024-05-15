package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "請求參數錯誤",
	CodeUserExist:       "用戶已存在",
	CodeUserNotExist:    "用戶不存在",
	CodeInvalidPassword: "用戶名或密碼錯誤",
	CodeServerBusy:      "服務忙碌中",
	CodeNeedLogin:       "需要登陸",
	CodeInvalidToken:    "無效的 Token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
