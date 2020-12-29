package main

import (
	"fmt"
)

// 構造体の定義（オブジェクト）
type Human struct {
	name  string
	age   int
	phone string
}

// HumanオブジェクトにSayHiメソッドを実装
func (h *Human) SayHi() {
	fmt.Println("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// HumanオブジェクトにSingメソッドを実装
func (h *Human) Sing(lyrics string) {
	fmt.Println("La lala,", lyrics)
}

type Student struct {
	Human  // 匿名フィールド
	school string
	loan   float32
}

type Employee struct {
	Human   // 匿名フィールド
	company string
	money   float32
}
