package e

var ErrMsg = map[int]string{
	SUCCESS:         "ok",
	ERROR:           "内部错误",
	UNKNOW_IDENTITY: "未知身份",
}

func GetMsg(code int) string {
	return ErrMsg[code]
}
