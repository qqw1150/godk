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
