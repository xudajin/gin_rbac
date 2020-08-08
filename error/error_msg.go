package error

var MsgFlags = map[int64]string{
	// 通用
	SUCCESS:        "成功",
	ERROR:          "失败",
	INVALID_PARAMS: "参数错误",
	CREATE_SUCCESS: "创建成功",

	//用户
	USER_NOT_FIND:  "用户不存在",
	PASSWORD_WRONG: "密码错误",
	USER_ISEXIST:   "用户已存在",

	// 权限
	PERMISSION_ISEXIST: "权限已存在",
}

func Msg(code int64) string {
	msg := MsgFlags[code]
	return msg
}
