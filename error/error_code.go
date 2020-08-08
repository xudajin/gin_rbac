package error

const (
	// 通用
	SUCCESS        = 200
	CREATE_SUCCESS = 201

	ERROR          = 500
	INVALID_PARAMS = 400

	//用户
	USER_NOT_FIND  = 1001
	PASSWORD_WRONG = 1002
	USER_ISEXIST   = 1003

	// 权限
	PERMISSION_ISEXIST = 1101
)
