package utils

import (
	"errors"
	"fmt"
	"sdyxmall/business-examine/wrapper/page"
	"strings"
)

// 生成相同条件查询的统计数量SQL
func GetSameConditionCountSQL(sqlStr string) (string, error) {
	index := strings.Index(sqlStr, "FROM")
	if index == -1 {
		index = strings.Index(strings.ToUpper(sqlStr), "FROM")
	}
	if index == -1 {
		return "", errors.New("invalid sql")
	}
	return "SELECT count(*) AS total " + sqlStr[index:], nil
}

// 从page对象生成limit 条件SQL语句
func GetLimitSQLFromPage(pageMeta *page.Page) string {
	if pageMeta == nil || !pageMeta.IsPageRequest() {
		return ""
	}
	pageNum := pageMeta.PageNum
	pageSize := pageMeta.PageSize
	return fmt.Sprintf(" LIMIT %d, %d", (pageNum-1)*pageSize, pageSize)
}

//将切片类型转化为sql语句中的字符串 eg：["a","b","c"] ---> ?,?,?  --->"a","b","c"
func ConvertToSQLParams(sliceData interface{}, rawSqlParams []interface{}) (placeHolder string, params []interface{}) {
	placeHolder = ""
	if IsEmptyData(rawSqlParams) {
		rawSqlParams = make([]interface{}, 0)
	}
	Foreach(sliceData, func(e interface{}) {
		placeHolder += "?,"
		rawSqlParams = append(rawSqlParams, e)
	})
	if len(placeHolder) > 0 {
		placeHolder = placeHolder[:len(placeHolder)-1]
	}
	params = rawSqlParams
	return
}
