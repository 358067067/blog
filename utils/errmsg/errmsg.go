package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	// code = 1000 用户模块错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST    = 1004
	ERROR_TOKEN_RUNTIME  = 1005
	ERROR_TOKEN_WRONG    = 1006
	ERROR_TOKEN_TYPE     = 1007
	// code = 2000 文章
	ERROR_ARTICLE_NIL = 1011
	// code = 3000 分类
	ERROR_CNAME_USED = 1021
)

var codeMsg = map[int]string{
	SUCCESS:              "OK",
	ERROR:                "FAIL",
	ERROR_USERNAME_USED:  "用户名已存在！",
	ERROR_PASSWORD_WRONG: "密码错误！",
	ERROR_USER_NOT_EXIST: "用户不存在！",
	ERROR_TOKEN_EXIST:    "token不存在！",
	ERROR_TOKEN_RUNTIME:  "token过期！",
	ERROR_TOKEN_WRONG:    "token不正确！",
	ERROR_TOKEN_TYPE:     "token格式错误！",
	ERROR_ARTICLE_NIL:    "文章部分数值不能为空！",
	ERROR_CNAME_USED:     "分类名称已存在！",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
