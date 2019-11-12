package e

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	UNKOWN_ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",

	CREATE_ERROR: "创建失败",
	DELETE_ERROR: "删除失败",
	UPDATR_ERROR: "更新失败",
	SELECT_ERROR: "查询失败",

	USERNAME_OR_PASSWORD_ERROR: "用户名或密码错误",
	USER_PASSWORD_NOT_CONSISTENT: "两次输入的密码不匹配",
	USERNAME_ALREADY_EXISTED: "用户名已存在",
	USER_SET_PASSWORD_ERROR: "设置密码失败",
	REGISTER_ERROR: "注册失败",

	OSS_CONFIG_ERROR: "OSS配置失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[UNKOWN_ERROR]
}
