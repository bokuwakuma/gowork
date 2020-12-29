package main

import (
	"fmt"
)

type ContextInterface interface {
	NewContext() *Context
	SetName()
}

type Context struct {
	id   string
	name string
}

// コンストラクタ
func (con *Context) NewContext() *Context {
	return &Context{}
}

// ポインタレシーバ
func (con *Context) SetName(name string) {
	con.name = name
	fmt.Printf("SetName_addr_con: %p\n", con)
	fmt.Printf("SetName: My name is %s\n", con.name)
}

func main() {
	con := new(Context)
	// con := NewContext()
	// con := &Context{}
	fmt.Printf("main_addr_con: %p\n", con)

	con.name = "java"
	fmt.Printf("My first name is %s\n", con.name)

	con.SetName("gopher")
	fmt.Printf("My second name is %s\n", con.name)
}
