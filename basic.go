package main

import "fmt"

var isActive bool                   // 全域性變數宣告
var enabled, disabled = true, false // 忽略型別的宣告

func test() {
	var available bool // 一般宣告
	valid := false     // 簡短宣告
	available = true   // 賦值操作
}

func main() {
	var x, y int
	x, y = 8, 7

	const Pi float32 = 3.1415926

	fmt.Println("Hello Go", 1+2)
	fmt.Println(x+y, Pi)
}
