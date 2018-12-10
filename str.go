package common

import (
	"strings"
)

type Str string

// 字符串清除
func (s *Str) Clean(cleanStr ...string) *Str {
	return s.ReplaceM("", cleanStr...)
}

// 字符串替换
func (s *Str) Replace(oldStr string, newStr string) *Str {
	*s = Str(strings.Replace(string(*s), oldStr, newStr, -1))
	return s
}

// 字符串替换，将所有的oldStr替换为newStr
func (s *Str) ReplaceM(newStr string, oldStr ...string) *Str {
	for i := 0; i < len(oldStr); i++ {
		s = s.Replace(oldStr[i], newStr)
	}
	return s
}

// 字符串添加
func (s *Str) Append(appendStr string) *Str {
	newStr := string(*s) + appendStr
	*s = Str(newStr)
	return s
}

// 获得string
func (s *Str) Val() string {
	return string(*s)
}

// 获得字符串倒序
func (s *Str) Reverse() *Str {
	runes := []rune(*s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	*s = Str(string(runes))
	return s
}
