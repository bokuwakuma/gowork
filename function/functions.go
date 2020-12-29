package main

import (
	"fmt"
	"math"
)

// 引数の省略（一番最後の型に合わせる）
func add(x, y int) int {
	return x + y
}

// 複数の戻り値
func swap(x, y string) (string, string) {
	return y, x
}

// Naked return
// 戻り値に名前を宣言する
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 変数宣言
// 初期化子を与えるとその型で宣言される
var c, python, java = true, false, "no!"
var i, j int = 1, 2

func main() {
	// スコープが狭い方の変数で上書きされる
	var i int // 初期値0 boolはfalse

	// 暗黙的な型宣言 Syntax Suger
	// 以下は全て同値
	var j int = 10
	var k = 11
	l := 12

	fmt.Println(add(4, 4))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(17))
	fmt.Println(split(20))
	fmt.Println(i, c, python, java)
	fmt.Println(j, k, l)
	fmt.Println(math.Pow(5, 2))
}
