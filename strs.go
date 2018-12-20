package strx

import "strings"

// bool to string
// return "true" or "false"
func BoolToString(source bool) string {
	if source {
		return "true"
	} else {
		return "false"
	}
}

// 判断是否是空，是空返回true，不是空返回false
func IsEmpty(source string) bool {
	if strings.Trim(source, " ") == "" {
		return true
	}
	return false
}

// 判断是否不是空，是空返回false，不是空返回true
func IsNotEmpty(source string) bool {
	return !IsEmpty(source)
}

// 字符串替换，将所有的oldStr替换为newStr
func ReplaceM(sourceStr string, newStr string, oldStr ...string) (resultStr string) {
	for i := 0; i < len(oldStr); i++ {
		sourceStr = strings.Replace(sourceStr, oldStr[i], newStr, -1)
	}
	return sourceStr
}

// 返回第一次出现 start与第一次出现end之间的字符串
func Between(sourceStr, startStr, endStr string) string {
	startIndex := strings.Index(sourceStr, startStr)
	endIndex := strings.Index(sourceStr, endStr)
	return sourceStr[startIndex+len(startStr) : endIndex]
}
