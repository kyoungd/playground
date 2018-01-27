package main

import (
	"fmt"
	"strings"
)

type shield struct {
	front bool
	back  bool
	right bool
	left  bool
}

type shBuilder struct {
	code string
}

func NewShieldBuilder() *shBuilder {
	return new(shBuilder)
}

func (sh *shBuilder) RaiseFront() *shBuilder {
	sh.code += "f"
	return sh
}

func (sh *shBuilder) RaiseBack() *shBuilder {
	sh.code += "b"
	return sh
}

func (sh *shBuilder) RaiseRight() *shBuilder {
	sh.code += "r"
	return sh
}

func (sh *shBuilder) RaiseLeft() *shBuilder {
	sh.code += "l"
	return sh
}

func (sh *shBuilder) Build() *shield {
	code := sh.code
	return &shield{
		front: strings.Contains(code, "f"),
		back:  strings.Contains(code, "b"),
		right: strings.Contains(code, "r"),
		left:  strings.Contains(code, "l"),
	}
}

func main() {
	builder := NewShieldBuilder()
	shield := builder.RaiseLeft().RaiseFront().Build()
	fmt.Printf("%+v \n", *shield)
}
