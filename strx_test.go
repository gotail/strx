package strx

import (
	"fmt"
	"testing"
)

func TestStr_HashCode(t *testing.T) {
	var testStr = "abcd"
	code := New(testStr).HashCode()
	t.Log(code)
}
func TestNew_NoConsecutiveSpaces(t *testing.T) {
	var testStr = " a   bcde   fgh      ijkl   mn"

	s := New(testStr).NoConsecutiveSpaces().Val()
	fmt.Println(s)
}

func TestIsEmpty(t *testing.T) {
	var testStr = "   "
	r := IsEmpty(testStr)
	r2 := IsNotEmpty(testStr)
	fmt.Println(r, r2)
}

func TestNew_TrimRight(t *testing.T) {
	var testStr = "世界和平fj"
	str := New(testStr)
	resultStr := str.TrimRight("fj").Val()
	t.Log(resultStr)
}

func TestNew_Between(t *testing.T) {
	var testStr = "今天天气真好啊，太阳当空照耀。"
	str := New(testStr)
	resultStr := str.Between("今天", "耀").Val()
	t.Log(resultStr)
}

func TestBetween(t *testing.T) {
	var newStr string = New(" 这*里|||有人民yyyy币         $yy#{money}   ..").
		NoConsecutiveSpaces().
		Trim(" ").
		Clean("*", "|", "yy").
		Replace("$", "￥").
		Replace("#{money}", "250").
		TrimRight(".").
		Append("元").
		Val()
	fmt.Println(newStr)
}

func TestStr_Clean(t *testing.T) {
	var testStr = "abbcde ffsa*b*34"
	str := New(testStr)
	resultStr := str.Clean(" ", "*", "bb").Val()
	t.Log(resultStr)
}

func TestStr_ReplaceM(t *testing.T) {
	var testStr = "a|d*sdf"

	str := New(testStr)
	m := str.ReplaceM("rrr", "|", "*")
	t.Log(m.Val())
}

func TestStr_Append(t *testing.T) {
	var testStr = "asdf"

	str := New(testStr)

	s := str.Append(" | sdf").Append(" - fff").Val()
	t.Log(s)
}

func TestStr_Replace(t *testing.T) {
	var testStr = "777-${a}-${b}"
	y := New(testStr)
	resultStr := y.Replace("${a}", "999").Replace("${b}", "234").Val()
	t.Log(resultStr)
}

func TestStr_Reverse(t *testing.T) {
	y := New("中国")
	reverse := y.Reverse().Val()
	t.Log(reverse)
}

func TestStr_ToLower(t *testing.T) {
	y := New("Hello,world!")
	lower := y.ToLower().Val()
	t.Log(lower)
}

func TestStr_ToUpper(t *testing.T) {
	y := New("Hello,world!")
	upper := y.ToUpper().Val()
	t.Log(upper)
}
