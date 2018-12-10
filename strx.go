package strx

import (
	"strings"
)

type New string

// 字符串清除
func (s *New) Clean(cleanStr ...string) *New {
	return s.ReplaceM("", cleanStr...)
}

// 字符串替换
func (s *New) Replace(oldStr string, newStr string) *New {
	*s = New(strings.Replace(string(*s), oldStr, newStr, -1))
	return s
}

// 字符串替换，将所有的oldStr替换为newStr
func (s *New) ReplaceM(newStr string, oldStr ...string) *New {
	for i := 0; i < len(oldStr); i++ {
		s = s.Replace(oldStr[i], newStr)
	}
	return s
}

// 字符串添加
func (s *New) Append(appendStr string) *New {
	newStr := string(*s) + appendStr
	*s = New(newStr)
	return s
}

// 获得string
func (s *New) Val() string {
	return string(*s)
}

// 获得字符串倒序
func (s *New) Reverse() *New {
	runes := []rune(*s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	*s = New(string(runes))
	return s
}
