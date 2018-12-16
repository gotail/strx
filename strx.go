package strx

import (
	"strings"
)

type New string

// bool to string
// return "true" or "false"
func BoolToString(source bool) string {
	if source {
		return "true"
	} else {
		return "false"
	}
}

// 字符串清除
func (s *New) Clean(cleanStr ...string) *New {
	return s.ReplaceM("", cleanStr...)
}

// 字符串替换右侧
func (s *New) TrimRight(trimStr string) *New {
	*s = New(strings.TrimRight(s.Val(), trimStr))
	return s
}

// 字符串替换左侧
func (s *New) TrimLeft(trimStr string) *New {
	*s = New(strings.TrimLeft(s.Val(), trimStr))
	return s
}

// 字符串替换两端
func (s *New) Trim(trimStr string) *New {
	*s = New(strings.Trim(s.Val(), trimStr))
	return s
}

// 字符串替换
func (s *New) Replace(oldStr string, newStr string) *New {
	*s = New(strings.Replace(s.Val(), oldStr, newStr, -1))
	return s
}

// 字符串替换，将所有的oldStr替换为newStr
func (s *New) ReplaceM(newStr string, oldStr ...string) *New {
	*s = New(ReplaceM(s.Val(), newStr, oldStr...))
	return s
}

// 字符串替换，将所有的oldStr替换为newStr
func ReplaceM(sourceStr string, newStr string, oldStr ...string) (resultStr string) {
	for i := 0; i < len(oldStr); i++ {
		sourceStr = strings.Replace(sourceStr, oldStr[i], newStr, -1)
	}
	return sourceStr
}

// 返回第一次出现 start与第一次出现end之间的字符串
func (s *New) Between(start string, end string) *New {
	resultStr := Between(s.Val(), start, end)
	*s = New(resultStr)
	return s
}

// 返回第一次出现 start与第一次出现end之间的字符串
func Between(sourceStr, startStr, endStr string) string {
	startIndex := strings.Index(sourceStr, startStr)
	endIndex := strings.Index(sourceStr, endStr)
	return sourceStr[startIndex+len(startStr) : endIndex]
}

// 字符串添加
func (s *New) Append(appendStr string) *New {
	newStr := s.Val() + appendStr
	*s = New(newStr)
	return s
}

// 获得string
func (s *New) Val() string {
	return string(*s)
}

// 获得字符串倒序
func (s *New) Reverse() *New {
	runes := []rune(s.Val())
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	*s = New(string(runes))
	return s
}
