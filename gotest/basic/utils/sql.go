package utils

//将动态参数解析出 占位符 以及 可用的编程
func SqlParamsConverter(params []interface{}) (placeholders interface{}) {
	if params == nil || len(params) == 0 {
		return nil
	}
	sqlParams := ""
	for range params {
		sqlParams += "?,"
	}
	sqlParams = sqlParams[:len(sqlParams)-1]
	return sqlParams
}
