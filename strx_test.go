package strx

import (
	"testing"
)

func TestStr_Clean(t *testing.T) {
	var x = "abbcde ffsa*b*34"
	str := New(x)
	clean := str.Clean(" ", "*", "bb")
	t.Log(clean.Val())
}

func TestStr_ReplaceM(t *testing.T) {
	var x = "a|d*sdf"

	str := New(x)
	m := str.ReplaceM("rrr", "|", "*")
	t.Log(m.Val())
}

func TestStr_Append(t *testing.T) {
	var x = "asdf"

	str := New(x)

	s := str.Append(" | sdf").Append(" - fff").Val()
	t.Log(s)
}

func TestStr_Replace(t *testing.T) {
	var x = "777-${a}-${b}"
	y := New(x)
	splite := y.Replace("${a}", "999").Replace("${b}", "234").Val()

	t.Log(splite)
}

func TestStr_Reverse(t *testing.T) {
	y := New("abcd")
	reverse := y.Reverse().Val()
	t.Log(reverse)
}
