package test

import (
	"godk/str"
	"testing"
)

func TestSubstr(t *testing.T) {
	t.Log(str.Substr("水电费水电费单", 1, 3))
}

func TestUcfirst(t *testing.T) {
	t.Log(str.Ucfirst("hello"))
}

func TestUcwords(t *testing.T) {
	t.Log(str.Ucwords("hello	word nihao	hello xixix"))
}

func TestMd5(t *testing.T) {
	t.Log(str.Md5("hello"))
}

func TestMd5File(t *testing.T) {
	t.Log(str.Md5File("../1.txt"))
}

func TestNl2br(t *testing.T) {
	t.Log(str.Nl2br("This\r\nis\n\ra\nsting\r"))
}

func TestNumberFormat(t *testing.T) {
	t.Log(str.NumberFormat(67657567123456.456, 2, ".", ","))
}
