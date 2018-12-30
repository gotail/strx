package strx

import (
	"strings"
)

func New(x string) *str {
	n := str(x)
	return &n
}

type str string

// 没有连续的空格
func (s *str) NoConsecutiveSpaces() *str {
	cs := "  "
	a := string(*s)
	for {
		a = strings.Replace(a, cs, " ", -1)
		if !strings.Contains(a, cs) {
			s = New(strings.Trim(a, " "))
			return s
		}
	}
}

// 字符串清除
func (s *str) Clean(cleanStr ...string) *str {
	return s.ReplaceM("", cleanStr...)
}

// 字符串替换右侧
func (s *str) TrimRight(trimStr string) *str {
	s = New(strings.TrimRight(s.Val(), trimStr))
	return s
}

// 字符串替换左侧
func (s *str) TrimLeft(trimStr string) *str {
	s = New(strings.TrimLeft(s.Val(), trimStr))
	return s
}

// 字符串替换两端
func (s *str) Trim(trimStr string) *str {
	s = New(strings.Trim(s.Val(), trimStr))
	return s
}

// 字符串替换
func (s *str) Replace(oldStr string, newStr string) *str {
	s = New(strings.Replace(s.Val(), oldStr, newStr, -1))
	return s
}

// 字符串替换，将所有的oldStr替换为newStr
func (s *str) ReplaceM(newStr string, oldStr ...string) *str {
	*s = str(ReplaceM(s.Val(), newStr, oldStr...))
	return s
}

// 返回第一次出现 start与第一次出现end之间的字符串
func (s *str) Between(start string, end string) *str {
	resultStr := Between(s.Val(), start, end)
	s = New(resultStr)
	return s
}

// 字符串添加
func (s *str) Append(appendStr string) *str {
	newStr := s.Val() + appendStr
	s = New(newStr)
	return s
}

// 获得string
func (s *str) Val() string {
	return string(*s)
}

const primeRK = 16777619

func (s *str) HashCode() uint32 {
	strTemp := *s
	hash := uint32(0)
	for i := 0; i < len(strTemp); i++ {
		hash = hash*primeRK + uint32(strTemp[i])
	}
	return hash
}

// 获得字符串倒序
func (s *str) Reverse() *str {
	runes := []rune(s.Val())
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	s = New(string(runes))
	return s
}

// 英文字符串全部置为大写
func (s *str) ToUpper() *str {
	strNew := strings.ToUpper(s.Val())
	return New(strNew)
}

//英文字母全部置为小写
func (s *str) ToLower() *str {
	strNew := strings.ToLower(s.Val())
	return New(strNew)
}
