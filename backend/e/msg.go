package e

var MsgFlags = map[int]string{
	SUCCESS:      "成功",
	SELECT_ERROR: "无效选择",
	CREATE_ERROR: "创建失败",
}

func GetMsg(code int) string {
	return MsgFlags[code]
}
