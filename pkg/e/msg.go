package e


var MsgFlags = map[int]string{
	SUCCESS:                         "success",
	ERROR:                           "server error",
	PARAMS_ERROR:                    "请求参数错误",
	NOT_FOUND:						 "请求地址错误",
}


func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
