package main

import (
	"fmt"
)

import (
	"github.com/goroom/config"
)

type TestStruct struct {
	A int
	B int
	C string
	D bool
}

func main() {
	var ts TestStruct

	//	data := `{"A":1, "B":2, "C" ""}`
	config.Unmarshal("./config.ini", &ts)
	fmt.Println(ts)
}
