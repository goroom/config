package config

import (
	"testing"
)

type A struct {
	A int
	B int
	C string
	D string
}

func TestReflect(t *testing.T) {
	var a A
	err := Unmarshal("./config.ini", &a)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(a)
}
