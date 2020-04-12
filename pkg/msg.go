package pkg

var Message = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	TokenExpired:     "令牌过期",
	TokenNotValidYet: "令牌尚未激活",
	TokenMalformed:   "非法令牌",
	TokenInvalid:     "无效令牌",
}

// GetMsg 获取错误信息
func GetMsg(code int) string {
	msg, ok := Message[code]
	if ok {
		return msg
	}
	return Message[ERROR]
}
