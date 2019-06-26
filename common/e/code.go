package e

/**
 * 自定内内部错误好，统一异常信息处理
 */
const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	ERR_NO_AUTH    = 401

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004


)